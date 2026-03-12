# SSH Key Manager - TODO

This file tracks ongoing tasks and future improvements for the SSH Key Manager project.

## In Progress

### CI/CD & Docker Improvements
- [ ] Replace `INSECURE_IGNORE_HOST_KEY` with `ssh-keyscan` for proper host key verification in CI
  - [ ] Add step to run `ssh-keyscan -p 2222 localhost` and save to known_hosts
  - [ ] Update CI workflow to create `~/.ssh` directory if needed
  - [ ] Remove `INSECURE_IGNORE_HOST_KEY` env var from tests
  - [ ] Test that host key verification works correctly

- [ ] Set up Docker image build and push to registry
  - [ ] Choose registry (Docker Hub, GitHub Container Registry, etc.)
  - [ ] Add Dockerfile for the service (separate from test Dockerfile)
  - [ ] Add GitHub Actions workflow for building and pushing images
  - [ ] Tag images with version and latest
  - [ ] Set up multi-arch builds (amd64, arm64)

## Pending

### Documentation
- [ ] CLI Documentation (`cmd/sshkeymanager`)
  - [ ] Installation instructions
  - [ ] Configuration (environment variables, config file)
  - [ ] Command-line flags and options
  - [ ] Usage examples
  - [ ] SSH key formats supported

- [ ] Daemon Documentation (`cmd/sshkeyserver`)
  - [ ] Architecture overview
  - [ ] Installation and setup
  - [ ] Configuration options
  - [ ] Running as a systemd service
  - [ ] Logging and monitoring
  - [ ] Security considerations

- [ ] HTTP API Documentation
  - [ ] API endpoints reference
  - [ ] Request/response formats
  - [ ] Authentication methods
  - [ ] Error codes and handling
  - [ ] Rate limiting
  - [ ] OpenAPI/Swagger spec

- [ ] Docker Service Documentation
  - [ ] Docker Compose setup
  - [ ] Environment variables
  - [ ] Volume mounts
  - [ ] Network configuration
  - [ ] Health checks
  - [ ] Production deployment guide

- [ ] Development Documentation
  - [ ] Project structure
  - [ ] Running tests locally
  - [ ] Contributing guidelines
  - [ ] Code style guide

### Testing
- [ ] Increase test coverage
- [ ] Add benchmarks
- [ ] Add integration tests for API endpoints
- [ ] Test with different SSH server configurations

### Features
- [ ] Add support for ed25519 keys
- [ ] Add support for SSH certificates
- [ ] Implement key rotation functionality
- [ ] Add audit logging
- [ ] Add metrics and monitoring endpoints

### Security
- [ ] Security audit
- [ ] Add rate limiting
- [ ] Implement proper secret management
- [ ] Add mTLS support for API
- [ ] Review and update dependencies regularly

## Completed

- [x] Update to Go 1.26.1
- [x] Update all legacy packages
- [x] Set up testing against a real SSH server using Docker
- [x] Create Makefile with test targets
- [x] Create GitHub Actions CI workflow
- [x] Fix tests for fresh Debian bookworm (23 users instead of 25)
- [x] Update docker-compose to v2
