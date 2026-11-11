---
title: TLS and Certificates
sidebar_position: 3
description: Secure gRPC connections to ChirpStack using TLS and custom CA certificates.
---

# TLS and Certificates

chirpctl communicates with ChirpStack through gRPC.

For remote infrastructure, **TLS is strongly recommended** to protect API credentials and network traffic.

chirpctl supports TLS server verification, custom CA certificates, and explicit TLS server names.

## Connection modes

chirpctl provides two transport modes:

| Configuration | Transport | Encrypted |
| --- | --- | --- |
| `insecure: false` | gRPC over TLS | Yes |
| `insecure: true` | Plaintext gRPC | No |

:::danger

The `insecure` setting does not simply disable certificate validation.

Setting `insecure: true` disables TLS entirely.

Do not enable it to work around an invalid certificate on a production endpoint.

:::

## Standard TLS connection

For a ChirpStack server using a certificate trusted by your operating system, use:

```yaml
contexts:
  production:
    addr: grpc.example.com:443
    apiVersion: v4
    insecure: false
    serverName: grpc.example.com
    caCert: ""
    token: ""
    timeout: 10s
```

The important settings are:

- `addr`: gRPC endpoint, including the port
- `insecure: false`: enable TLS
- `serverName`: expected certificate identity
- `caCert`: optional custom CA certificate

If `caCert` is empty, the client uses the system's default trusted certificate authorities.

## Understanding the server address

The `addr` property must use the `host:port` format:

```yaml
addr: grpc.example.com:443
```

It is not an HTTP URL.

Do not use:

```yaml
addr: https://grpc.example.com
```

The gRPC client expects a network address, not a URL with a scheme.

:::note

The port depends on your ChirpStack deployment, reverse proxy, and TLS termination configuration.

Port `443` is common for publicly exposed TLS services, but is not guaranteed.

:::

## TLS server name

TLS validates the identity of the server against its certificate.

The `serverName` property controls the name expected during TLS verification.

For example:

```yaml
addr: 192.0.2.10:443
serverName: grpc.example.com
insecure: false
```

This configuration connects to the specified address while validating the certificate for `grpc.example.com`.

It can be useful when:

- Accessing a service through a private IP address
- Using a reverse proxy
- Connecting through a VPN
- Testing a specific endpoint without changing DNS

### Default behavior

When no explicit `serverName` is configured, chirpctl derives the TLS server name from the host portion of `addr`.

For example:

```yaml
addr: grpc.example.com:443
serverName: ""
```

The effective TLS server name is derived from `grpc.example.com`.

For predictable production configurations, setting the expected hostname explicitly is reasonable.

## Custom certificate authorities

Some private ChirpStack deployments use certificates issued by an internal certificate authority.

If the certificate chain is not trusted by the operating system, you may configure a custom CA certificate.

For example:

```yaml
contexts:
  private-network:
    addr: grpc.internal.example.com:443
    apiVersion: v4
    insecure: false
    caCert: /etc/chirpctl/certs/internal-ca.pem
    serverName: grpc.internal.example.com
    token: ""
    timeout: 10s
```

The file must contain one or more trusted CA certificates in PEM format.

### Important trust-store behavior

When a custom `caCert` is supplied, the current chirpctl gRPC client builds its trust pool from that PEM file.

This replaces the default system trust pool for that connection rather than automatically appending to it.

If your deployment requires multiple private CAs, the configured PEM file must contain the appropriate certificates.

### Configure the CA path from the CLI

To use a custom CA for a single command:

```bash
chirpctl \
  --cacert /etc/chirpctl/certs/internal-ca.pem \
  get gateways
```

To override the expected server identity:

```bash
chirpctl \
  --server-name grpc.internal.example.com \
  get gateways
```

These options override the corresponding connection settings for the invocation.

## Plaintext gRPC

For a local test server that does not support TLS, chirpctl can use plaintext gRPC:

```yaml
contexts:
  local:
    addr: localhost:8080
    apiVersion: v4
    insecure: true
    token: ""
    timeout: 10s
```

This disables transport encryption.

You can also use the command-line flag:

```bash
chirpctl \
  --addr localhost:8080 \
  --insecure \
  get gateways
```

:::warning

Use plaintext gRPC only in an environment where the transport is appropriately isolated and trusted.

For production, prefer TLS.

If an API token is provided over plaintext gRPC, its confidentiality is not protected by TLS.

:::

## TLS and reverse proxies

Many ChirpStack installations place a reverse proxy or ingress controller in front of the gRPC server.

A typical architecture is:

```text
chirpctl
   |
   | gRPC + TLS
   v
Reverse proxy / ingress
   |
   | Internal upstream connection
   v
ChirpStack gRPC API
```

The reverse proxy must support the gRPC protocol and HTTP/2 as required by the deployment.

TLS may terminate at the reverse proxy rather than inside ChirpStack.

From chirpctl's perspective, the external endpoint must expose a compatible gRPC service.

## TLS versus mTLS

Standard TLS authenticates the server to the client and protects the connection.

Mutual TLS (mTLS) additionally authenticates the client using a client certificate.

chirpctl's configuration schema currently includes the following fields:

```yaml
mTls:
  enabled: false
  clientCert: ""
  clientKey: ""
```

However, the current gRPC connection implementation does not load and attach the configured client certificate and private key.

:::caution

Do not rely on the `mTls` configuration for client-certificate authentication in the current release.

The schema anticipates this capability, but it is not yet wired into the runtime transport.

:::

API-token authentication remains separate from TLS.

## Troubleshooting TLS

### Certificate signed by unknown authority

The server certificate cannot be validated against the trusted CA certificates.

Check:

- Whether the server presents the expected certificate chain
- Whether the CA certificate is trusted by the operating system
- Whether `caCert` points to the correct PEM file
- Whether intermediate certificates are configured correctly

Do not solve this error by blindly enabling `--insecure`.

### Certificate is not valid for the requested hostname

The certificate identity does not match the configured TLS server name.

Verify:

- The hostname in `addr`
- The configured `serverName`
- The certificate's Subject Alternative Name entries
- Reverse-proxy certificate selection

### TLS handshake failed

Possible causes include:

- Connecting with TLS to a plaintext endpoint
- Connecting without TLS to a TLS endpoint
- An unsupported or misconfigured proxy
- A certificate or protocol configuration problem

### Deadline exceeded

A timeout may indicate an unreachable server, blocked port, stalled TLS handshake, or slow gRPC operation.

Check the server address and network path before increasing the timeout.

## Recommended production configuration

For a production ChirpStack instance:

```yaml
contexts:
  production:
    addr: grpc.example.com:443
    apiVersion: v4
    insecure: false
    caCert: ""
    serverName: grpc.example.com
    token: ""
    timeout: 15s
```

Supply the API token through a protected secret source.

Use a custom `caCert` only when necessary for your certificate infrastructure.

## Related documentation

- [Contexts](./contexts.md)
- [Authentication](./authentication.md)
- [Configuration Precedence](./configuration-precedence.md)
