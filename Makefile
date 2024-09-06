include .env

# In case you want something else (e.g. podman)
OCI ?= docker

# Configuration -------------------------------------------------->
.PHONY: default-config
default-config: .env

# Create new configuration file (.env) from example (.example.env)
.env: .example.env
	cp $< $@

# Docker -------------------------------------------------->
.PHONY: docker-up
docker-up: docker-build docker-run

.PHONY: docker-build
docker-build:
	${OCI} build --tag users-service:${USERS_APP__VERSION} .

.PHONY: docker-run
docker-run:
	${OCI} run \
		--rm \
		--detach \
		--name users-service-c \
		--publish ${USERS_HTTP__PORT}:${USERS_HTTP__PORT} \
		--env-file .env \
		--health-cmd "curl --fail http://localhost:${USERS_HTTP__PORT}/healthcheck || exit 1" \
			--health-start-period 5s \
			--health-interval 30s \
			--health-timeout 3s \
			--health-retries 3 \
		users-service:${USERS_APP__VERSION}

.PHONY: docker-stop
docker-stop:
	${OCI} container stop users-service-c

.PHONY: docker-clean
docker-clean:
	${OCI} image rm users-service:${USERS_APP__VERSION}

# Dev -------------------------------------------------->
# Run server instance with hot-reaload
.PHONY: run-dev
run-dev:
	air

.PHONY: setup-dev
setup-dev: default-config install-pre-commit

# Install pre-commit hooks
.PHONY: install-pre-commit
install-pre-commit:
	# You'll need pre-commit installed (check https://pre-commit.com/ for instructions)
	pre-commit install
