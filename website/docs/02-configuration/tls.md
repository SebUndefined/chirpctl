
---
title: TLS and Certificates
sidebar_position: 3
description: Configure encrypted gRPC connections, custom certificate authorities, TLS verification and mutual TLS.
---

# TLS and Certificates

chirpctl communicates with ChirpStack using gRPC and supports several transport security configurations.

The recommended mode is **TLS with server certificate verification**, which protects both the connection and the identity of the remote server.

Other modes are available for private infrastructure, development environments and advanced deployments.

## Available connection modes

| Mode | Encrypted | Server verified | Typical use |
| --- | --- | --- | --- |
| Standard TLS | Yes | Yes | Public and production servers |
| TLS with custom CA | Yes | Yes | Private PKI and self-signed certificates |
| TLS with verification disabled | Yes | No | Testing and diagnostics |
| Plaintext gRPC | No | No | Local services and trusted tunnels |
| Mutual TLS (mTLS) | Yes | Yes* | Client certificate authentication |

*The server is verified when `tlsSkipVerify` is disabled and the appropriate CA is trusted.*

## Standard TLS

TLS is enabled by default when `insecure` is set to `false`.

```yaml
contexts:
  production:
    addr: grpc.example.com:443
    apiVersion: v4
    insecure: false
    tlsSkipVerify: false
    serverName: grpc.example.com
```

With this configuration, chirpctl:

1. Establishes an encrypted TLS connection.
2. Verifies the server's certificate chain.
3. Verifies that the certificate matches the expected server name.
4. Uses the operating system's trusted certificate authorities.

The `addr` field must contain a `host:port` address, not an HTTP URL.

## Using a custom CA certificate

Private ChirpStack installations may use certificates signed by an internal certificate authority.

Use `caCert` to add the appropriate CA certificate:

```yaml
contexts:
  private:
    addr: grpc.internal.example.com:443
    apiVersion: v4
    insecure: false
    tlsSkipVerify: false
    caCert: /etc/chirpctl/certs/internal-ca.pem
    serverName: grpc.internal.example.com
```

The certificate file must be PEM encoded.

With the current certificate-loading design, chirpctl adds certificates from `caCert` to the operating system's trusted certificate pool.

This allows both publicly trusted certificates and certificates issued by the configured private CA to be validated.

### Self-signed certificates

A self-signed certificate does not automatically require disabling TLS verification.

When the certificate can be used as an appropriate trust anchor, explicitly trust it through `caCert`.

The certificate must still satisfy the applicable TLS verification requirements, including the expected server identity.

:::tip

For production environments using an internal PKI, prefer `caCert` over `tlsSkipVerify`.

This preserves both encryption and server authentication.

:::

## TLS without certificate verification

chirpctl provides the `tlsSkipVerify` setting for environments where certificate verification must be bypassed temporarily.

```yaml
contexts:
  testing:
    addr: grpc-testing.example.com:443
    apiVersion: v4
    insecure: false
    tlsSkipVerify: true
```

You can also enable this behavior for an individual command:

```bash
chirpctl --tls-skip-verify get gateways
```

In this mode:

- TLS encryption remains enabled.
- The server certificate chain is not verified.
- The server's hostname is not verified.

:::danger

Disabling certificate verification weakens connection security and permits man-in-the-middle attacks.

Do not use `tlsSkipVerify` as a permanent workaround for an incorrectly configured production certificate.

:::

If `caCert` is also configured, it does not restore standard server verification while `tlsSkipVerify` is enabled.

## Plaintext gRPC

Set `insecure` to `true` to disable TLS entirely.

```yaml
contexts:
  local:
    addr: 127.0.0.1:1700
    apiVersion: v4
    insecure: true
```

Equivalent command-line usage:

```bash
chirpctl \
  --addr 127.0.0.1:1700 \
  --insecure \
  get gateways
```

:::important

`insecure: true` disables TLS encryption. It does not merely disable certificate verification.

Use `tlsSkipVerify` when TLS encryption is required but certificate verification must be bypassed.

:::

### Using plaintext gRPC through an SSH tunnel

A common deployment pattern exposes ChirpStack only through a private network.

For example:

```text
chirpctl
    |
    | Plaintext gRPC
    v
127.0.0.1:1700
    |
    | Encrypted SSH tunnel
    v
Bastion / private server
    |
    | Private network
    v
ChirpStack gRPC API
```

The local chirpctl connection uses plaintext gRPC, while the SSH tunnel encrypts its own network segment.

A suitable context is:

