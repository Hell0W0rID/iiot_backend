#!/bin/bash

# Docker build and deployment script for IIoT Backend
# Resolves all Go 1.21 dependencies in containerized environment

set -e

echo "Building IIoT Backend with Go 1.21 in containers..."

# Build infrastructure services
echo "Starting infrastructure services..."
docker-compose up -d postgres redis consul mosquitto

# Wait for infrastructure to be ready
echo "Waiting for infrastructure services..."
sleep 10

# Build and start microservices
echo "Building and starting microservices..."
docker-compose build --no-cache
docker-compose up -d iiot-metadata iiot-data iiot-command iiot-gateway

# Show service status
echo "Service status:"
docker-compose ps

echo "Deployment complete! Services available at:"
echo "- Gateway: http://localhost:8000"
echo "- Metadata: http://localhost:8001"
echo "- Data: http://localhost:8002"
echo "- Command: http://localhost:8003"
echo "- Consul UI: http://localhost:8500"