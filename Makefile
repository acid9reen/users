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
USERS_SERVICE_CONTAINER__NETWORK_NAME ?= users-service
USERS_SERVICE_CONTAINER__IMAGE_NAME ?= users-service
USERS_SERVICE_CONTAINER__CONTAINER_NAME ?= users-service-c

.PHONY: docker-up
docker-up: docker-build docker-create-network docker-run

.PHONY: docker-build
docker-build:
	${OCI} build --tag ${USERS_SERVICE_CONTAINER__IMAGE_NAME}:${USERS_APP__VERSION} .

.PHONY: docker-run
docker-run:
	${OCI} run \
		--rm \
		--detach \
		--name ${USERS_SERVICE_CONTAINER__CONTAINER_NAME} \
		--network ${USERS_SERVICE_CONTAINER__NETWORK_NAME} \
		--publish ${USERS_HTTP__PORT}:${USERS_HTTP__PORT} \
		--env-file .env \
		--health-cmd "curl --fail http://localhost:${USERS_HTTP__PORT}/healthcheck || exit 1" \
			--health-start-period 5s \
			--health-interval 30s \
			--health-timeout 3s \
			--health-retries 3 \
		${USERS_SERVICE_CONTAINER__IMAGE_NAME}:${USERS_APP__VERSION}

.PHONY: docker-stop
docker-stop:
	${OCI} container stop ${USERS_SERVICE_CONTAINER__CONTAINER_NAME}

.PHONY: docker-clean
docker-clean:
	-${OCI} image rm ${USERS_SERVICE_CONTAINER__IMAGE_NAME}:${USERS_APP__VERSION}
	-${OCI} network rm ${USERS_SERVICE_CONTAINER__NETWORK_NAME}

.PHONY: docker-create-network
docker-create-network:
	-${OCI} network create --driver bridge ${USERS_SERVICE_CONTAINER__NETWORK_NAME}

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
