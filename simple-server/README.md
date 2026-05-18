# Go HTTP Server Documentation

## Overview

This project is a simple HTTP server written in Go using the built-in `net/http` package. The server:

* Serves static files from a `static` directory
* Handles a basic `/hello` route
* Processes form submissions through the `/form` route

The application runs on port `8080`.

---

## Project Structure

```text
project/
├── main.go
└── static/
    ├── index.html
    ├── form.html
```

---

## Features

## Static File Serving

The server serves files from the `./static` directory.

Example:

```go
fileServer := http.FileServer(http.Dir("./static"))
http.Handle("/", fileServer)
```

This allows users to access:

```text
http://localhost:8080/
```

---

## Hello Route

Endpoint:

```text
GET /hello
```

Returns:

```text
hello world
```

### Validation

The handler checks:

* The request path must be `/hello`
* The HTTP method must be `GET`

If the path or method is invalid, an error response is returned.

Example:

```go
if r.URL.Path != "/hello" {
    http.Error(w, "404 Not Found", http.StatusNotFound)
    return
}
```

---

## Form Submission Route

Endpoint:

```text
POST /form
```

Processes form data submitted by the client.

### Expected Form Fields

| Field     | Description  |
| --------- | ------------ |
| `name`    | User name    |
| `address` | User address |

### Example Response

```text
Post request successful
Name = John
Address = Johannesburg
```

---

## Code Explanation

### Main Function

The `main()` function:

1. Creates a static file server
2. Registers HTTP routes
3. Starts the HTTP server on port `8080`

```go
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

---

### helloHandler

Handles requests to:

```text
/hello
```

#### Responsibilities of helloHandler

* Validates URL path
* Validates HTTP method
* Sends response to client

#### Supported Method

```text
GET
```

#### Example

```go
fmt.Fprintf(w, "hello world")
```

---

### formHandler

Handles form submissions sent to:

```text
/form
```

#### Responsibilities of formHandler

* Parses form data
* Reads submitted values
* Returns confirmation response

### Form Parsing

```go
r.ParseForm()
```

### Reading Form Values

```go
name := r.FormValue("name")
address := r.FormValue("address")
```

---
