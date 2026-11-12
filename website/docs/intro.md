---
id: intro
title: Overview
sidebar_label: Overview
sidebar_position: 1
slug: /
description: Learn what chirpctl is and how to get started with ChirpStack from your terminal.
---

# Meet chirpctl

**chirpctl** is an open-source, cross-platform command-line interface for managing [ChirpStack](https://www.chirpstack.io/) infrastructure.

Built with Go and powered by gRPC, it brings common LoRaWAN network operations directly into your terminal.

Whether you manage a single ChirpStack instance or work across multiple environments, chirpctl provides a consistent, script-friendly interface.

## Why chirpctl?

ChirpStack offers a powerful web interface and API. But some operations are more convenient when performed directly from a terminal, especially when working with multiple environments or automation.

chirpctl is designed around a familiar command-line workflow:

```bash
chirpctl get gateways
```

Its goal is to make routine ChirpStack operations easier to execute, repeat, and integrate into your existing tools.

## Key capabilities

### Native gRPC integration

Communicate directly with the ChirpStack v4 gRPC API without requiring an additional REST proxy.

### Multiple contexts

Define reusable connection profiles to manage development, staging, and production environments.

### Secure connections

Configure TLS, custom CA certificates, server names, and API tokens.

### Tenant-aware configuration

Associate a default tenant with a context and override it when supported by individual commands.

### Structured output

Choose the format that best fits your workflow:

- Styled tables for interactive terminal use
- Plain tables for simple text output
- JSON and YAML for structured data
- CSV for exports and external tools

### Cross-platform distribution

Precompiled binaries are available for Linux, macOS, and Windows, targeting AMD64 and ARM64 architectures.

## Supported ChirpStack versions

| ChirpStack version | Status |
| --- | --- |
| v4 | Supported |
| v3 | Planned |
| Future versions | Architecture designed for extensibility |

:::info

Support for ChirpStack v3 is not implemented yet. The current release targets ChirpStack v4.

:::

## Get started

New to chirpctl? Follow these two steps:

1. [Install chirpctl](./01-getting-started/installation.md) on your operating system.
2. [Complete the Quick Start](./01-getting-started/quickstart.md) to configure your first context and query ChirpStack.

## Explore the documentation

Once you have a working installation, explore:

- **Configuration** — Contexts, authentication, TLS, and configuration precedence.
- **Guides** — Practical examples for common operations and scripting.
- **CLI Reference** — Commands, arguments, flags, and environment variables.
- **Development** — Project architecture, testing, and contributing.
- **Troubleshooting** — Diagnose connection, authentication, and installation issues.

## Open source

chirpctl is developed in the open. You can explore its source code, report issues, and follow upcoming features on GitHub.

- [GitHub repository](https://github.com/sebundefined/chirpctl)
- [Latest releases](https://github.com/sebundefined/chirpctl/releases/latest)
- [Report an issue](https://github.com/sebundefined/chirpctl/issues)
