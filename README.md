# resultx

A lightweight, zero-dependency Go module for creating standardized API responses with built-in support for pagination, metadata, and type-safe generics.

## Features

- **Generic Result Type** - Type-safe `Result[T]` wrapper for any data type
- **Standardized Responses** - Consistent JSON structure for success and error responses
- **Built-in Pagination** - Automatic page calculation, navigation flags, and limit validation
- **Functional Options** - Flexible metadata attachment using the options pattern
- **Zero Dependencies** - Pure Go implementation using only the standard library
- **JSON Ready** - Full serialization support with proper `omitempty` handling

## Installation

```bash
go get github.com/oshturhq/resultx
```

## Quick Start

```go
package main

import (
    "github.com/oshturhq/resultx"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // Success response
    user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
    result := resultx.Ok(user, "User retrieved successfully")

    // Error response
    errResult := resultx.Fail[User]("USER_NOT_FOUND", errors.New("user does not exist"))
}
```

## Usage

### Success Response

```go
data := MyData{Value: "example"}
result := resultx.Ok(data, "Operation completed successfully")
```

**JSON Output:**
```json
{
  "success": true,
  "data": {
    "value": "example"
  },
  "message": "Operation completed successfully"
}
```

### Error Response

```go
result := resultx.Fail[MyData]("VALIDATION_ERROR", errors.New("invalid input provided"))
```

**JSON Output:**
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "invalid input provided"
  }
}
```

### With Pagination

```go
// Create pagination request (with validation)
pageReq := resultx.NewPaginationRequest(0, 10) // offset: 0, limit: 10

// Fetch your data...
users := []User{...}
totalCount := int64(100)

// Create pagination response
pagination := resultx.NewPagination(totalCount, pageReq.GetOffset(), pageReq.GetLimit())

// Return result with pagination metadata
result := resultx.Ok(users, "Users retrieved", resultx.WithPagination(pagination))
```

**JSON Output:**
```json
{
  "success": true,
  "data": [...],
  "message": "Users retrieved",
  "meta": {
    "pagination": {
      "total": 100,
      "page": 1,
      "totalPages": 10,
      "limit": 10,
      "offset": 0,
      "hasNext": true,
      "hasPrev": false
    }
  }
}
```

### Search Requests

```go
searchReq := resultx.NewSearchRequest("  john doe  ")
query := searchReq.QueryValue() // Returns: "john doe" (trimmed)
```

## API Reference

### Result[T]

```go
type Result[T any] struct {
    Success bool      `json:"success"`
    Data    *T        `json:"data,omitempty"`
    Message string    `json:"message,omitempty"`
    Error   *Error    `json:"error,omitempty"`
    Meta    *Metadata `json:"meta,omitempty"`
}
```

### Functions

| Function | Description |
|----------|-------------|
| `Ok[T](val T, message string, opts ...MetaOption) *Result[T]` | Creates a successful result with data |
| `Fail[T](code string, err error, opts ...MetaOption) *Result[T]` | Creates a failed result with error details |
| `NewPaginationRequest(offset, limit int) PaginationRequest` | Creates a validated pagination request |
| `NewPagination(total int64, offset, limit int) Pagination` | Creates pagination metadata from total count |
| `NewSearchRequest(query string) SearchRequest` | Creates a search request wrapper |
| `WithPagination(pagination Pagination) MetaOption` | Option to attach pagination to result |

### Pagination Defaults

| Setting | Value |
|---------|-------|
| Default Limit | 10 |
| Maximum Limit | 100 |
| Minimum Offset | 0 |

## Complete Example

```go
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"

    "github.com/oshturhq/resultx"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
    // Parse pagination from query params
    offset := 0 // parse from r.URL.Query()
    limit := 10 // parse from r.URL.Query()

    pageReq := resultx.NewPaginationRequest(offset, limit)

    // Simulate database query
    products := []Product{
        {ID: 1, Name: "Laptop", Price: 999.99},
        {ID: 2, Name: "Mouse", Price: 29.99},
    }
    totalCount := int64(50)

    // Create response
    pagination := resultx.NewPagination(totalCount, pageReq.GetOffset(), pageReq.GetLimit())
    result := resultx.Ok(products, "Products retrieved", resultx.WithPagination(pagination))

    // Send JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    productID := 999 // parse from URL

    // Simulate not found error
    result := resultx.Fail[Product]("PRODUCT_NOT_FOUND", errors.New("product with ID 999 not found"))

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(result)
}
```
