# go-assert

A lightweight, zero-dependency assertion framework for Go. It provides a flexible way to handle assertions in development, testing, and production using Go build tags.

---

## Features

- Three Modes: Choose between panic, io.Writer logging, or no-op (zero overhead) using build tags.
- Fluent Testing API: Comprehensive suite of assertions for testing.T (Equal, DeepEqual, Panic, InDelta, etc.).
- Type Safe: Leverages Go Generics (comparable, cmp.Ordered) for compile-time safety.
- Idiomatic: Clean, minimal, and follows Go's "less is more" philosophy.

---

## Installation

```sh
go get github.com/Rafael24595/go-assert
```

---

## Build Configurations

You can control how the package behaves globally by using -tags during compilation or testing:

| Build Tag | Behavior | Logic Implementation | Best Use Case |
|:----------|:---------|:---------------------|:--------------|
| `g_ast_dbg` | **Panic** | `panic(message)`  | Local development & Debugging |
| `g_ast_wrt` | **Write** | `writer.Write([]byte(msg))` | CI/CD, Staging or Logging |
| *None (default)*| **No-op** | `func() {}` (Empty body) | Production (Zero overhead) |

---

### Usage Examples:

```sh
# Run tests with panic on assertion failure
go test -tags g_ast_dbg ./...

# Build your app with logging assertions
go build -tags g_ast_wrt -o myapp .
```

---

## Quick Start

### 1. Global Assertions (Runtime)

Configure the output once (optional) and use assertions anywhere in your code.

```go
import (
	"os"

	assert "github.com/Rafael24595/go-assert/assert/runtime"
)

func main() {
    // Optional: Set a custom writer (only works once)
    assert.DefaultWriter(os.Stderr)

    val := 10
    assert.True(val > 5, "Value must be greater than 5")
    assert.Unreachable("This code should never execute")
}
```

### 2. Testing API
Use the robust set of tools for your *_test.go files.

```go
package main

import (
	"testing"

	assert "github.com/Rafael24595/go-assert/assert/test"
)

func TestCalculation(t *testing.T) {
	have := 10 + 5
	assert.Equal(t, 15, have, "Math should work")

	assert.InDelta(t, 0.333, 1.0/3.0, 0.001)

	assert.Panic(t, func() {
		panic("boom")
	})
}
```

---

## API Overview

- Comparison: Equal, NotEqual, Greater, Less, DeepEqual, InDelta.
- Nullability: Nil, NotNil.
- Boolean: True, False, LazyTrue, LazyFalse.
- Collections: Len, Contains, NotContains.
- Errors & Flow: Error, NotError, Panic, NotPanic, Unreachable.
