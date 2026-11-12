---
title: Installation
sidebar_position: 1
description: Install chirpctl on Linux, macOS, or Windows.
---

# Installation

chirpctl is distributed as a standalone Go executable for Linux, macOS, and Windows.

You do not need to install Go to use a precompiled release.

## Download chirpctl

Download the latest archive from the GitHub Releases page:

**[Download the latest release](https://github.com/sebundefined/chirpctl/releases/latest)**

Choose the archive that matches your operating system and CPU architecture.

| Operating system | Architecture | Archive |
| --- | --- | --- |
| Linux | AMD64 | `linux_amd64.tar.gz` |
| Linux | ARM64 | `linux_arm64.tar.gz` |
| macOS | Intel | `darwin_amd64.tar.gz` |
| macOS | Apple Silicon | `darwin_arm64.tar.gz` |
| Windows | AMD64 | `windows_amd64.zip` |
| Windows | ARM64 | `windows_arm64.zip` |

Archive filenames follow this pattern:

```text
chirpctl_<version>_<os>_<arch>.<extension>
```

For example:

```text
chirpctl_0.1.0_darwin_arm64.tar.gz
```

The version number may differ from the example above.

## Linux

### 1. Extract the archive

Download the appropriate Linux archive, then extract it:

```bash
tar -xzf chirpctl_<version>_linux_amd64.tar.gz
```

For an ARM64 system, use the corresponding `linux_arm64` archive.

### 2. Install the executable

Install chirpctl into your local user binary directory:

```bash
mkdir -p "$HOME/.local/bin"

install -m 755 chirpctl "$HOME/.local/bin/chirpctl"
```

Ensure that `$HOME/.local/bin` is included in your `PATH`.

For example, with Bash:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

Add this export to your shell configuration if necessary to make it persistent.

### 3. Verify the installation

```bash
chirpctl --version
```

You should see the installed chirpctl version.

## macOS

chirpctl provides native executables for both Intel and Apple Silicon Macs.

### 1. Choose the correct architecture

- Apple Silicon (M1, M2, M3, and later): `darwin_arm64`
- Intel Mac: `darwin_amd64`

You can check your architecture using:

```bash
uname -m
```

`arm64` indicates Apple Silicon. `x86_64` indicates Intel.

### 2. Extract and install

Extract the downloaded archive:

```bash
tar -xzf chirpctl_<version>_darwin_arm64.tar.gz
```

Install the executable:

```bash
mkdir -p "$HOME/.local/bin"

install -m 755 chirpctl "$HOME/.local/bin/chirpctl"
```

Ensure that `$HOME/.local/bin` is included in your shell's `PATH`.

### 3. Verify the installation

```bash
chirpctl --version
```

### macOS security warning

macOS may prevent an unsigned executable downloaded from the internet from opening.

If macOS blocks chirpctl, first verify that you downloaded the correct official release and checked its checksum.

Then, if you trust the executable, use the approval process under **System Settings → Privacy & Security**.

:::warning

The current distribution should not be assumed to be Apple-notarized. Do not disable macOS security protections globally to install chirpctl.

:::

## Windows

### 1. Download the archive

Choose the appropriate Windows ZIP archive from GitHub Releases.

For most Intel or AMD Windows computers, select:

```text
chirpctl_<version>_windows_amd64.zip
```

### 2. Extract the executable

Extract the ZIP archive into a directory of your choice, for example:

```text
C:\Tools\chirpctl\
```

The directory should contain:

```text
chirpctl.exe
```

### 3. Add chirpctl to PATH

To use chirpctl from any terminal:

1. Open Windows Environment Variables.
2. Edit the user `Path` variable.
3. Add `C:\Tools\chirpctl`.
4. Open a new PowerShell terminal.

### 4. Verify the installation

```powershell
chirpctl.exe --version
```

If you have not added the executable to your `PATH`, run it from its directory:

```powershell
.\chirpctl.exe --version
```

## Verify downloaded archives

Each release includes a SHA-256 checksum file.

Before installing, compare your downloaded archive's checksum with the entry provided in the release.

On Linux:

```bash
sha256sum chirpctl_<version>_linux_amd64.tar.gz
```

On macOS:

```bash
shasum -a 256 chirpctl_<version>_darwin_arm64.tar.gz
```

On Windows PowerShell:

```powershell
Get-FileHash .\chirpctl_<version>_windows_amd64.zip -Algorithm SHA256
```

The calculated checksum must match the corresponding value in the release checksum file.

## Updating chirpctl

To update, download the latest release for your platform and replace the existing executable.

Your chirpctl configuration is stored separately from the executable and does not need to be recreated for a normal binary update.

Check the release notes before upgrading, especially when command behavior or configuration formats change.

## Next step

Your installation is ready.

Continue with the **[Quick Start](./quickstart.md)** to configure your first ChirpStack connection.
