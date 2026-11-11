# go-assert

A lightweight, zero-dependency assertion library for Go. It covers two needs:

- **Tests:** assertions for `testing.T`, in two flavours: `assert`, which reports a failure and lets the test continue, and `require`, which reports it and stops the test.
- **Running code:** invariants that panic, write to an `io.Writer` or compile to nothing, chosen with build tags.

---

## Features

- **Two testing flavours, one implementation:** `assert` and `require` expose the same assertions and share their logic and messages; only what happens after a failure differs.
- **Three runtime modes:** panic, `io.Writer` logging or no-op (zero overhead), selected with build tags.
- **Type safe:** uses Go generics (`comparable` and numeric constraints) for compile-time checks.
- **Idiomatic:** small API, plain functions, no DSL.

---

## Installation

```sh
go get github.com/Rafael24595/go-assert
```

Requires Go 1.25.5 or later.

---

## Packages

| Package | Purpose |
|:--|:--|
| `github.com/Rafael24595/go-assert/assert` | Test assertions that report a failure and let the test continue (`t.Errorf`). |
| `github.com/Rafael24595/go-assert/require` | The same assertions, but they stop the test at the first failure (`t.Fatalf`). |
| `github.com/Rafael24595/go-assert/invariant` | Runtime invariants controlled through build tags. |

---

## Quick Start

### 1. Tests

```go
package user

import (
	"testing"

	"github.com/Rafael24595/go-assert/assert"
	"github.com/Rafael24595/go-assert/require"
)

func TestLoadUser(t *testing.T) {
	u, err := Load("123")

	require.Nil(t, err) // without a user there is nothing left to check

	assert.Equal(t, "123", u.ID)
	assert.Inside(t, "@", u.Email, "user %s must have a valid email", u.ID)
	assert.Size(t, 2, u.Roles)
}
```

### 2. Invariants in running code

Configure the output once (optional) and use the invariants anywhere in your code.

```go
import (
	"os"

	"github.com/Rafael24595/go-assert/invariant"
)

func main() {
	// Optional: set a custom writer for the g_ast_wrt mode (only the first call counts)
	invariant.DefaultWriter(os.Stderr)

	val := 10
	invariant.True(val > 5, "value must be greater than 5")
	invariant.Unreachable("this code should never execute")
}
```

---

## Testing Assertions (`assert` and `require`)

Both packages provide exactly the same assertions with the same signatures and messages; a test in the repository fails if one package gains an assertion the other lacks.

### `assert` or `require`?

| | `assert` | `require` |
|:--|:--|:--|
| The values do not match | Reports with `t.Errorf`; the test **continues** | Reports with `t.Fatalf`; the test **stops** |
| The assertion is misused (e.g. `Size` on a `bool`, `Same` on a value type, `Inside` on an `int`) | Stops (`t.Fatalf`) | Stops (`t.Fatalf`) |
| `Unreachable` is reached | Stops (`t.Fatalf`) | Stops (`t.Fatalf`) |

Use `require` when the rest of the test depends on the check, typically for set-up steps and before indexing a slice; use `assert` for the checks themselves, so a single run reports every mismatch.

```go
list, err := repo.List()

require.Nil(t, err)
require.Size(t, 2, list)          // the next lines index the slice

assert.Equal(t, "first", list[0].Name)
assert.Equal(t, "second", list[1].Name)
```

### Messages

Every assertion accepts an optional message after its arguments. When the first value is a string it is used as a format string; otherwise the values are joined with spaces. The message is placed before the failure description:

```go
assert.Equal(t, "1.25", got, "version of %s", lang)
// version of golang - Expected '1.25', but got '1.24'
```

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
| `Capacity` | Asserts `cap(have) == want`. |

### Collection Assertions

| Assertion | Description |
|:--|:--|
| `Inside` | Asserts that a container includes a value. |
| `NotInside` | Asserts that a container does not include a value. |

Supported containers:

- `string`
- `slice`
- `array`
- `map` (its keys)

### Error Assertions

