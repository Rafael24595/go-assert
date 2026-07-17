# go-assert

A lightweight, zero-dependency assertion framework for Go. It provides a flexible way to handle assertions in development, testing, and production using Go build tags.

---

## Features

- Three Modes: Choose between panic, io.Writer logging, or no-op (zero overhead) using build tags.
- Comprehensive Testing API: Assertions for `testing.T` including equality, panic, delta, collections, and concurrency helpers.
- Type Safe: Leverages Go Generics (`comparable` and numeric constraints) for compile-time safety.
- Idiomatic: Clean, minimal, and follows Go's "less is more" philosophy.

---

## Installation

```sh
go get github.com/Rafael24595/go-assert
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

Use assertions directly with `testing.T`.

```go
package main

import (
	"testing"

	assert "github.com/Rafael24595/go-assert/assert/test"
)

func TestCollections(t *testing.T) {
	assert.Equal(t, 4, 2+2)

	assert.GreaterThan(t, 3, "golang")

	assert.Size(t, 2, map[string]int{
		"go":   1,
		"rust": 2,
	})

	assert.Panic(t, func() {
		panic("boom")
	})
}
```

---

## Packages

The project provides two independent assertion packages:

| Package | Purpose |
|:--|:--|
| `github.com/Rafael24595/go-assert/assert/runtime` | Runtime assertions controlled through build tags. |
| `github.com/Rafael24595/go-assert/assert/test` | Assertions for `testing.T` in unit tests. |

## Runtime Assertions (`assert/runtime`)

Runtime assertions are controlled globally using Go build tags.

### Available Assertions

| Assertion | Description |
|:--|:--|
| `Unreachable` | Marks code paths that should never execute. |
| `True` | Ensures a condition is true. |
| `False` | Ensures a condition is false. |
| `LazyTrue` | Lazily evaluates a condition. |
| `LazyFalse` | Lazily evaluates a condition. |

### Build Configurations

You can control how the package behaves globally by using -tags during compilation or testing:

| Build Tag | Behavior | Logic Implementation | Best Use Case |
|:----------|:---------|:---------------------|:--------------|
| `g_ast_dbg` | **Panic** | `panic(message)`  | Local development & Debugging |
| `g_ast_wrt` | **Write** | `writer.Write([]byte(msg))` | CI/CD, Staging or Logging |
| *None (default)*| **No-op** | `func() {}` (Empty body) | Production (Zero overhead) |

### Usage Examples:

```sh
# Run tests with panic on assertion failure
go test -tags g_ast_dbg ./...

# Build your app with logging assertions
go build -tags g_ast_wrt -o myapp .
```

---

## Testing Assertions (`assert/test`)

The `assert/test` package provides helpers for `testing.T`.

### Equality Assertions

| Assertion | Description |
|:--|:--|
| `Nil` | Asserts that the value is nil. |
| `NotNil` | Asserts that the value is not nil. |
| `True` | Asserts that the value is true. |
| `False` | Asserts that the value is false. |
| `Equal` | Asserts equality between comparable values. |
| `NotEqual` | Asserts inequality between comparable values. |
| `DeepEqual` | Asserts deep equality using `reflect.DeepEqual`. |
| `NotDeepEqual` | Asserts deep inequality using `reflect.DeepEqual`. |
| `Same` | Asserts that two reference types point to the same object. |
| `NotSame` | Asserts that two reference types do not point to the same object. |
| `InDelta` | Asserts floating point proximity within a delta. |

### Ordering Assertions

These assertions support both:

- numeric values
- measurable values (`string`, `slice`, `array`, `map`, `chan`)

Measurable values are compared using `len(...)`.

| Assertion | Description |
|:--|:--|
| `Size` | Asserts `len(have) == want`. |
| `Empty` | Asserts `len(have) == 0`. |
| `NotEmpty` | Asserts `len(have) > 0`. |
| `LessThan` | Asserts `have < want`. |
| `LessOrEqual` | Asserts `have <= want`. |
| `GreaterThan` | Asserts `have > want`. |
| `GreaterOrEqual` | Asserts `have >= want`. |

### Collection Assertions

| Assertion | Description |
|:--|:--|
| `Inside` | Asserts that a container includes a value. |
| `NotInside` | Asserts that a container does not include a value. |

Supported containers:

- `string`
- `slice`
- `array`
- `map`

### Error Assertions

| Assertion | Description |
|:--|:--|
| `ErrorType` | Asserts that an error is of a specific type. |
| `ErrorNotType` | Asserts that an error is not of a specific type. |
| `ErrorIs` | Asserts that an error matches the expected value (sentinel), even when wrapped. |
| `ErrorIsNot` | Asserts that an error does not match the specified value (sentinel), even when wrapped. |

### Panic Assertions

| Assertion | Description |
|:--|:--|
| `Panic` | Asserts that a function panics. |
| `PanicWithMessage` | Asserts a panic message. |
| `NotPanic` | Asserts that a function does not panic. |

### Concurrency Assertions

| Assertion | Description |
|:--|:--|
| `WillClose` | Asserts that a channel closes or receives before timeout. |

---

## Why go-assert?

`go-assert` focuses on:

- zero-cost runtime assertions in production
- simple and idiomatic testing helpers
- minimal API surface
- no magic or reflection-heavy DSLs
