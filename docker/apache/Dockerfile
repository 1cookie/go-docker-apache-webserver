# Use official Apache image
FROM httpd:latest

# Ensure log directories exist and have correct permissions
RUN mkdir -p /usr/local/apache2/logs /tmp/logs && \
    chown -R www-data:www-data /usr/local/apache2/logs /tmp/logs && \
    chmod -R 755 /usr/local/apache2/logs /tmp/logs

# Copy the custom Apache configuration file
COPY ./configs/apache.local.conf /usr/local/apache2/conf/httpd.conf

# Expose Apache's default port
EXPOSE 80
