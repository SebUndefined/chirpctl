---
title: Contexts
sidebar_position: 1
description: Configure and manage multiple ChirpStack environments with chirpctl.
---

# Contexts

A **context** is a reusable connection profile that tells chirpctl how to connect to a ChirpStack instance.

Each context can define its own:

- gRPC server address
- ChirpStack API version
- TLS configuration
- API token
- Default tenant ID
- Request timeout

Contexts allow you to work across multiple environments without repeatedly entering connection parameters.

## Why use contexts?

A typical LoRaWAN deployment may include several ChirpStack environments:

| Context       | Purpose                           |
|---------------|-----------------------------------|
| `development` | Local development and testing     |
| `staging`     | Preproduction validation          |
| `production`  | Production LoRaWAN infrastructure |

Instead of manually changing the server address and token for every command, chirpctl loads the appropriate settings from your context.

For example:

```bash
chirpctl --context staging get gateways
```

This command uses the `staging` connection profile for the current invocation.

It does not change the default context stored in your configuration.

## Initialize your configuration

The recommended way to create your first context is the interactive wizard:

```bash
chirpctl config init
```

The wizard guides you through the initial connection settings.

:::note

The initialization wizard requires an interactive terminal. It is not intended to run inside a non-interactive CI/CD pipeline.

For automation, configure chirpctl using a configuration file and environment variables.

:::

## Configuration file structure

chirpctl uses a YAML configuration file.

A configuration containing two contexts may look like this:

```yaml
version: 1
defaultContext: production

contexts:
  staging:
    addr: grpc-staging.example.com:443
    apiVersion: v4
    insecure: false
    caCert: ""
    serverName: grpc-staging.example.com
    token: ""
    tenantID: "11111111-1111-4111-8111-111111111111"
    timeout: 10s
    mTls:
      enabled: false
      clientCert: ""
      clientKey: ""

  production:
    addr: grpc.example.com:443
    apiVersion: v4

    # Transport
    insecure: false
    tlsSkipVerify: false
    caCert: /etc/chirpctl/certs/internal-ca.pem
    serverName: grpc.example.com

    # Mutual TLS (optional)
    mTls:
      enabled: false
      clientCert: ""
      clientKey: ""

    # Authentication
    token: ""
    tenantID: "11111111-1111-4111-8111-111111111111"

    # Request settings
    timeout: 15s

ui:
  output: table-styled
```

The empty `token` fields in this example are intentional.

You can supply an API token through the `CHIPCTL_TOKEN` environment variable instead of persisting it in YAML.

:::warning

A configuration file may contain API credentials in plaintext.

Never commit a configuration containing real tokens to source control.

The example tenant IDs and server addresses above are placeholders.

:::

## Configuration properties

### Top-level properties

| Property         | Description                       |
|------------------|-----------------------------------|
| `version`        | Configuration file schema version |
| `defaultContext` | Name of the default context       |
| `contexts`       | Map of named connection profiles  |
| `ui.output`      | Preferred output format           |

The current configuration file format uses `version: 1`.

### Context properties

| Property          | Required | Description                                         |
|-------------------|----------|-----------------------------------------------------|
| `addr`            | Yes      | gRPC server address in `host:port` format           |
| `apiVersion`      | Yes      | ChirpStack API version, currently `v4`              |
| `insecure`        | No       | Use plaintext gRPC without TLS encryption           |
| `tlsSkipVerify`   | No       | Enable TLS without verifying the server certificate |
| `caCert`          | No       | Path to an additional trusted CA certificate bundle |
| `mTls.enabled`    | No       | Enable mutual TLS client authentication             |
| `mTls.clientCert` | No       | Path to the client certificate                      |
| `mTls.clientKey`  | No       | Path to the client private key                      |
| `serverName`      | No       | Expected TLS server name                            |
| `token`           | No       | ChirpStack API token                                |
| `tenantID`        | No       | Default tenant UUID                                 |
| `timeout`         | No       | Command timeout, such as `10s`                      |
| `mTls`            | No       | Reserved client-certificate configuration           |

:::caution

The `mTls` fields exist in the configuration schema, but client-certificate authentication should not be considered operational until supported by the gRPC connection implementation.

