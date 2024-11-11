# Makefile for Docker Compose management

# Variables
DOCKER_COMPOSE = docker-compose

# Default: Run the 'up' target
default: up

# Bring up services in detached mode
up:
	@$(DOCKER_COMPOSE) up -d
	@echo "Services are up and running."

# Bring up services with logs for live feedback
up-with-logs:
	@$(DOCKER_COMPOSE) up

# Build or rebuild services
build:
	@$(DOCKER_COMPOSE) build
	@echo "Build complete."

# Build without cache (useful for full rebuilds)
build-no-cache:
	@$(DOCKER_COMPOSE) build --no-cache
	@echo "Build complete without cache."

# Stop services
down:
	@$(DOCKER_COMPOSE) down
	@echo "Services stopped."

# View logs from all services
logs:
	@$(DOCKER_COMPOSE) logs -f

# Restart services
restart:
	@$(DOCKER_COMPOSE) restart
	@echo "Services restarted."

# Run Apache config test inside the container
configtest:
	@$(DOCKER_COMPOSE) run apache apachectl configtest

# Clean up (stop and remove all containers, images, volumes, and networks)
clean:
	@$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans
	@echo "Cleanup complete."

# Tail logs from a specific service (e.g., `make logs-apache`)
logs-%:
	@$(DOCKER_COMPOSE) logs -f $*

.PHONY: up up-with-logs build build-no-cache down logs restart configtest clean logs-%
