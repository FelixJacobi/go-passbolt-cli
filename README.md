# go-pbolt-cli

<p align="center">
  <img src="https://raw.githubusercontent.com/pbolt-dev/.github/main/assets/pbolt-logo.svg" alt="pbolt logo" width="140" height="140" />
</p>

<p align="center">
  <strong>Fast, automation-friendly CLI tool and credential helper compatible with Passbolt.</strong>
</p>

<p align="center">
  <a href="https://github.com/pbolt-dev/go-pbolt-cli/releases"><img src="https://img.shields.io/github/v/release/pbolt-dev/go-pbolt-cli?style=flat-square&color=0bddec" alt="Latest Release" /></a>
  <a href="https://github.com/pbolt-dev/go-pbolt"><img src="https://img.shields.io/badge/Go%20SDK-go--pbolt-199ee9?style=flat-square&logo=go" alt="go-pbolt SDK" /></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-fdae15?style=flat-square" alt="License: MIT" /></a>
</p>

---

## ⚡ Overview

**pbolt** (`go-pbolt-cli`) is a community-driven CLI tool designed to interact with [Passbolt](https://www.passbolt.com), an open-source password manager for teams.

It is built for developers, system administrators, and CI/CD automation pipelines that need reliable, scriptable secret management without compromising cryptographic security.

If you are developing Go applications and need an API client, check out our SDK: [**`go-pbolt`**](https://github.com/pbolt-dev/go-pbolt).

---

## 🔒 API Compatibility

`go-pbolt-cli` is fully wire-compatible with standard Passbolt server instances (Community and Pro).

> [!IMPORTANT]
> Compatibility is continuously tested against Passbolt API 5.6.1 and later versions.

---

## 📦 Installation

### Download Pre-built Binary
Download the archive or deb/rpm package for Linux, macOS, or Windows from our [Releases page](https://github.com/pbolt-dev/go-pbolt-cli/releases/latest).

For Debian/Ubuntu:
```bash
sudo dpkg -i pbolt_*_linux_amd64.deb
```

### Via Go Install
```bash
go install github.com/pbolt-dev/go-pbolt-cli@latest
```
*(Note: Installing via Go installs the binary using the Go module path. Release packages include tab completions and manpages).*

---

## 🚀 Getting Started

To get started, configure your server connection, private key, and credentials:

```bash
pbolt configure --serverAddress https://passbolt.example.org --userPassword 'your-passphrase' --userPrivateKeyFile 'keys/privatekey.asc'
```

Alternatively, you can supply your private key directly via `--userPrivateKey`:
```bash
pbolt configure --serverAddress https://passbolt.example.org --userPassword 'your-passphrase' --userPrivateKey '-----BEGIN PGP PRIVATE KEY BLOCK-----...'
```

### Configuration Files & Locations
`pbolt configure` prioritizes modern XDG configuration directories:
- **Primary:** `$XDG_CONFIG_HOME/pbolt/pbolt.toml` (`~/.config/pbolt/pbolt.toml`)
- **Backward compatibility:** Existing configurations in `~/.config/go-passbolt-cli/go-passbolt-cli.toml` are automatically discovered and supported.

The configuration file is maintained with strict `0600` permissions (readable only by you) in a `0700` directory.

---

## 💡 Usage

The CLI follows a standard command structure:

```bash
pbolt <action> <entity> [arguments]
```

- `<action>`: `create`, `get`, `list`, `update`, `delete`, `share`, etc.
- `<entity>`: `resource` (passwords/secrets), `folder`, `user`, `group`.

### Examples

**Create a password resource:**
```bash
pbolt create resource --name "Production Database" --password "super-secret-password"
```

**List users:**
```bash
pbolt list user --column id,username,role
```

**Share a password:**
```bash
pbolt share resource --id <RESOURCE_ID> --type 1 --user <USER_ID>
```
*(Permission types: `1` = Read-only, `7` = Update, `15` = Owner, `-1` = Revoke).*

---

## 🛡️ Multi-Factor Authentication (MFA)

TOTP MFA is supported across three modes:

| Mode | Description |
| :--- | :--- |
| `none` | Fails immediately if MFA challenge is issued by server. |
| `interactive-totp` *(Default)* | Interactively prompts for TOTP code in your terminal. |
| `noninteractive-totp` | Automatically generates TOTP codes using the `--mfaTotpToken` secret. |

---

## 🤖 Automation & Scripting

- **JSON Output:** Use `-j` or `--json` with `get`, `list`, and `create` for structured machine-readable output.
- **Environment Secret Injection (`exec`):** Injects secrets directly into child processes without writing to shell history or disk:
  ```bash
  export GITHUB_TOKEN=passbolt://<RESOURCE_ID>
  pbolt exec -- gh auth login
  ```

---

## ⚖️ Trademark Notice

*Passbolt is a registered trademark of Passbolt SA. **pbolt** is an independent community project and is not affiliated with, endorsed by, or sponsored by Passbolt SA.*

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
