.PHONY: test test-unit test-integration docker-up docker-down docker-restart docker-logs wait-for-ssh deps fmt lint ci clean

DEBUG ?= YES
INSECURE_IGNORE_HOST_KEY ?= YES
SSH_HOST ?= localhost
SSH_PORT ?= 2222

export DEBUG
export INSECURE_IGNORE_HOST_KEY
export SSH_HOST
export SSH_PORT

# Run all tests (unit + integration with fresh container)
test: test-unit docker-restart wait-for-ssh test-integration docker-down

# Unit tests only - packages that don't require SSH
test-unit:
	go test -v -race ./passwd/... ./group/... ./authorized_keys/...

# Integration tests - requires running SSH container
test-integration:
	go test -v -race .

# Docker management
docker-up:
	docker compose up -d

docker-down:
	docker compose down -v

docker-restart: docker-down
	docker compose up -d --build

wait-for-ssh:
	@for i in $$(seq 1 30); do \
		if nc -z $(SSH_HOST) $(SSH_PORT) 2>/dev/null; then \
			echo "SSH server is ready"; \
			exit 0; \
		fi; \
		echo "Waiting for SSH server... ($$i/30)"; \
		sleep 1; \
	done; \
	echo "Timeout waiting for SSH server"; \
	exit 1

docker-logs:
	docker compose logs -f

# Development tasks
deps:
	go mod download
	go mod tidy

fmt:
	go fmt ./...

lint:
	go vet ./...

clean:
	docker compose down -v 2>/dev/null || true

# CI pipeline - runs everything
ci: deps fmt lint docker-restart wait-for-ssh test
