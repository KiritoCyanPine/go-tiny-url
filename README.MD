# Go Tiny URL

A modern URL shortening service built with Go, providing a clean web interface and RESTful API for creating and managing shortened URLs.

## Features

- URL shortening with custom path generation
- Smart URL sanitization and redirect handling

## Technologies Used

- Go 1.21.9
- Gin Web Framework for routing and HTTP handling
- HTML Templates for web interface
- Environment-based Configuration
- Testify framework for testing

## Setup and Installation

### Prerequisites

- Go 1.21.9 or later
- Git

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/kiritocyanpine/go-tiny-url.git
   cd go-tiny-url
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Configure environment variables:
   Create a `.env` file with the following options:

   ```
   HOST_ADDRESS=http://localhost:8080  # Default if not set
   HOST_PORT=8080                      # Optional custom port
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start at `http://localhost:8080` by default.

## Development and Testing

The project uses the testify framework for testing. Mock implementations are available in the `mocks/` directory for testing database interactions and business logic.

To run tests:

```bash
go test ./...
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
