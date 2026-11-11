package check

import "github.com/Rafael24595/go-assert/internal"

type Report func(format string, args ...any)

func prefix(message ...any) string {
	if len(message) == 0 {
		return ""
	}
	return internal.FormatMessage(message...) + " - "
}
