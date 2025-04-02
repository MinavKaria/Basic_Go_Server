# Go Server

This project is a simple HTTP server built with Go. It demonstrates how to set up a basic web server, handle routes, and connect to a PostgreSQL database.

## Project Structure

```
go-server
├── main.go          # Entry point of the application
├── handlers         # Contains HTTP request handlers
│   └── handlers.go  # Functions to handle HTTP requests
├── models           # Defines data structures
│   └── models.go    # Data models and methods
├── config           # Configuration settings
│   └── config.go    # Load configuration from config.yaml
├── database         # Database connection and operations
│   └── database.go  # Functions to connect and interact with PostgreSQL
├── config.yaml      # Configuration file for database connection
├── Dockerfile       # Instructions for Docker image build
├── docker-compose.yaml # Docker Compose configuration
├── go.mod           # Module dependencies
└── go.sum           # Checksums for module dependencies
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- PostgreSQL database
- Docker and Docker Compose (optional)

### Installation

1. Clone the repository:

   ```
   git clone <your-repository-url>
   cd go-server
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Configure database connection:

   Update the `config.yaml` file with your database credentials:

   ```yaml
   database:
     host: localhost # Use 'postgres' for Docker setup
     port: 5432
     user: postgres
     password: yourpassword
     dbname: demo
   ```

### Running the Server

#### Locally:

```
go run main.go
```

The server will start listening on `localhost:8080`.

#### With Docker:

```
docker-compose up -d
```

### API Endpoints

- `GET /users` - Get all users
- `GET /users/get?id=N` - Get user by ID
- `POST /users` - Create a new user

Example of creating a user:

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Minav Karia","email":"minavpkaria@gmail.com"}'
```

### Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.

### License

This project is licensed under the MIT License. See the LICENSE file for more details.