| Assertion | Description |
|:--|:--|
| `ErrorType` | Asserts that an error is of a specific type, even when wrapped (`errors.As`). |
| `ErrorNotType` | Asserts that an error is not of a specific type. |
| `ErrorIs` | Asserts that an error matches the expected value (sentinel), even when wrapped (`errors.Is`). |
| `ErrorIsNot` | Asserts that an error does not match the specified value (sentinel), even when wrapped. |

### Panic Assertions

| Assertion | Description |
|:--|:--|
| `Panic` | Asserts that a function panics. |
| `PanicWithMessage` | Asserts that a function panics with a given message. |
| `NotPanic` | Asserts that a function does not panic. |

### Concurrency Assertions

| Assertion | Description |
|:--|:--|
| `WillClose` | Asserts that a channel closes or receives before a timeout. |

### Flow Assertions

| Assertion | Description |
|:--|:--|
| `Unreachable` | Fails and stops the test if a code path that should never run is reached. |

### Custom Assertions

The assertions take the small `T` interface (`Helper`, `Errorf`, `Fatalf`) instead of `*testing.T`, so `*testing.T`, `*testing.B` and your own test doubles all work. `assert.T` and `require.T` are the same type, so a helper written for one works with both:

```go
import (
	"testing"

	"github.com/Rafael24595/go-assert/assert"
)

func AssertValidUser(t assert.T, u User) {
	t.Helper()

	assert.NotEmpty(t, u.ID, "user ID should not be empty")
	assert.Inside(t, "@", u.Email, "user must have a valid email")
}

func TestUserCreation(t *testing.T) {
	u := User{ID: "123", Email: "gopher@golang.org"}

	AssertValidUser(t, u)
}
```

---

## Runtime Invariants (`invariant`)

Invariants state what must always hold while the program runs. Their behaviour is chosen globally with build tags, so they can be strict in development and cost nothing in production.

### Available Invariants

| Invariant | Description |
|:--|:--|
| `Unreachable` | Marks code paths that should never execute. |
| `True` | Ensures a condition is true. |
| `False` | Ensures a condition is false. |
| `LazyTrue` | Ensures a lazily evaluated condition is true. |
| `LazyFalse` | Ensures a lazily evaluated condition is false. |

### Build Configurations

| Build Tag | Behavior | Implementation | Best Use Case |
|:----------|:---------|:---------------|:--------------|
| `g_ast_dbg` | **Panic** | `panic(message)` | Local development & debugging |
| `g_ast_wrt` | **Write** | `writer.Write([]byte(message))` | CI/CD, staging or logging |
| *None (default)* | **No-op** | Empty function body | Production (zero overhead) |

```sh
# Run tests with a panic on a broken invariant
go test -tags g_ast_dbg ./...

# Build your app with logged invariants
go build -tags g_ast_wrt -o myapp .
```

With no tag the invariants do nothing, but their arguments are still evaluated; use `LazyTrue` / `LazyFalse` when computing the condition is expensive.

---

## Migrating from v0.10

v0.11 reorganises the packages and changes when `assert` stops a test.

| Before | Now |
|:--|:--|
| `assert "github.com/Rafael24595/go-assert/assert/test"` | `"github.com/Rafael24595/go-assert/assert"` |
| `assert "github.com/Rafael24595/go-assert/assert/runtime"` | `"github.com/Rafael24595/go-assert/invariant"` (calls become `invariant.True(...)`) |
| — | `"github.com/Rafael24595/go-assert/require"` (new) |

- In `assert`, `Size`, `Empty`, `Capacity`, `Panic`, `PanicWithMessage`, `NotPanic` and `WillClose` **no longer stop the test** on a mismatch, like the rest of the package. Where a test relied on that (for example `assert.Size` before indexing a slice), use the `require` version.
- The deprecated `Contains`, `NotContains`, `Len`, `Less` and `Greater` have been removed; use `Inside`, `NotInside`, `Size`, `LessThan` and `GreaterThan`.
- Failure messages now show the values involved: `Nil` prints the value and its type, `ErrorIs` / `ErrorIsNot` print both errors.

---

## Why go-assert?

`go-assert` focuses on:

- zero-cost runtime invariants in production
- simple and idiomatic testing helpers, with a clear choice between continuing and stopping
- a minimal API surface
- no magic or reflection-heavy DSLs
