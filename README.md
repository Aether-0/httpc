# HTTPC - HTTP Method Tester Tool


**HTTPC** is a lightweight and efficient Go-based utility designed to test various HTTP methods (e.g., GET, POST, PUT, DELETE) against a list of URLs. It is perfect for security researchers, developers, and system administrators who want to quickly identify which HTTP methods are supported by a server.

---

## Features

- **Multiple HTTP Methods**: Test a wide range of HTTP methods (e.g., GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, TRACE, CONNECT).
- **Concurrency**: Test multiple URLs concurrently for faster results.
- **Customizable**: Set concurrency levels, timeouts, retries, and user-agent headers.
- **Status Code Filtering**: Filter results based on specific HTTP status codes.
- **Output to File**: Save results to a file for further analysis.
- **Verbose and Silent Modes**: Choose between detailed logs or silent mode for minimal output.
- **Color-Coded Output**: Easily identify results with color-coded status codes.

---

## Installation

1. **Install Go**: Ensure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/Aether-0/httpc.git
   cd httpc
   ```

3. **Build the Tool**:
   ```bash
   go build -o httpc
   ```

4. **Run the Tool**:
   ```bash
   ./httpc -url https://example.com
   ```

---

## Usage

### Command-Line Arguments

| Argument           | Description                                                                 |
|--------------------|-----------------------------------------------------------------------------|
| `-c`               | Set concurrency level (default: `10`).                                     |
| `-timeout`         | Set request timeout (default: `5s`).                                       |
| `-retries`         | Set number of retries for failed requests (default: `1`).                  |
| `-ua`              | Set custom User-Agent header (default: `Go-HTTP-Client/1.1`).              |
| `-silent`          | Enable silent mode (only output results).                                  |
| `-v`               | Enable verbose mode for detailed logs.                                     |
| `-o`               | Save results to a file (e.g., `results.txt`).                              |
| `-m`               | Comma-separated HTTP methods to test (default: all methods).               |
| `-url`             | Test a single URL.                                                         |
| `-status`          | Comma-separated status codes to filter (e.g., `200,404`).                  |

### Examples

1. **Test a Single URL**:
   ```bash
   ./httpc -url https://example.com
   ```

2. **Test Multiple URLs from stdin**:
   ```bash
   cat urls.txt | ./httpc
   ```

3. **Test Specific HTTP Methods**:
   ```bash
   ./httpc -url https://example.com -m "GET,POST,PUT"
   ```

4. **Filter by Status Codes**:
   ```bash
   ./httpc -url https://example.com -status "200,404"
   ```

5. **Save Results to a File**:
   ```bash
   ./httpc -url https://example.com -o results.txt
   ```

6. **Increase Concurrency**:
   ```bash
   ./httpc -url https://example.com -c 20
   ```

---

## Output Example


```
   __   __  __       _____
  / /  / /_/ /____  / ___/
 / _ \/ __/ __/ _ \/ /__  
/_//_/\__/\__/ .__/\___/  
            /_/           

HTTPC - HTTP Method Tester

Default Methods: GET, PUT, POST, DELETE, PATCH, HEAD, OPTIONS, TRACE, CONNECT

[INFO] Starting HTTPC Tool...
[INFO] Concurrency Level: 10, Timeout: 5s, Retries: 1
[GET] https://example.com - 200
[POST] https://example.com - 405
[PUT] https://example.com - 403
[DELETE] https://example.com - 404
[SUCCESS] All requests completed.
```

---

## Configuration

### Custom HTTP Methods
You can specify custom HTTP methods using the `-m` flag:
```bash
./httpc -url https://example.com -m "GET,POST,PUT"
```

### Status Code Filtering
Filter results to show only specific status codes:
```bash
./httpc -url https://example.com -status "200,404"
```

---

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, feel free to open an issue or submit a pull request.

---

## License

This project is licensed under the **MTFK CPC (Mother Fucker Copy Cat)** License.  
Basically, do whatever the fuck you want, but don't be a copycat.  

---

## Contact

For questions or feedback, reach out to me on Telegram: [@k4b00m3](https://t.me/k4b00m3).

---

Happy Testing! 🚀
