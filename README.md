# Analysis Utility

## Project Overview
Analysis Utility is a lightweight Go-based HTTP server that provides runtime metrics and structured logging to aid in monitoring and analysis. The server exposes Prometheus metrics on the `/metrics` endpoint, supports configurable runtime settings, and handles graceful shutdowns for robustness.

## Features
- HTTP server built with Gin framework
- Prometheus metrics endpoint (`/metrics`) exposing process and Go runtime metrics
- Structured logging using zerolog with both console and file outputs
- Graceful shutdown on termination signals (SIGINT, SIGTERM)
- Configuration through YAML file
- Dockerized for easy deployment and development

## Prerequisites
- Docker and Docker Compose for containerized deployment (optional)

## Installation & Build

Clone the repository:

```bash
git clone <github.com/avraam311/analysis-utility>
cd analysis-utility
```

## Configuration

The server loads its configuration from `config/local.yaml`. You can customize settings such as the server listening port in this file. Example:

```yaml
server:
  port: ":8080"
```

## Docker Setup

Build and run the application using Docker Compose:

```bash
docker-compose up --build
```

This will build the container image using the multi-stage `cmd/Dockerfile` and start the app container exposing port 8080.

Logs will be persisted to the `./logs` directory on the host.

## Project Structure

- `cmd/` – Contains the main application entrypoint and Dockerfile
- `internal/api/http/server/` – HTTP server setup with routing and middleware
- `internal/infra/config/` – Configuration loading utilities
- `internal/infra/logger/` – Logging setup using zerolog
- `internal/infra/prometheus/` – Prometheus metrics registry and collectors
- `config/` – YAML configuration files

## Logging

Logging is provided by [zerolog](https://github.com/rs/zerolog), configured to output to both the console in a human-friendly format and to a file at `/app/logs/app.log` within the container. Logs include timestamps and structured fields for ease of parsing.

## Metrics

Prometheus metrics are exposed on the `/metrics` HTTP endpoint. It includes:
- Process metrics (CPU, memory usage, file descriptors, etc.)
- Go runtime metrics (goroutines, garbage collection stats, etc.)

This allows easy integration with Prometheus-based monitoring systems.

## Graceful Shutdown

The application listens for system termination signals (`SIGINT`, `SIGTERM`) and attempts a graceful shutdown with a 5-second timeout. This ensures ongoing requests are completed before the server stops.
