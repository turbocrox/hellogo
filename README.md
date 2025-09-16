# HelloGo

A simple Go web application using the Chi router framework. This project demonstrates basic HTTP server setup with CORS support, JSON responses, and error handling.

## Features

- Lightweight HTTP server with Chi router
- CORS middleware for cross-origin requests
- JSON response utilities
- Environment variable configuration
- Basic API endpoints for readiness and error testing

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/turbocrox/hellogo.git
   cd hellogo
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory:
   ```
   PORT=8080
   ```

4. Build and run:
   ```bash
   go build
   ./hellogo
   ```

## Usage

The server will start on the port specified in the `PORT` environment variable. You can test the endpoints using curl or a browser.

### API Endpoints

- **GET /v1/ready**: Readiness check endpoint (currently returns an error for demonstration)
- **GET /v1/err**: Error demonstration endpoint

#### Example Requests

```bash
# Test readiness
curl http://localhost:8080/v1/ready

# Test error handling
curl http://localhost:8080/v1/err
```

## Project Structure

- `main.go`: Application entry point and server setup
- `handler_read.go`: Handler for the /ready endpoint
- `handler_err.go`: Handler for the /err endpoint
- `json.go`: Utility functions for JSON responses and error handling
- `go.mod` & `go.sum`: Go module dependencies

## Dependencies

- [Chi](https://github.com/go-chi/chi): Lightweight router
- [Chi CORS](https://github.com/go-chi/cors): CORS middleware
- [godotenv](https://github.com/joho/godotenv): Environment variable loader


## License

This project is open source. Feel free to use and modify as needed.
