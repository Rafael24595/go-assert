package internal

import (
	"fmt"
	"strings"
)

// FormatMessage processes the input arguments to return a formatted string.
// If the first argument is a string, it is treated as a format specifier.
// Otherwise, it returns a string representation of all arguments.
func FormatMessage(message ...any) string {
	if len(message) == 0 {
		return ""
	}

	if format, ok := message[0].(string); ok {
		return fmt.Sprintf(format, message[1:]...)
	}

	parts := make([]string, len(message))
	for i, m := range message {
		parts[i] = fmt.Sprint(m)
	}
	
	return strings.Join(parts, " ")
}