```yaml
contexts:
  private-tunnel:
    addr: 127.0.0.1:1700
    apiVersion: v4
    insecure: true
    token: ""
    timeout: 15s
```

With the tunnel established, provide the API token through an environment variable:

```bash
export CHIPCTL_TOKEN="YOUR_API_TOKEN"

chirpctl --context private-tunnel get gateways
```

chirpctl allows Bearer authentication over explicitly selected plaintext gRPC connections.

:::warning

An SSH tunnel protects traffic only along the tunnel itself.

If the connection continues from the SSH endpoint to ChirpStack over another network segment, that segment must also be appropriately secured.

Avoid exposing plaintext gRPC endpoints to untrusted networks.

:::

## TLS server name

The `serverName` property controls the server identity expected during TLS verification.

```yaml
addr: 192.0.2.10:443
serverName: grpc.example.com
insecure: false
```

This allows the TCP connection to target an IP address while validating the certificate against a DNS hostname.

When `serverName` is omitted, chirpctl derives it from the host portion of `addr`.

The server name is also relevant to TLS Server Name Indication (SNI) when a DNS name is used.

## Mutual TLS (mTLS)

Mutual TLS allows the server to authenticate the client using a client certificate.

It provides an additional authentication mechanism at the transport layer.

A mutual TLS configuration looks like this:

```yaml
contexts:
  secure:
    addr: grpc.secure.example.com:443
    apiVersion: v4
    insecure: false
    tlsSkipVerify: false
    caCert: /etc/chirpctl/certs/internal-ca.pem
    serverName: grpc.secure.example.com

    mTls:
      enabled: true
      clientCert: /etc/chirpctl/certs/client.crt
      clientKey: /etc/chirpctl/certs/client.key
```

When mTLS is enabled, chirpctl loads the configured certificate and private key and makes the certificate available during the TLS handshake.

The certificate and private key must form a valid pair.

### mTLS command-line options

When the corresponding flags are registered, the equivalent configuration is:

```bash
chirpctl \
  --mtls \
  --client-cert /etc/chirpctl/certs/client.crt \
  --client-key /etc/chirpctl/certs/client.key \
  get gateways
```

### mTLS and API tokens

mTLS and ChirpStack API-token authentication serve different purposes.

- **mTLS** authenticates the client at the transport layer.
- **API tokens** authenticate requests to ChirpStack.

Depending on the deployment, both may be required.

:::note

The server or reverse proxy must be configured to request or require client certificates.

Enabling mTLS in chirpctl does not make the server require client authentication.

:::

## TLS configuration reference

| YAML property | CLI flag | Description |
| --- | --- | --- |
| `insecure` | `--insecure` | Use plaintext gRPC |
| `tlsSkipVerify` | `--tls-skip-verify` | Disable TLS server certificate verification |
| `caCert` | `--cacert` | Additional trusted CA certificate bundle |
| `serverName` | `--server-name` | Expected TLS server name |
| `mTls.enabled` | `--mtls` | Enable mTLS client authentication |
| `mTls.clientCert` | `--client-cert` | Client certificate path |
| `mTls.clientKey` | `--client-key` | Client private key path |

## Security recommendations

For production deployments:

1. Prefer TLS with certificate verification.
2. Use a trusted public CA or an explicitly configured private CA.
3. Avoid `tlsSkipVerify` except for controlled diagnostic situations.
4. Restrict access to private keys and API tokens.
5. Use plaintext gRPC only when the complete network path is appropriately protected.
6. Prefer TLS 1.2 or newer.
7. Use mTLS when your infrastructure requires client-certificate authentication.

## Troubleshooting

### Certificate signed by unknown authority

The server certificate cannot be validated against the trusted certificate authorities.

Check the certificate chain and the configured `caCert` file.

### Certificate hostname mismatch

The server certificate does not match the configured TLS server name.

Verify the hostname, `serverName`, and certificate Subject Alternative Name entries.

### TLS handshake failure

Check:

- Whether the endpoint actually supports TLS.
- Whether the expected gRPC port is used.
- Whether the reverse proxy supports gRPC and HTTP/2.
- Whether the appropriate certificates are configured.

### Failed mTLS authentication

Verify:

- The client certificate and private key match.
- The certificate is valid and not expired.
- The server trusts the client certificate issuer.
- The reverse proxy or gRPC server is configured for mTLS.

### Plaintext connection rejected

Confirm whether the server expects TLS.

Do not enable plaintext mode against a TLS-only endpoint.

## Related documentation

- [Contexts](./contexts.md)
- [Authentication](./authentication.md)
- [Configuration Precedence](./configuration-precedence.md)
