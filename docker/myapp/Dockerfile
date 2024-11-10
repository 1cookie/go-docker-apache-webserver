# Build stage for the Go app
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o myapp ./cmd/myapp

# Run stage with Apache
FROM httpd:latest
COPY --from=builder /app/myapp /usr/local/bin/myapp

# Set up Apache as a reverse proxy and static file server
COPY ./configs/apache.local.conf /usr/local/apache2/conf/httpd.conf
COPY ./web/static /usr/local/apache2/htdocs/

# Ensure permissions are correct
RUN chmod +x /usr/local/bin/myapp

# Expose Go's port internally, but only expose port 80 externally
EXPOSE 80
CMD ["sh", "-c", "/usr/local/bin/myapp & httpd-foreground"]
