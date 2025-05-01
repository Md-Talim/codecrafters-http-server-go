[![progress-banner](https://backend.codecrafters.io/progress/http-server/33046764-71e9-4a66-bfbc-de337903bad4)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

# HTTP Server From Scratch

[![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8.svg)](https://golang.org)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen.svg)]()

A high-performance, concurrent HTTP/1.1 server implementation in Go, built from the ground up following RFC 2616 specifications. This project demonstrates deep understanding of network programming, TCP/IP protocols, and Go's concurrency patterns.

## 🚀 Key Features

*   **HTTP/1.1 Protocol Implementation:** Adheres to core aspects of the HTTP/1.1 specification (RFC 2616).
*   **Concurrent Request Handling:** Leverages Go's concurrency model (goroutines) to efficiently handle multiple client connections simultaneously.
*   **Custom Router:** Implements a basic but effective path-based routing system to direct requests to appropriate handlers.
*   **RESTful Endpoints:** Provides several standard endpoints:
    *   `GET /`: Basic health check.
    *   `GET /echo/{message}`: Echoes back the provided message in the response body.
    *   `GET /user-agent`: Returns the client's `User-Agent` header value.
    *   `GET /files/{filename}`: Serves static files from a specified directory.
    *   `POST /files/{filename}`: Allows clients to upload files to the server's specified directory.
*   **Static File Serving & Upload:** Capable of serving existing files and accepting file uploads with appropriate `Content-Type` handling.
*   **Header Parsing:** Correctly parses HTTP request headers, including `User-Agent`, `Content-Length`, `Content-Type`, and `Connection`.
*   **HTTP Compression:** Supports `gzip` compression based on the `Accept-Encoding` request header.
*   **Robust Error Handling:** Returns appropriate HTTP status codes (e.g., 200 OK, 201 Created, 404 Not Found, 500 Internal Server Error) and informative error messages.

## 🛠️ Technology Stack

- **Language:** Go (Golang)
- **Testing:** Go testing package
- **Network:** Native Go TCP/IP stack
- **Protocols:** HTTP/1.1
- **Core Libraries:** Go Standard Library (`net`, `bufio`, `os`, `fmt`, `strings`, `sync`, `compress/gzip`)

## 🏗️ Architecture

The project follows clean architecture principles with clear separation of concerns:

```
├── app/
│   ├── main.go          # Entry point
│   └── main_test.go     # Integration tests
├── internal/
│   ├── handlers/        # Request handlers
│   ├── protocol/        # HTTP protocol implementation
```

## 📥 Installation

Follow these steps to get the server running on your local machine.

**Prerequisites:**

*   Go (Version 1.18+ Recommended)
*   Git (for cloning the repository)

**Steps:**

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/md-talim/codecrafters-http-server-go.git
    # Or: git clone [Your Fork URL]
    ```

2.  **Navigate to the project directory:**
    ```bash
    cd codecrafters-http-server-go
    ```

3.  **Build the application:**
    ```bash
    go build -o http-server-go ./app/main.go
    ```
    This will create an executable file named `http-server-go` (or `http-server-go.exe` on Windows) in the current directory.

4.  **Environment Variables:**
    *   No specific environment variables are required for basic operation. Configuration is primarily handled via command-line arguments (see Configuration).

5.  **Run the application:**
    ```bash
    ./http-server-go [--directory /path/to/serve/files]
    ```
    *   Replace `/path/to/serve/files` with the actual directory you want the server to use for file operations (GET/POST `/files/...`). If omitted, file operations might default to the current working directory or be disabled depending on implementation details.
    *   The server will start listening on `0.0.0.0:4221`.

## 🚦 Usage

Once the server is running, you can interact with it using an HTTP client like `curl` or a web browser.

*   **Health Check:**
    ```bash
    curl -v http://localhost:4221/
    ```
    *Expected Response:* `HTTP/1.1 200 OK`

*   **Echo Service:**
    ```bash
    curl -v http://localhost:4221/echo/hello-world
    ```
    *Expected Response:* `HTTP/1.1 200 OK` with `hello-world` in the body.

*   **User-Agent:**
    ```bash
    curl -v --header "User-Agent: foobar/1.2.3" http://localhost:4221/user-agent
    ```
    *Expected Response:* `HTTP/1.1 200 OK` with `foobar/1.2.3` in the body.

*   **Get File (Requires `--directory` flag during startup):**
    ```bash
    # Assuming a file named 'example.txt' exists in the specified directory
    curl -v http://localhost:4221/files/example.txt
    ```
    *Expected Response:* `HTTP/1.1 200 OK` with the content of `example.txt`.

*   **Upload File (Requires `--directory` flag during startup):**
    ```bash
    curl -v --data "This is the file content." http://localhost:4221/files/newfile.txt
    ```
    *Expected Response:* `HTTP/1.1 201 Created`

*   **Not Found:**
    ```bash
    curl -v http://localhost:4221/nonexistent-path
    ```
    *Expected Response:* `HTTP/1.1 404 Not Found`

## 🧪 Testing

Run the test suite:

```bash
go test ./app
```

## 📖 Learning Resources

This project was built as part of the "Build Your Own HTTP Server" challenge from [CodeCrafters](https://codecrafters.io/). Key learning resources:

- [RFC 2616 - HTTP/1.1](https://www.w3.org/Protocols/rfc2616/rfc2616.html)
- [Go Net Package Documentation](https://golang.org/pkg/net/)

## 🙏 Acknowledgements

- [CodeCrafters](https://codecrafters.io/) for the project challenge and structure
- The Go team for excellent networking primitives and documentation
