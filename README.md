# IIOT Backend

A cutting-edge Industrial IoT (IIoT) backend system that leverages advanced microservices architecture for scalable and intelligent device integration and management.

## Overview

This project is a comprehensive Industrial IoT backend system built with Go and Echo framework, featuring:

- **High-Performance Architecture**: Built on Echo v4.11.4 for maximum throughput
- **Modular Design**: Six core modules for contracts, bootstrap, configuration, messaging, registry, and secrets management
- **IIOT v1 API**: RESTful API endpoints for device management and data collection
- **Database Integration**: PostgreSQL support with automated migrations
- **Microservices Ready**: Containerized deployment support with Docker

## Key Components

### Core Modules
- **go-mod-core-contracts**: Core data structures and DTOs
- **go-mod-bootstrap**: Service initialization and startup management
- **go-mod-configuration**: Configuration management and providers
- **go-mod-messaging**: MQTT, NATS, and message broker integrations
- **go-mod-registry**: Service discovery and registration
- **go-mod-secrets**: Security and secret management

### Services
- **Metadata Service**: Device and device profile management
- **Data Service**: Event and measurement data handling
- **Command Service**: Device command execution
- **Scheduler Service**: Job scheduling and automation
- **Notification Service**: Alert and notification management
- **Rules Service**: Business logic and rule processing
- **Management Service**: System administration
- **Registry Service**: Service discovery

## API Endpoints

### IIOT v1 API
- `GET /api/v1/ping` - Health check
- `GET /api/v1/config` - Service configuration
- `GET /api/v1/devicehandler/all` - List all device handlers
- `GET /api/v1/device/all` - List all devices
- `GET /api/v1/dataevent/all` - List all data events

## Quick Start

### Prerequisites
- Go 1.19.3 or later
- PostgreSQL database
- Docker (optional)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd iiot-backend
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run database migrations:
```bash
go run main.go
```

5. Start the server:
```bash
go run main.go
```

The server will start on port 8000 by default.

### Docker Deployment

```bash
docker-compose up -d
```

## Configuration

The application uses environment variables for configuration:

- `DATABASE_URL`: PostgreSQL connection string
- `PORT`: Server port (default: 8000)
- `LOG_LEVEL`: Logging level (debug, info, warn, error)

## Architecture

The IIOT Backend follows a microservices architecture with:

- **Echo Framework**: High-performance HTTP router and middleware
- **PostgreSQL**: Primary database for persistent storage
- **Modular Packages**: Separated concerns with clear interfaces
- **RESTful APIs**: Standard HTTP methods and status codes
- **Middleware Stack**: CORS, logging, authentication, and recovery

## Development

### Project Structure
```
iiot-backend/
├── config/              # Configuration management
├── middleware/          # HTTP middleware
├── models/             # Data models
├── pkg/                # Shared modules
│   ├── go-mod-core-contracts/
│   ├── go-mod-bootstrap/
│   ├── go-mod-configuration/
│   ├── go-mod-messaging/
│   ├── go-mod-registry/
│   └── go-mod-secrets/
├── services/           # Business logic services
├── utils/              # Utility functions
└── main.go            # Application entry point
```

### Building
```bash
go build -o iiot-backend main.go
```

### Testing
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions, please open an issue in the GitHub repository.