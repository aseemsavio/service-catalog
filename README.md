# Design Considerations for service-catalog

This document outlines the key design considerations for implementing `service-catalog` in a cloud-native environment.

### Framework Selection

- Chi is chosen for its lightweight and modular design, making it suitable for building microservices.
- An alternative considered was Gin, but Chi's simplicity and ease of use were prioritized.

---

# How to

## Prerequisites

- Go 1.25 or higher
- Docker

## Setup

1. Clone the repository
2. Navigate to the project directory
3. Run `make tidy` to install dependencies
4. Build the project using `make build`
5. Bring up the Docker container with `make up`
6. Once the docker containers are up, run the application with `make run`
7. To run the integration tests, make sure you have the docker daemon running and execute `make test-integration`.
