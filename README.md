# Design Considerations for service-catalog

A robust REST API for managing service metadata in a microservice architecture.

## Design Considerations

### Architecture

The Service Catalog follows a layered architecture pattern:

- **HTTP Layer** (`internal/http`): Handles request parsing, validation, and response formatting
- **Service Layer** (`internal/service`): Contains business logic and orchestration
- **Repository Layer** (`internal/repo`): Manages data persistence with PostgreSQL
- **Supporting Components**:
    - Configuration management (`internal/config`)
    - Structured logging (`internal/logger`) using Zap
    - Database migrations (`internal/migrations`)

### Database

- PostgreSQL was chosen for its robustness, JSON support, and transaction capabilities
- GORM is used as an ORM for simplified data access
- Database migrations ensure schema consistency across environments

### API Design

- RESTful principles with appropriate HTTP methods and status codes
- Pagination implemented for resource collections
- Query parameters for filtering, sorting, and pagination control
- Consistent error response format

## Assumptions

1. **Authentication/Authorization**: The service assumes authentication is handled by an upstream service (API gateway), although asked for in the requirements. This was ignored due to time constraints and the focus on core functionality.
2. **Deployment**: Designed to be containerized and deployed in a Kubernetes environment
3. **Database**: Assumes PostgreSQL is available and configured properly
4. **User Requirements**: Users need to browse, search, and filter services with pagination support

## Trade-offs

### SQL vs NoSQL
- **Decision**: Used PostgreSQL (SQL) over NoSQL options
- **Pros**: ACID compliance, structured schema, simplicity for relational data (due to the presence of versions).
- **Cons**: Less flexibility for rapidly changing data structures (doesn't matter in this case as the schema is mostly stable).

### ORM vs Raw SQL
- **Decision**: Used GORM instead of raw SQL
- **Pros**: Faster development, SQL injection protection, simpler code
- **Cons**: Potential performance overhead, less control over complex queries (Ignored this con as the requirement is pretty straight forward. For more complex use cases, gorm may be bypassed with raw sql).

### Pagination Implementation
- **Decision**: Offset-based pagination vs cursor-based
- **Pros**: Simple implementation, familiar to API consumers
- **Cons**: Less efficient for large datasets (performance degrades as offset increases)

### API Design
- **Decision**: Query parameters for filtering/sorting vs POST body
- **Pros**: RESTful, cacheable services, easier for clients to use
- **Cons**: None comes to mind for this use case.

## Running the Application

```bash
# Set environment variables or use .env file
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=password
export POSTGRES_DB=servicecatalog
export HTTP_PORT=8080
export LOG_LEVEL=info

# Run the application
make run
```

## API Usage

### Pagination

The API supports pagination with the following query parameters:

- `page`: Page number (starts at 1, default: 1)
- `page_size`: Number of items per page (default: 20)
- `query`: Text search across service attributes
- `sort`: Field to sort by
- `order`: Sort direction (asc/desc)

Example: `/services?page=2&page_size=10&query=auth&sort=name&order=asc`

---

# How to Set up and Run this Project

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
