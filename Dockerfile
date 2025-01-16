# Use the official Golang image to build the application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary file
COPY main.go .

# Compile the Go application with optimization flags
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o conf_generator main.go

# Use a minimal nginx image for the final container
FROM nginx:stable-alpine

# Copy only the required files from the builder stage
COPY --from=builder /app/conf_generator /conf_generator
COPY nginx.conf.template .
COPY ./start.sh /start.sh

# Configure permissions and clean up in a single layer
RUN chmod +x /start.sh && \
    chown -R nginx:nginx /conf_generator /start.sh

# Expose the default Nginx port
EXPOSE 80

# Set the default stop signal
STOPSIGNAL SIGQUIT

# Run the custom script as the entrypoint
ENTRYPOINT ["/start.sh"]
