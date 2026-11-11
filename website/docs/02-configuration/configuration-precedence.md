---
title: Configuration Precedence
sidebar_position: 4
description: Understand how chirpctl resolves YAML configuration, environment variables and command-line flags.
---

# Configuration Precedence

chirpctl supports multiple sources of configuration.

Connection settings can be defined in a YAML file, overridden through environment variables, or supplied using command-line flags.

Understanding the resolution order is important when working across environments or using chirpctl in automation.

## Resolution order

For most connection settings, chirpctl applies the following precedence:

| Priority | Source | Example |
| --- | --- | --- |
| 1 — Highest | Command-line flags | `--addr grpc.example.com:443` |
| 2 | Environment variables | `CHIPCTL_ADDR` |
| 3 | Selected YAML context | `contexts.production.addr` |
| 4 — Lowest | Built-in defaults | Default timeout |

The effective configuration is produced by merging these sources.

An explicitly supplied command-line flag overrides the corresponding environment variable and YAML value.

## Example

Consider the following context:

```yaml
defaultContext: production

contexts:
  production:
    addr: grpc-production.example.com:443
    apiVersion: v4
    insecure: false
    token: ""
    timeout: 10s
```

Now configure an environment variable:

```bash
export CHIPCTL_ADDR=grpc-staging.example.com:443
```

Running:

```bash
chirpctl get gateways
```

uses:

```text
grpc-staging.example.com:443
```

The environment variable overrides the address stored in the YAML context.

However:

```bash
chirpctl \
  --addr grpc-test.example.com:443 \
  get gateways
```

uses:

```text
grpc-test.example.com:443
```

because the explicit command-line flag has the highest priority.

Only the overridden property changes. Other settings continue to come from the selected context or their own overrides.

## Context selection

Context selection follows a separate precedence rule:

| Priority | Source |
| --- | --- |
| 1 | `--context` |
| 2 | `CHIPCTL_PROFILE` |
| 3 | YAML `defaultContext` |

For example:

```yaml
defaultContext: production

contexts:
  production:
    addr: grpc-production.example.com:443
    apiVersion: v4

  staging:
    addr: grpc-staging.example.com:443
    apiVersion: v4
```

Without any override:

```bash
chirpctl get gateways
```

selects `production`.

With an environment variable:

```bash
export CHIPCTL_PROFILE=staging
chirpctl get gateways
```

selects `staging`.

An explicit flag takes precedence:

```bash
chirpctl --context production get gateways
```

selects `production`, even when `CHIPCTL_PROFILE=staging`.

:::note

The environment variable is named `CHIPCTL_PROFILE`, not `CHIPCTL_CONTEXT`.

The name reflects the current configuration resolver implementation.

:::

## Supported environment variables

The following variables are handled by the connection options resolver:

| Environment variable | Corresponding setting |
| --- | --- |
| `CHIPCTL_PROFILE` | Selected context |
| `CHIPCTL_ADDR` | gRPC server address |
| `CHIPCTL_INSECURE` | Plaintext gRPC mode |
| `CHIPCTL_CACERT` | Custom CA certificate path |
| `CHIPCTL_SERVER_NAME` | TLS server name |
| `CHIPCTL_TOKEN` | API token |
| `CHIPCTL_TIMEOUT` | Command timeout |
| `CHIPCTL_OUTPUT` | Output format |
| `CHIPCTL_VERBOSE` | Verbose logging |
| `CHIPCTL_API_VERSION` | ChirpStack API version |

Use the `CHIPCTL_` prefix exactly as shown.

### Boolean values

Boolean variables such as `CHIPCTL_INSECURE` and `CHIPCTL_VERBOSE` accept values that Go can parse as booleans.

For example:

```bash
export CHIPCTL_VERBOSE=true
```

Or:

```bash
export CHIPCTL_INSECURE=false
```

An empty but present boolean environment variable is treated as `true` by the current resolver.

For predictable behavior, use explicit `true` or `false` values.

### Timeouts

The `CHIPCTL_TIMEOUT` variable accepts Go duration strings.

Examples:

```bash
export CHIPCTL_TIMEOUT=5s
```

```bash
export CHIPCTL_TIMEOUT=30s
```

```bash
export CHIPCTL_TIMEOUT=1m
```

The environment override is applied when the duration is valid and positive.

