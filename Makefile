# Configuration -------------------------------------------------->
.PHONY: default-config
default-config: .env

# Create new configuration file (.env) from example (.example.env)
.env: .example.env
	cp $< $@

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
