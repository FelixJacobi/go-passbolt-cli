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

### Interactive authentication agent (opt-in)

Use `pbolt --agent list user` to start an in-memory authentication agent. It
asks for approval for each command, and asks for your passphrase and TOTP when
needed. Later commands reuse the authenticated session. Authentication expires
after one hour without approved use; active use resets that idle timer. Server
session expiry and MFA policy still apply.

To enable this for your configuration, add these settings to your TOML file:

```toml
agent = true
agent-dialog = "auto"
```

Keep `userPassword` and your MFA seed out of the file to enter them interactively.
`--agent=false` uses ordinary CLI authentication for one invocation.

The default dialog backend uses Zenity or KDialog under X.org/Wayland, system
dialogs on macOS and Windows, or a full-terminal dialog when no desktop helper
is available. Install `zenity` or `kdialog` for Linux desktop dialogs. Select
`--agent-dialog=terminal` to always use the terminal, or
`--agent-dialog=pinentry --agent-pinentry=/path/to/pinentry` to use GnuPG's
pinentry helper. An explicit denial never falls back to another dialog.
Pipelines work: approval uses the desktop or controlling terminal, leaving the
command's stdin and stdout available for data. Without either UI, agent access
fails. `--yes` does not bypass approval.

```sh
pbolt --agent list user --json | jq .
pbolt agent status
pbolt agent lock   # clear the session matching this configuration
pbolt agent stop   # clear all cached sessions and stop the process
```

The agent runs Passbolt operations itself and keeps private keys and session
cookies inside its process. `exec` receives only the requested resource values
and launches your program in the foreground, preserving its terminal. Cached
sessions are isolated by server, account key and authentication/transport policy.
The agent never writes cached authentication to disk or a credential store.
`--debug` is unavailable in agent mode because SDK debug logging exposes secrets.
After upgrading the CLI, stop an older running agent before using the new binary.
See [SECURITY.md](SECURITY.md) for the local security boundary.

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

### Enable or Disable a User

Administrators can disable an account without deleting it, or enable it again:

```bash
pbolt update user --id USER_UUID --disabled=true
pbolt update user --id USER_UUID --disabled=false
```

Omitting `--disabled` leaves the account state unchanged. Account state changes
cannot be combined with `--firstname`, `--lastname`, or `--role` in one command.

---

## 🤖 Automation & Scripting

- **JSON Output:** Use `-j` or `--json` with `get`, `list`, and `create` for structured machine-readable output.
- **Environment Secret Injection (`exec`):** Injects secrets directly into child processes without writing to shell history or disk:
  ```bash
  export GITHUB_TOKEN=passbolt://<RESOURCE_ID>
  pbolt exec -- gh auth login
  ```

### Read optional user MFA and directory details

User detail JSON preserves `is_mfa_enabled` when the server discloses it. Missing
MFA data is omitted; an explicit `false` remains false. Administrators can request
SCIM directory linkage in the server's original JSON schema:

```sh
pbolt get user --id USER_UUID --scim-entry --json
pbolt list user --json --scim-entry
pbolt list user --json --include-mfa
pbolt list user --mfa-enabled=false --json
pbolt list user --column Username --column IsMFAEnabled
```

MFA filtering is tri-state: omitting `--mfa-enabled` leaves the filter unset.
Explicit true and false both request the MFA data required by the server filter.
Selecting the MFA column also requests its data. For a CEL expression on
`is_mfa_enabled`, use `--include-mfa`; unavailable values are null/unknown rather
than false. Ordinary list requests do not request these administrator-only
associations. The server controls field visibility and plugin availability.
SCIM linkage appears in list JSON; this does not expose custom SSO claim mappings
or server-internal directory records.

Optional profile,role and timestamp columns are null in CEL expressions when
not returned by the server, and render as blank table cells instead of panicking.

### Request recovery mail or renew an invitation

Using your configured connection, request account recovery mail by username:

```sh
pbolt recover user --username ada@example.com
```

For an account that has not completed setup, this renews the invitation.
Success does not confirm that an account exists or that an email was delivered.

### Change a group member's manager role

Promote an existing member to manager, or demote a manager while retaining membership:

```sh
pbolt update group --id GROUP_UUID --manager USER_UUID
pbolt update group --id GROUP_UUID --user USER_UUID
```

Demotion requires the SDK to send an explicit `is_admin: false` membership update.
The command retains the normal dry-run and secret re-encryption workflow.

Group detail JSON exposes each user's `membership_id` separately from the existing
`id` user UUID. The table accepts `--column MembershipID`. Group updates are
idempotent and reread the group to verify requested names and memberships; shared
secrets still use the SDK's normal dry-run and re-encryption flow. A verification
error may follow a partially applied server change, so reread before retrying.

### Manage organization settings

Read a plugin's organization settings as JSON:

```sh
pbolt get organization-settings --plugin email-notifications
```

Submit a plugin-specific JSON object from a file or stdin:

```sh
pbolt update organization-settings --plugin email-notifications --input-file notifications.json
pbolt update organization-settings --plugin email-notifications --input-file - < notifications.json
```

Updates are silent by default. Add `--json` to print the server's response.
Input uses the plugin API's original JSON schema, without field renaming or
notification-key flattening. Read responses may contain server-only fields and
are not necessarily suitable for direct reuse as update input. Prepare any
required encrypted or signed credentials before submitting them. SSO writes
create settings; activation is a separate server workflow. Metadata settings
writes do not perform metadata migrations.

Delete or disable settings where supported. Some plugins have one settings
record; others require a settings UUID returned by their API:

```sh
pbolt delete organization-settings --plugin directory-sync
pbolt delete organization-settings --plugin self-registration --id SETTINGS_UUID
```

The client supports these operations; availability and permissions also depend
on the server's edition and enabled plugins. Plugin names support shell
completion, filtered by the command's supported operations.

| Plugin | Get | Update | Delete |
| --- | --- | --- | --- |
| account-recovery-policies | yes | yes | — |
| directory-sync | yes | yes | no ID |
| email-notifications | yes | yes | — |
| export-policies | yes | — | — |
| locale | — | yes | — |
| metadata-keys | yes | yes | — |
| metadata-types | yes | yes | — |
| mfa | yes | yes | — |
| mfa-policies | yes | yes | — |
| password-expiry | yes | yes | requires `--id` |
| password-policies | yes | yes | — |
| scim | yes | yes | requires `--id` |
| secret-revisions | yes | yes | no ID |
| self-registration | yes | yes | requires `--id` |
| smtp | yes | yes | — |
| sso | yes | yes | requires `--id` |
| user-key-policies | yes | — | — |
| user-passphrase-policies | yes | yes | — |

---

## ⚖️ Trademark Notice

*Passbolt is a registered trademark of Passbolt SA. **pbolt** is an independent community project and is not affiliated with, endorsed by, or sponsored by Passbolt SA.*

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