Use standard TLS and API-token authentication for the current release.

:::

## List configured contexts

Use:

```bash
chirpctl config get-contexts
```

This displays the available contexts and identifies the default one.

For structured output:

```bash
chirpctl config get-contexts --output json
```

Or:

```bash
chirpctl config get-contexts --output yaml
```

The output provides connection metadata without exposing the actual stored API tokens.

A `hasToken` field indicates whether a token is stored in the context. It does not necessarily reflect a token supplied through an environment variable.

## Inspect the default context

To display the default context:

```bash
chirpctl config current-context
```

For JSON output:

```bash
chirpctl config current-context --output json
```

The command reads the default context recorded in the configuration file.

It should not be confused with the `--context` flag, which selects a context for an individual command.

## Create or update a context

Use `config set-context` to create or update a named profile.

For example, to configure a staging environment:

```bash
chirpctl config set-context \
  --context staging \
  --addr grpc-staging.example.com:443 \
  --api-version v4 \
  --timeout 10s
```

The `--context` flag identifies the profile to create or update.

When updating an existing context, only the explicitly supplied fields are intended to change.

For example:

```bash
chirpctl config set-context \
  --context staging \
  --timeout 20s
```

This changes the timeout while preserving the other existing connection settings.

:::tip

Avoid passing real API tokens through `--token` in shared terminals or automation environments.

Command-line arguments can be exposed through shell history, process inspection, or CI logs.

Prefer an environment variable or the interactive initialization wizard.

:::

## Select a context for one command

You can override the default context without editing the configuration file:

```bash
chirpctl --context production get gateways
```

The selection applies to that invocation only.

This is useful for scripts and administrative operations across multiple environments.

## Change the default context

The `defaultContext` property identifies the profile used when no context override is supplied.

For example:

```yaml
defaultContext: production
```

To change it manually:

1. Open your configuration file.
2. Update `defaultContext`.
3. Ensure the specified context exists in `contexts`.
4. Save the file.
5. Verify the result with `chirpctl config current-context`.

Only one default context can be selected at a time.

## Delete a context

To delete a context:

```bash
chirpctl config delete-context staging
```

The default context cannot be deleted while it remains selected as the default.

Choose another default context before deleting it.

:::warning

Deleting a context removes its saved connection settings, including any token stored in that profile.

Keep a secure backup when necessary.

:::

## Tenant-aware contexts

A context can store a default ChirpStack tenant ID:

```yaml
tenantID: "11111111-1111-4111-8111-111111111111"
```

This can simplify operations that support tenant filtering.

For example:

```bash
chirpctl get gateways
```

The gateway list command can use the tenant associated with the active context.

To override the tenant for one request:

```bash
chirpctl get gateways \
  --tenant-id 22222222-2222-4222-8222-222222222222
```

For a global administrator who needs to request gateways across tenants:

```bash
chirpctl get gateways --all-tenants
```

The `--all-tenants` flag requests an unfiltered gateway list.

It does not bypass ChirpStack authorization or grant additional privileges.

The `--tenant-id` and `--all-tenants` flags are mutually exclusive.

## Context selection precedence

chirpctl selects the effective context in this order:

1. Explicit `--context` flag
2. `CHIPCTL_PROFILE` environment variable
3. `defaultContext` from the YAML configuration file

For example:

```bash
chirpctl --context staging get gateways
```

takes precedence over:

```bash
CHIPCTL_PROFILE=production
```

For details about how individual connection parameters are merged, see [Configuration Precedence](./configuration-precedence.md).

## Best practices

- Use separate contexts for development, staging, and production.
- Prefer dedicated API tokens with appropriate permissions.
- Keep TLS enabled for remote infrastructure.
- Avoid saving sensitive tokens in version-controlled files.
- Always verify the selected context before modifying production resources.
- Use explicit `--context` flags in scripts to avoid depending on workstation defaults.
- Avoid concurrent processes modifying the same configuration file.

## Related documentation

- [Authentication](./authentication.md)
- [TLS and Certificates](./tls.md)
- [Configuration Precedence](./configuration-precedence.md)
