# Dynamic Nginx Reverse Proxy Configuration Generator

A Docker-based solution that dynamically generates Nginx reverse proxy configurations from environment variables.

## Overview

This project provides a containerized solution that combines Nginx with a Go-based configuration generator. It automatically creates Nginx reverse proxy configurations based on environment variables, making it easy to set up multiple domain mappings.

## How It Works

1. The system uses a Go program to read environment variables in the format:
   ```
   APP_n=domain:host:port
   ```

2. For each environment variable, it generates an Nginx configuration file using the template, setting up:
   - Domain name (server_name)
   - Upstream service host
   - Port number
   - Standard proxy headers

3. All configurations are validated before Nginx starts

## Features

- Dynamic configuration generation
- Multiple domain support
- Secure default settings
- Automatic configuration validation
- Docker-optimized build process
- Alpine-based for minimal image size

## Usage

1. Build the Docker image:
   ```bash
   docker build -t nginx-rp .
   ```

2. Run the container with environment variables:
   ```bash
   docker run -d \
     -p 80:80 \
     -v /host_path:/etc/nginx/conf.d
     -v /tmp/logs:/var/log/nginx
     -e "APP_1=example.com:backend:8080" \
     -e "APP_2=api.example.com:api-service:3000" \
     nginx-rp
   ```

## File Structure

- `main.go` - Go configuration generator
- `nginx.conf.template` - Nginx configuration template
- `Dockerfile` - Multi-stage build definition
- `start.sh` - Container entrypoint script

## Requirements

- Docker
- Environment variables following the format: domain:host:port

## Security Features

- Runs Nginx as non-root user
- Includes secure proxy headers
- Proper error handling
- Graceful shutdown support