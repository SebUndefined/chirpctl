---
title: Authentication
sidebar_position: 2
description: Authenticate chirpctl requests using ChirpStack API tokens.
---

# Authentication

chirpctl authenticates with ChirpStack using an **API token**.

The token identifies the caller and allows ChirpStack to evaluate its permissions for each request.

Authentication is performed over the ChirpStack gRPC API.

## API token authentication

For authenticated requests, chirpctl sends the API token as a Bearer credential in gRPC metadata:

```text
authorization: Bearer <token>
```

You do not need to construct this header manually.

chirpctl attaches the credential to outgoing gRPC requests when a token is configured.

:::warning

An API token is a secret.

Anyone who obtains a valid token may be able to perform actions permitted by that token.

Treat API tokens like passwords.

:::

## Obtain an API token

Before connecting, obtain a token from the ChirpStack instance you intend to manage.

The exact token creation procedure depends on the ChirpStack version and the permissions of your account.

For ChirpStack v4:

1. Sign in to the ChirpStack web interface.
2. Locate the API key management functionality available to your account.
3. Create a key with the intended access scope.
4. Store the generated token securely.

API keys may have different scopes and permissions.

Choose the least privileged key that supports your intended operations.

For authoritative API and authorization details, refer to the [ChirpStack documentation](https://www.chirpstack.io/docs/).

## Configure a token

There are several ways to supply an API token.

### Option 1 — Interactive configuration

When creating your first context, run:

```bash
chirpctl config init
```

The wizard allows you to configure your credentials interactively.

Be aware that credentials saved to the configuration file remain stored locally.

### Option 2 — Environment variable

For terminal sessions and automation, use the `CHIPCTL_TOKEN` environment variable.

On Linux or macOS:

```bash
export CHIPCTL_TOKEN="YOUR_API_TOKEN"
```

On PowerShell:

```powershell
$env:CHIPCTL_TOKEN = "YOUR_API_TOKEN"
```

Replace the placeholder with a real token.

The variable is used by chirpctl when resolving connection settings.

It can override a token stored in the selected YAML context.

:::tip

In CI/CD pipelines, inject `CHIPCTL_TOKEN` from your CI provider's secret store.

Do not hardcode credentials directly in workflow files, scripts, or committed configuration.

:::

### Option 3 — Configuration file

A context can contain a saved token:

```yaml
contexts:
  production:
    addr: grpc.example.com:443
    apiVersion: v4
    insecure: false
    token: "YOUR_API_TOKEN"
```

This is convenient for local usage, but the token is stored as plaintext in the YAML file.

Protect the file with appropriate operating-system permissions.

### Option 4 — Command-line flag

chirpctl also provides a `--token` flag.

However, passing credentials as command-line arguments can expose them in shell history, process listings, or logs.

Avoid this method for long-lived production credentials.

## Authentication versus authorization

Authentication answers:

**Who is making this request?**

Authorization answers:

**Is the caller allowed to perform this operation?**

A valid API token does not automatically grant access to every resource or tenant.

For example:

```bash
chirpctl get gateways
```

may succeed while returning only gateways visible to the caller.

A request against another tenant may fail if the token does not have sufficient permissions.

## Tenant access

A context may contain a default tenant ID:

```yaml
tenantID: "11111111-1111-4111-8111-111111111111"
```

For gateway listing operations, you can supply an explicit tenant filter:

```bash
chirpctl get gateways \
  --tenant-id 11111111-1111-4111-8111-111111111111
```

The tenant ID is a request filter, not an authentication credential.

It does not grant access to that tenant.

### Global administrators

A global administrator may be able to request data across tenants, depending on the token's privileges.

For gateway listing:

```bash
chirpctl get gateways --all-tenants
```

This disables the context's default tenant filter for that request.

ChirpStack still enforces its authorization rules.

## Secure token handling

Follow these recommendations:

### Use dedicated credentials

Prefer a token intended specifically for chirpctl operations instead of reusing unrelated service credentials.

### Apply least privilege

Give each token only the permissions necessary for its intended use.

### Separate environments

Use distinct tokens for development, staging, and production whenever possible.

### Keep tokens out of Git

Never commit actual credentials to:

- `config.yaml`
- Example configuration files
- Documentation
- Test fixtures
- CI/CD workflow definitions

### Rotate credentials

If a token may have been exposed:

1. Revoke or disable it in ChirpStack.
2. Generate a replacement.
3. Update the relevant secret store or local configuration.
4. Verify the new token with a permitted operation.


## Transport security and API tokens

chirpctl supports API-token authentication over both TLS and explicitly configured plaintext gRPC connections.

With TLS, credentials are protected by the encrypted connection.

With plaintext gRPC, the API token is transmitted without gRPC-level encryption.

This can be appropriate when the connection is carried through a trusted, encrypted transport such as an SSH tunnel.

For example:

```yaml
contexts:
  private-tunnel:
    addr: 127.0.0.1:1700
    apiVersion: v4
    insecure: true
    token: ""
```

The API token can be supplied separately:

```bash
export CHIPCTL_TOKEN="YOUR_API_TOKEN"

chirpctl --context private-tunnel get gateways
```

:::warning

Plaintext gRPC does not protect the confidentiality of API tokens by itself.

Only use plaintext mode when the complete communication path is appropriately secured.

Avoid sending credentials over publicly accessible or untrusted plaintext connections.

:::

For general production deployments, TLS with certificate verification remains the recommended configuration.

See [TLS and Certificates](./tls.md) for more information.


## Troubleshooting

### Unauthenticated

A gRPC `Unauthenticated` response may indicate:

- A missing API token
- An invalid or revoked token
- A token not accepted by the target ChirpStack instance
- A server-side authentication configuration issue

Check that the intended context and token source are selected.

### PermissionDenied

A `PermissionDenied` response indicates that the caller is not authorized to perform the requested operation.

Check the token's scope and permissions.

### Unexpected empty results

An empty list is not necessarily an authentication failure.

Verify tenant filters and resource visibility.

For gateway queries, inspect whether the active context includes a default tenant ID.

## Related documentation

- [Contexts](./contexts.md)
- [TLS and Certificates](./tls.md)
- [Configuration Precedence](./configuration-precedence.md)
