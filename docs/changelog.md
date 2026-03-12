# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- GitHub Actions CI workflow for automated testing
- Makefile with comprehensive test and docker targets
- Docker Compose setup for SSH test server
- Integration tests against real SSH server
- Documentation structure (docs/ directory)
- This changelog file

### Changed
- **BREAKING**: Updated Go version from 1.17 to 1.26.1
- Updated all dependencies to latest versions
- Migrated from deprecated `github.com/go-yaml/yaml` to `gopkg.in/yaml.v3`
- Updated Docker base image from `debian:buster` to `debian:bookworm-slim`
- Updated docker-compose commands from v1 (`docker-compose`) to v2 (`docker compose`)

### Fixed
- Fixed user count in tests for fresh Debian bookworm (23 users instead of 25)
- Fixed SSH host key verification in CI by using `INSECURE_IGNORE_HOST_KEY` environment variable
- Fixed container lifecycle management in testserver to check if already running

### Security
- Added host key callback configuration options
- Added support for `SSH_HOST_KEY` environment variable for fixed host keys
- Added `INSECURE_IGNORE_HOST_KEY` option for testing environments (use with caution)

## [0.1.0] - Previous Version

### Added
- Initial implementation of SSH Key Manager
- CLI tool for managing SSH keys
- HTTP API server
- Support for managing `/etc/passwd`, `/etc/group`, and `~/.ssh/authorized_keys`
- Basic test suite

### Notes
- This version used Go 1.17
- Dependencies were outdated
- Limited test coverage
