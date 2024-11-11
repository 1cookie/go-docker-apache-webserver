# Scaffold for a Go App on `Apache` Web Server Using Docker

## Overview

This scaffold demonstrates how to set up a Go application behind an `Apache` web server using Docker. The setup configures
`Apache` as a reverse proxy, directing client requests to the Go app on a separate port.

## Architecture

`Apache` runs on port `80`, the default `HTTP` port. Clients connect here, allowing seamless access via http://localhost/.
Go app runs on port `8080`, but it is only accessible to `Apache`, which forwards requests to this port using reverse proxy
settings (ProxyPass and ProxyPassReverse).
When a client makes a request, it hits `Apache` on port `80`. `Apache` acts as a middle layer, routing specific requests to
the Go app on port 8080 without clients directly accessing it. This configuration provides flexibility and security by
exposing only the `Apache` server to external requests.

## Purpose of the Reverse Proxy

`Apache`’s reverse proxy configuration in this setup achieves several goals:

**Control Access and Routing:** By exposing only port 80, `Apache` can restrict direct access to other services (e.g., the Go
app or database).

**Serve Multiple Applications:** `Apache` can proxy requests to multiple services, centralizing requests under a single domain
and port.

**Centralized Logging and Security:** Requests go through `Apache`, so logs and security rules are applied to all services
proxied through it.

## Running the App - `Makefile`

    make up

## How to Use Each Command
`make up`: Starts all services in detached mode (running in the background).
`make up-with-logs`: Starts all services and follows their logs.
`make build`: Builds or rebuilds the Docker images.
`make build-no-cache`: Rebuilds the images without cache, useful if you want a clean rebuild.
`make down`: Stops and removes containers.
`make logs`: Shows logs for all services.
`make restart`: Restarts the services.
`make configtest`: Runs an `Apache` configuration test inside the `Apache` container to check syntax.
`make clean`: Stops containers and removes all images, volumes, and orphaned containers, performing a complete cleanup.
`make logs-<service_name>`: Allows you to view logs for a specific service. For example, make logs-`Apache` or make logs-mysql.

Then visit: http://localhost/api/users or http://localhost/api/some-route

## Useful Commands

To view logs for each service:

```
docker-compose logs mysql
docker-compose logs apache
docker-compose logs myapp
```

## To access the Go app container:

    docker-compose exec myapp sh