# IIoT Backend - Docker Microservices Deployment Guide

## Current Setup Status

**Go Version:** 1.19.3 (compatible with container environment)

**Architecture:** Containerized microservices with shared infrastructure

## Service Architecture

### Infrastructure Services
- **PostgreSQL 15** - Primary database (port 5432)
- **Redis 7** - Caching and message queue (port 6379)
- **Consul 1.16** - Service discovery and configuration (port 8500)
- **Mosquitto MQTT** - IoT device communication (ports 1883, 9001)

### Application Services
- **iiot-gateway** - API Gateway and load balancer (port 8000)
- **iiot-metadata** - Device management service (port 8001)
- **iiot-data** - Data collection service (port 8002)
- **iiot-command** - Device command service (port 8003)

## Deployment Commands

### Start Infrastructure Only
```bash
docker-compose up -d postgres redis consul mosquitto
```

### Start All Services
```bash
docker-compose up -d
```

### Start Specific Service
```bash
docker-compose up -d iiot-metadata
```

### Build and Start
```bash
docker-compose up --build
```

### View Service Logs
```bash
docker-compose logs -f iiot-gateway
```

### Scale Services
```bash
docker-compose up -d --scale iiot-data=3
```

## Service URLs

- **API Gateway:** http://localhost:8000
- **Metadata Service:** http://localhost:8001
- **Data Service:** http://localhost:8002
- **Command Service:** http://localhost:8003
- **Consul UI:** http://localhost:8500
- **MQTT Broker:** tcp://localhost:1883
- **MQTT WebSocket:** ws://localhost:9001

## Benefits of Containerization

1. **Dependency Isolation** - Each service runs in isolated container
2. **Version Control** - Go 1.19 locked in container regardless of host
3. **Scalability** - Individual services can be scaled independently
4. **Development Consistency** - Same environment across all machines
5. **Infrastructure as Code** - Complete stack defined in docker-compose.yml

## Health Monitoring

All services include health checks and automatic restart policies. Monitor with:
```bash
docker-compose ps
docker-compose logs
```

## Production Considerations

- Enable authentication in production
- Configure TLS certificates
- Set up proper logging aggregation
- Implement service mesh for advanced networking
- Add monitoring with Prometheus/Grafana