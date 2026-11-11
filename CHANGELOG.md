# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

Releases before this file was introduced are documented in the [GitHub releases](https://github.com/passbolt/go-passbolt-cli/releases).

## [Unreleased]

## [0.7.0] - 2026-11-11
### Changed
- Renamed Go module to `github.com/haspr-dev/go-haspr-cli`
- Updated SDK dependency to `github.com/haspr-dev/go-haspr v0.10.0`
- Adopted `haspr` CLI branding with primary executable `haspr` and helper executables `git-credential-haspr`, `docker-credential-haspr`
- Added backward compatibility aliases for `pbolt` and `passbolt` binaries and configuration directories

## [0.6.0] - 2026-11-10
### Added
- **Authentication Agent**: Opt-in in-memory authentication agent (`--agent`) reusing authenticated sessions behind interactive approval dialogs (desktop, terminal, pinentry) with idle expiry.
- **Native Credential Helpers**:
  - `git-credential-pbolt` (and `git-credential-passbolt`) Git credential helper matching HTTP(S) URLs against Passbolt resource URIs.
  - `docker-credential-pbolt` (and `docker-credential-passbolt`) Docker credential helper using upstream Docker credentials interface.
  - `credential-helper aws`: AWS `credential_process` integration.
  - `credential-helper k8s`: Kubernetes `client-go` exec authentication integration.
- **Raw Secret Output**: Added `--raw` flag and `--field` to `pbolt get resource` for pipeline integration without tables, JSON, or trailing newlines.
- **Organization Settings**: Added commands to manage plugin organization settings: `pbolt get organization-settings`, `update organization-settings`, and `delete organization-settings`.
- **User Management**:
  - Reversible user account state management via `pbolt update user --active=true/false`.
  - Added user recovery and invitation renewal commands (`pbolt recover user`).
  - Exposed account state, optional MFA status, and SCIM directory linkage in user details.
- **Group Management**:
  - Exposed complete membership UUIDs (`membership_id`) separately from user UUIDs in group details and `--column MembershipID`.

### Fixed
- Fixed group manager demotion when updating group memberships to prevent accidental manager privilege retention.
- Handled missing profile/role associations gracefully in user details.

### Changed
- Bumped `github.com/pbolt-dev/go-pbolt` dependency to `v0.9.0`.
- Enhanced GoReleaser configuration to build helper binaries (`git-credential-pbolt`, `docker-credential-pbolt`, and compatibility aliases).


## [0.5.4] - 2026-11-10
### Changed
- Renamed Go module to `github.com/pbolt-dev/go-pbolt-cli`
- Updated SDK dependency to `github.com/pbolt-dev/go-pbolt v0.8.5`
- Adopted `pbolt` CLI branding and configuration compatibility

### Maintenance
- ci: updated `golangci-lint` to `v2.15` in GitHub Actions for Go 1.27 compatibility
- ci: enabled Dependabot for Go modules and GitHub Actions workflows
- Security: enabled Dependabot vulnerability alerts and automated security fixes


## [0.5.3] - 2026-11-05
### Added
- PB-54610: Skip Resources this build cannot read in `list resource`
- PB-54610: Explain schema validation errors in `create`, `update`, `share` and `get`

### Fixed
- PB-54147: Propagate the real error when interactive TOTP is exhausted
- PB-56059: Fix the `--user` and `--manager` filters of `list group` and the `--group` filter of `list user`
- PB-56059: Fix the help text of the `list group` and `list user` filters

### Security
- PB-54765: Sanitize terminal output in every command to prevent escape sequence injection (Medium)
- PB-54765: Escape non-BMP characters as `\uXXXX` in JSON output
- PB-54766: Write the CLI config through a hardened file descriptor (Medium)
- PB-54766: Repair default config permissions and warn on custom paths
- PB-54766: Create the KeePass export with owner-only permissions
- PB-54766: Ask for confirmation before `configure --debug` prints the saved settings, and print them to stderr
- PB-54766: Document cleartext credential storage in the config file
- PB-54978: Restrict the config file and the KeePass export to their owner on Windows
- PB-54978: Document where the config file is unsafe to store
- PB-54987: Warn instead of refusing when a secret file stays readable by others or belongs to someone else
- PB-54147: Move to the Go 1.27.1 security patch
- PB-55389: Update Go to 1.27.2 (security fixes in crypto/tls, net/http, net/textproto and os)
- PB-55610: Bump Go version to 1.27.3 (security fixes in crypto/x509)

### Maintenance
- PB-54610: Bump the Go SDK to v0.8.4
- PB-54147: Bump Go to 1.27 (the Go SDK stays buildable on Go 1.26.10)
- PB-54147: Pin the omitempty semantics the `--json` output relies on
- PB-54765: Forbid direct fmt printing in golangci-lint
- PB-54765: Add integration tests to avoid terminal escape regression
- PB-54766: Add regression tests for config file permissions
- PB-54978: Move the secure-file code into util/securefile
- PB-54610: Run the txtar scenarios as benchmarks
- PB-54610: Cover every bundled resource type end to end
- PB-54610: Reorganize integration test runner
- PB-55287: Reorganize e2e test files in CLI and SDK
- PB-54998: Update GitHub Actions to stable versions
- PB-55888: Update Go dependencies
- PB-55888: Migrate cel-go to the cel.dev/cel-go module path
- PB-55888: Keep the previous table style after the pterm upgrade
- PB-56059: Add integration tests for the list filters
- Renovate: Update goreleaser/goreleaser-action action to v7.3.0
- Renovate: Update actions/upload-artifact action to v7.0.2

## [0.5.2] - 2026-08-28
### Fixed
- PB-53937: Sign shared v5 metadata with both the user key and the metadata key (Go SDK v0.8.3)

### Security
- PB-53335: Fix crypto/tls CVE-2026-42505 (Medium)
- PB-53334: Minor upgrade for golang.org/x/text (High)
- PB-53929: Minor upgrade for github.com/moby/go-archive (High)
- PB-53356: Minor upgrade for github.com/google/cel-go (Medium)

### Maintenance
- PB-53915: Use the SDK custom field helpers in the KeePass export (Go SDK v0.8.2)
- PB-53919: Bump go version to 1.26.6
- PB-54017: Update dependencies in SDK and CLI
- PB-54017: Replace a deprecated gokeepasslib call in the KeePass export
- PB-54022: Use go 1.25 syntax
- PB-53686: Document the minimum supported Passbolt API version
- PB-54031: Add changelog files to SDK and CLI
- PB-53914: Remove the community disclaimer from the README
- PB-53788: Update devcontainer version in go CI
- PB-53937: Sign the metadata private key payload with a timestamp when bootstrapping the test environment
- Renovate: Add renovate.json
- Renovate: Adapt renovate config
