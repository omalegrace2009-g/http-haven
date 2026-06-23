# HTTP Haven

## Description

HTTP Haven is a Go web server project designed to demonstrate fundamental HTTP concepts using the `net/http` package.

The project consists of a collection of endpoints that explore common web development concepts such as:

* Routing
* Query parameters
* Request methods
* Request headers
* Status codes
* Authentication
* Redirects
* Request body processing

Each endpoint focuses on a specific aspect of HTTP communication and server-side request handling.

---

## Author

* Omale Grace

---

## Learning Objectives

This project demonstrates how to:

* Create HTTP servers using Go.
* Register and manage routes.
* Handle GET and POST requests.
* Parse query parameters.
* Read request bodies.
* Access HTTP headers.
* Return appropriate HTTP status codes.
* Implement simple API key authentication.
* Configure HTTP redirects.

---

## Features

### Exercise 1 — Ping-Pong Server

**Route:** `/ping`

Returns a simple response:

```text
pong
```

Used to verify that the server is running correctly.

---

### Exercise 2 — Query Parameters

**Route:** `/hello`

Accepts an optional query parameter:

```text
/hello?name=Alice
```

Response:

```text
Hello, Alice!
```

If no name is provided:

```text
Hello, Guest!
```

Only GET requests are allowed.

---

### Exercise 3 — Text Counter

**Route:** `/count`

#### GET

Displays usage instructions.

#### POST

Accepts text in the request body and returns its length.

Example:

```text
Request Body:
Golang

Response:
6
```

---

### Exercise 4 — Basic Math API

**Route:** `/calculate`

Performs arithmetic operations using query parameters.

Example:

```text
/calculate?op=add&a=12&b=8
```

Response:

```text
20
```

Supported validation includes:

* Invalid numbers
* Missing parameters
* Unsupported operations

Returns:

```text
400 Bad Request
```

for invalid input.

---

### Exercise 5 — User-Agent Echo

**Route:** `/agent`

Reads and returns the client's User-Agent header.

Example:

```http
User-Agent: CustomTester/1.0
```

Response:

```text
CustomTester/1.0
```

---

### Exercise 6 — Secure Dashboard

**Route:** `/dashboard`

A protected endpoint that requires an API key.

Required header:

```http
X-API-Key: secret123
```

Successful response:

```text
Welcome
```

Missing or incorrect keys return:

```text
401 Unauthorized
```

---

### Exercise 7 — Redirector

**Route:** `/legacy`

Redirects permanently to:

```text
/v2
```

using:

```text
301 Moved Permanently
```

The `/v2` endpoint serves the updated resource.

---

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd http-haven
```

---

## Running the Application

Start the server:

```bash
go run .
```

The server runs on:

```text
http://localhost:8080
```

---

## API Endpoints

| Method    | Route      | Description                    |
| --------- | ---------- | ------------------------------ |
| GET       | /ping      | Health check endpoint          |
| GET       | /hello     | Greets a user                  |
| GET, POST | /count     | Counts text length             |
| GET       | /calculate | Performs arithmetic operations |
| GET       | /agent     | Returns User-Agent header      |
| GET       | /dashboard | Protected dashboard endpoint   |
| GET       | /legacy    | Redirects to `/v2`             |
| GET       | /v2        | Updated endpoint               |

---

## Testing

The repository includes an automated verification script:

```bash
./test_endpoints.sh
```

The script validates:

* Endpoint responses
* HTTP methods
* Status codes
* Authentication
* Redirect behavior

Example:

```bash
chmod +x test_endpoints.sh
./test_endpoints.sh
```

---

## HTTP Status Codes Used

| Status Code | Meaning            |
| ----------- | ------------------ |
| 200         | OK                 |
| 301         | Moved Permanently  |
| 400         | Bad Request        |
| 401         | Unauthorized       |
| 405         | Method Not Allowed |

---

## Technologies Used

* Go
* net/http
* Bash
* cURL

---

## Example Verification Output

```text
✔ PASS: Got 'pong'
✔ PASS: Query param parsed successfully
✔ PASS: POST request calculated length correctly
✔ PASS: Access granted with correct token header
✔ PASS: Route /legacy issues a 301 Permanent Redirect
```

---

## Project Goal

The purpose of HTTP Haven is to provide practical experience with core HTTP server concepts in Go while building confidence in handling requests, responses, headers, authentication, and routing.

---

## License

This project was developed for educational purposes.
