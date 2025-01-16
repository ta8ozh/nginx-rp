#!/bin/bash
set -eo pipefail  # Exit on error, undefined var, and pipeline failures

echo "Generating Nginx configuration..."
./conf_generator

echo "Checking Nginx configuration..."
if ! nginx -t; then
    echo "Nginx configuration is incorrect. Exiting..." >&2
    exit 1
fi

echo -e "\nNginx configuration:"
nginx -T
echo

# Start Nginx in foreground
echo "Running Nginx in the foreground..."
exec nginx -g "daemon off;"