If no valid timeout is configured, chirpctl uses its built-in timeout.

## Command-line overrides

The corresponding root-level flags include:

| Flag | Purpose |
| --- | --- |
| `--context` | Select a context |
| `--addr` | Override gRPC endpoint |
| `--insecure` | Use plaintext gRPC |
| `--cacert` | Override CA certificate path |
| `--server-name` | Override TLS server name |
| `--token` | Override API token |
| `--timeout` | Override timeout |
| `--output`, `-o` | Select output format |
| `--verbose`, `-v` | Enable verbose logging |
| `--api-version` | Select API version |

For example:

```bash
chirpctl \
  --context production \
  --timeout 30s \
  --output json \
  get gateways
```

This command selects the production context, applies a 30-second timeout, and requests JSON output.

:::warning

Avoid using the `--token` flag with real credentials in scripts or shared environments.

Use a secret-managed environment variable instead.

:::

## Configuration file location

The configuration file is loaded before connection settings are resolved.

You can override its location using:

```text
CHIPCTL_CONFIG
```

On Linux or macOS:

```bash
export CHIPCTL_CONFIG="$HOME/.config/chirpctl/config.yaml"
```

On PowerShell:

```powershell
$env:CHIPCTL_CONFIG = "$HOME\.config\chirpctl\config.yaml"
```

This variable tells chirpctl which configuration file to load.

If an existing directory is supplied instead of a file, chirpctl uses `config.yaml` within that directory.

If a relative path is supplied, it is resolved relative to the current working directory.

:::tip

Using `CHIPCTL_CONFIG` is convenient for isolated development environments and integration tests.

It prevents automated tests from modifying your regular configuration file.

:::

The default configuration location depends on the operating system and the chirpctl build.

## Output configuration

A YAML file can contain a preferred output format:

```yaml
ui:
  output: table-styled
```

Supported formats include:

| Format | Intended usage |
| --- | --- |
| `table-styled` | Human-readable terminal output |
| `table` | Plain table |
| `json` | Scripts and structured processing |
| `yaml` | Structured configuration and inspection |
| `csv` | Spreadsheet and external data processing |

You can request a format explicitly:

```bash
chirpctl get gateways --output json
```

The command-line flag is recommended when scripts require a deterministic output format.

:::note

The UI output preference is applied separately from some connection settings.

If your environment combines `ui.output` and `CHIPCTL_OUTPUT`, verify their interaction against the installed chirpctl version.

An explicit `--output` flag is the safest choice for automation.

:::

## Command-specific options

Not every option participates in the global connection resolver.

For example, gateway listing supports:

```bash
chirpctl get gateways --tenant-id <tenant-uuid>
```

and:

```bash
chirpctl get gateways --all-tenants
```

These are command-specific filters.

The effective tenant selection also takes into account the tenant configured in the selected context when supported by the command.

Do not assume that every YAML context field has a corresponding environment variable or root-level flag.

## Using chirpctl in CI/CD

For automated workflows, prefer explicit configuration and environment selection.

A recommended approach is:

1. Provide a configuration file for the intended environment.
2. Select the context explicitly.
3. Inject API credentials from the CI secret store.
4. Use an explicit output format.
5. Avoid printing sensitive configuration or credentials in logs.

For example:

```bash
chirpctl \
  --context staging \
  --output json \
  get gateways
```

The API token can be supplied using `CHIPCTL_TOKEN`.

:::caution

Environment variables are not automatically secret.

On shared systems, processes and debugging tools may expose environment data.

Use isolated runners, secret masking and least-privilege credentials.

:::

## Troubleshooting configuration resolution

### The wrong server is selected

Check:

1. The explicit `--addr` flag
2. The `CHIPCTL_ADDR` environment variable
3. The selected context's `addr` field

### The wrong context is selected

Check:

1. The `--context` flag
2. The `CHIPCTL_PROFILE` environment variable
3. The YAML `defaultContext` property

### An option appears to be ignored

Verify that:

- The flag belongs to the command being executed.
- The environment variable name is correct.
- The value is valid for its expected type.
- A higher-priority override is not present.

### A command fails before connecting

chirpctl must load its configuration before executing most resource commands.

Verify that the configuration file exists, can be read, and contains the required context.

## Related documentation

- [Contexts](./contexts.md)
- [Authentication](./authentication.md)
- [TLS and Certificates](./tls.md)
