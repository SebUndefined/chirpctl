<p align="center">
  <img src="./docs/assets/chirpctl-banner.png" width="80%" alt="Sublime's custom image"/>
</p>

<div align="center">

# A cross-platform CLI to interact with ChirpStack over gRPC

</div>

<br>
<div align="center">

💡 [Get started](https://sebundefined.github.io/chirpctl/docs/)&ensp;•&ensp;🌐 [Website](https://sebundefined.github.io/chirpctl/)&ensp;•&ensp;📚 [Documentation](https://sebundefined.github.io/chirpctl/docs/)&ensp;
</div>

**chirpctl** is an open-source, cross-platform command-line interface for managing [ChirpStack](https://www.chirpstack.io/) infrastructure.

Built with Go and powered by gRPC, it brings common LoRaWAN network operations directly into your terminal.

Whether you manage a single ChirpStack instance or work across multiple environments, chirpctl provides a consistent, script-friendly interface.

# Quick Start

This guide walks you through your first chirpctl connection to ChirpStack v4.

By the end, you will have:

- Configured a connection profile
- Verified your active context
- Listed your ChirpStack gateways
- Requested structured JSON output

## Prerequisites

Before continuing, make sure you have:

1. [Installed chirpctl](./installation.md).
2. Access to a running ChirpStack v4 instance.
3. The gRPC server address and port.
4. A valid ChirpStack API token.
5. A tenant ID, if required by your account or operation.

:::tip

The ChirpStack web interface and gRPC server may use different addresses or ports.

Use the address of the gRPC endpoint, not automatically the URL of the web interface.

:::

## Step 1 — Initialize your configuration

Start the interactive configuration wizard:

```bash
chirpctl config init
```

The wizard will guide you through the initial setup.

You will configure:

| Setting | Description |
| --- | --- |
| Context name | A unique name for the connection |
| API version | ChirpStack API version |
| Server address | gRPC host and port |
| TLS configuration | Whether the connection uses TLS |
| API token | Credentials used for authentication |
| Tenant ID | Optional default tenant identifier |
| Timeout | Maximum time allowed for operations |

For example, you might create a context named `production` pointing to:

```text
grpc.example.com:443
```

Use the actual endpoint and credentials of your ChirpStack instance.

:::warning

API tokens are secrets. Do not commit your configuration file to Git, share it publicly, or include token values in issue reports.

:::

## Step 2 — Check your active context

Display the current context:

```bash
chirpctl config current-context
```

This shows the selected context and its connection settings without exposing the API token.

You can also list all configured contexts:

```bash
chirpctl config get-contexts
```

Using multiple contexts helps keep environments such as development, staging, and production separate.

## Step 3 — List your gateways

With your context configured, query ChirpStack:

```bash
chirpctl get gateways
```

chirpctl connects to the configured ChirpStack gRPC endpoint and retrieves the available gateways.

By default, the command uses a human-readable table format.

:::note

The gateways returned depend on the permissions associated with your API token and any tenant filtering.

An empty result does not necessarily indicate a connection failure.

:::

## Step 4 — Request JSON output

For automation or integration with other tools, request JSON:

```bash
chirpctl get gateways --output json
```

You can also request YAML:

```bash
chirpctl get gateways --output yaml
```

Or CSV:

```bash
chirpctl get gateways --output csv
```

Structured output is useful when writing scripts or integrating chirpctl into existing workflows.

## Step 5 — Work with multiple contexts

If you have multiple ChirpStack instances configured, you can select a context explicitly for an invocation:

```bash
chirpctl --context production get gateways
```

The `--context` option allows you to target a specific environment without changing the default context stored in your configuration.

This is especially useful when comparing development, staging, and production systems.

## Step 6 — Target a tenant

Some operations may require a tenant identifier.

If you configured a default tenant in your context, chirpctl can use it for commands that support tenant filtering.

For gateway queries, you can provide an explicit tenant ID:

```bash
chirpctl get gateways --tenant-id <tenant-uuid>
```

Replace `<tenant-uuid>` with your actual tenant identifier.

:::info

Tenant access is enforced by ChirpStack. Providing a tenant ID does not grant additional permissions to the API token.

:::


## Choose your connection security

chirpctl supports several gRPC connection modes.

Choose the mode that matches your ChirpStack deployment:

| Environment                               | Recommended configuration |
|-------------------------------------------|---------------------------|
| Public server with valid certificate      | Standard TLS              |
| Private server with internal CA           | TLS with `caCert`         |
| Testing server with untrusted certificate | TLS with `tlsSkipVerify`  |
| Local service over a trusted SSH tunnel   | Plaintext with `insecure` |
| Server requiring client certificates      | mTLS                      |

For a standard production endpoint, use TLS with certificate verification.

For a private server with an internal CA, configure the trusted CA certificate.

For an SSH tunnel exposing a local plaintext gRPC endpoint, use `insecure: true`.

See [TLS and Certificates](../02-configuration/tls.md) for detailed examples and security recommendations.


## Troubleshooting your first connection

### Unauthenticated

If ChirpStack returns `Unauthenticated`, check:

- Whether the API token is valid.
- Whether the token has the required permissions.
- Whether your context uses the intended credentials.

### Connection refused or deadline exceeded

Verify:

- The gRPC server address and port.
- Network access to the server.
- Firewall and reverse-proxy configuration.
- TLS settings and certificate configuration.

### TLS handshake or certificate errors

Verify that the gRPC endpoint uses the expected TLS configuration and that the certificate matches the server name.

Avoid disabling transport security merely to bypass certificate validation.

## Explore more commands

chirpctl provides built-in help:

```bash
chirpctl --help
```

You can also display help for specific command groups:

```bash
chirpctl config --help
```

```bash
chirpctl get --help
```

## What's next?

You have successfully configured your first connection and executed your first ChirpStack query.

Continue with:

- **Configuration** to understand contexts, authentication, TLS, and precedence.
- **Guides** for practical workflows.
- **CLI Reference** for available commands and flags.

You can also browse the source code on [GitHub](https://github.com/sebundefined/chirpctl).
