# Pinned Company

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/pinned-company)](https://goreportcard.com/report/github.com/patricksferraz/pinned-company)
[![GoDoc](https://godoc.org/github.com/patricksferraz/pinned-company?status.svg)](https://godoc.org/github.com/patricksferraz/pinned-company)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, scalable company management system built with Go, featuring a clean architecture design and robust infrastructure.

## 🚀 Features

- RESTful API built with Fiber
- Clean Architecture implementation
- PostgreSQL database with GORM ORM
- Docker and Kubernetes support
- Swagger API documentation
- Kafka integration for event-driven architecture
- Environment-based configuration
- Database migrations support
- Comprehensive testing setup

## 🛠️ Tech Stack

- **Backend**: Go 1.18+
- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: GORM
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Message Broker**: Kafka
- **API Documentation**: Swagger
- **Environment Management**: godotenv

## 📋 Prerequisites

- Go 1.18 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)
- Make (for using Makefile commands)

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/patricksferraz/pinned-company.git
cd pinned-company
```

2. Copy the environment file and configure it:
```bash
cp .env.example .env
```

3. Start the application using Docker Compose:
```bash
docker-compose up -d
```

4. Run database migrations:
```bash
make migrate
```

5. Start the application in development mode:
```bash
make run
```

## 🏗️ Project Structure

```
.
├── app/          # Application layer
├── cmd/          # Command line entry points
├── domain/       # Domain models and interfaces
├── infra/        # Infrastructure implementations
├── k8s/          # Kubernetes configurations
└── utils/        # Utility functions and helpers
```

## 🔧 Development

### Available Make Commands

- `make build` - Build Docker containers
- `make ps` - Show running containers status
- `make logs` - Show containers logs
- `make up` - Start containers in detached mode
- `make start` - Start stopped containers
- `make stop` - Stop running containers
- `make down` - Stop and remove containers, networks, and volumes
- `make attach` - Attach to a running container (requires SERVICE parameter)
- `make prune` - Remove unused Docker data
- `make test` - Run test suite using Docker
- `make gtest` - Run Go tests with coverage report
- `make gen` - Generate protobuf files

### Hot Reload

The project uses Air for hot reloading during development. Configuration can be found in `.air.toml`.

## 📚 API Documentation

Once the application is running, you can access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

## 🐳 Docker Support

The project includes Docker configuration for both development and production environments:

- Development: Uses Docker Compose with hot-reload support
- Production: Multi-stage Dockerfile for optimized builds

## 🧪 Testing

Run the test suite:
```bash
make test
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Ferraz** - *Initial work* - [patricksferraz](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://github.com/go-gorm/gorm)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
