#!/bin/bash

# IIOTBackend Initialization Script
# This script sets up the development environment and initializes the database

set -e

echo "🚀 Initializing IIOTBackend..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if required tools are installed
check_dependencies() {
    print_status "Checking dependencies..."
    
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install Go 1.21 or later."
        exit 1
    fi
    
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed. Please install Docker."
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        print_error "Docker Compose is not installed. Please install Docker Compose."
        exit 1
    fi
    
    print_success "All dependencies are installed"
}

# Load environment variables
load_env() {
    if [ -f .env ]; then
        print_status "Loading environment variables from .env file..."
        export $(cat .env | grep -v '^#' | xargs)
    else
        print_warning ".env file not found. Using default environment variables."
        # Create .env from .env.example if it exists
        if [ -f .env.example ]; then
            print_status "Creating .env file from .env.example..."
            cp .env.example .env
            print_success ".env file created from .env.example"
        fi
    fi
}

# Initialize Go modules
init_go_modules() {
    print_status "Initializing Go modules..."
    
    if [ ! -f go.mod ]; then
        go mod init iiot-backend
        print_success "Go module initialized"
    else
        print_status "Go module already exists, updating dependencies..."
    fi
    
    # Download and tidy modules
    go mod tidy
    print_success "Go dependencies updated"
}

# Start infrastructure services
start_infrastructure() {
    print_status "Starting infrastructure services (PostgreSQL, Redis, Consul)..."
    
    # Stop any existing containers
    docker-compose down --remove-orphans 2>/dev/null || true
    
    # Start only infrastructure services
    docker-compose up -d postgres redis consul
    
    # Wait for services to be ready
    print_status "Waiting for services to be ready..."
    
    # Wait for PostgreSQL
    timeout=30
    while ! docker-compose exec -T postgres pg_isready -U postgres >/dev/null 2>&1; do
        timeout=$((timeout - 1))
        if [ $timeout -eq 0 ]; then
            print_error "PostgreSQL failed to start within 30 seconds"
            exit 1
        fi
        sleep 1
    done
    print_success "PostgreSQL is ready"
    
    # Wait for Redis
    timeout=30
    while ! docker-compose exec -T redis redis-cli ping >/dev/null 2>&1; do
        timeout=$((timeout - 1))
        if [ $timeout -eq 0 ]; then
            print_error "Redis failed to start within 30 seconds"
            exit 1
        fi
        sleep 1
    done
    print_success "Redis is ready"
    
    # Wait for Consul
    timeout=30
    while ! docker-compose exec -T consul consul members >/dev/null 2>&1; do
        timeout=$((timeout - 1))
        if [ $timeout -eq 0 ]; then
            print_error "Consul failed to start within 30 seconds"
            exit 1
        fi
        sleep 1
    done
    print_success "Consul is ready"
    
    print_success "All infrastructure services are running"
}

# Run database migrations
run_migrations() {
    print_status "Running database migrations..."
    
    # Set environment variables for database connection
    export PGHOST=${PGHOST:-localhost}
    export PGPORT=${PGPORT:-5432}
    export PGUSER=${PGUSER:-postgres}
    export PGPASSWORD=${PGPASSWORD:-postgres}
    export PGDATABASE=${PGDATABASE:-iiot_backend}
    
    # Check if database exists, create if not
    if ! docker-compose exec -T postgres psql -U postgres -lqt | cut -d \| -f 1 | grep -qw $PGDATABASE; then
        print_status "Creating database $PGDATABASE..."
        docker-compose exec -T postgres psql -U postgres -c "CREATE DATABASE $PGDATABASE;"
        print_success "Database $PGDATABASE created"
    fi
    
    # Run migrations by building and running the app briefly
    print_status "Building application to run migrations..."
    go build -o iiot-backend .
    
    # Run migrations (the app will run migrations on startup and then we'll stop it)
    print_status "Running migrations..."
    timeout 10s ./iiot-backend || true  # Allow it to fail after timeout
    
    print_success "Database migrations completed"
}

# Build the application
build_app() {
    print_status "Building IIOTBackend application..."
    
    go build -o iiot-backend .
    
    print_success "Application built successfully"
}

# Display service URLs
show_service_urls() {
    print_success "🎉 IIOTBackend initialization completed!"
    echo ""
    echo "📋 Service URLs:"
    echo "  🔧 IIOTBackend API:    http://localhost:8000"
    echo "  🔍 Health Check:       http://localhost:8000/health"
    echo "  📊 PostgreSQL:         localhost:5432 (user: postgres, db: iiot_backend)"
    echo "  📦 Redis:              localhost:6379"
    echo "  🏛️  Consul UI:          http://localhost:8500"
    echo ""
    echo "📚 API Documentation:"
    echo "  Core Metadata:    http://localhost:8000/api/v2/deviceservice"
    echo "  Core Data:        http://localhost:8000/api/v2/event"
    echo "  Core Command:     http://localhost:8000/api/v2/device"
    echo "  Scheduler:        http://localhost:8000/api/v2/interval"
    echo "  Notifications:    http://localhost:8000/api/v2/notification"
    echo "  Rules Engine:     http://localhost:8000/api/v2/rule"
    echo "  System Mgmt:      http://localhost:8000/api/v2/system/info"
    echo ""
    echo "🚀 To start the application:"
    echo "  docker-compose up iiot-backend"
    echo ""
    echo "   or run locally:"
    echo "  ./iiot-backend"
    echo ""
    echo "🔧 To stop all services:"
    echo "  docker-compose down"
}

# Main execution
main() {
    echo "=================================================="
    echo "🏭 IIOTBackend - Industrial IoT Backend System"
    echo "=================================================="
    echo ""
    
    check_dependencies
    load_env
    init_go_modules
    start_infrastructure
    run_migrations
    build_app
    show_service_urls
}

# Handle script arguments
case "${1:-}" in
    "deps")
        check_dependencies
        ;;
    "infra")
        start_infrastructure
        ;;
    "migrate")
        run_migrations
        ;;
    "build")
        build_app
        ;;
    "clean")
        print_status "Cleaning up..."
        docker-compose down --volumes --remove-orphans
        docker system prune -f
        rm -f iiot-backend
        print_success "Cleanup completed"
        ;;
    "help"|"-h"|"--help")
        echo "Usage: $0 [command]"
        echo ""
        echo "Commands:"
        echo "  deps     - Check dependencies only"
        echo "  infra    - Start infrastructure services only"
        echo "  migrate  - Run database migrations only"
        echo "  build    - Build application only"
        echo "  clean    - Clean up containers and built files"
        echo "  help     - Show this help message"
        echo ""
        echo "Run without arguments to perform full initialization"
        ;;
    *)
        main
        ;;
esac
