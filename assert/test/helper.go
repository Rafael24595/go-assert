package assert

import (
	"github.com/Rafael24595/go-assert/assert/internal"
)

func formatMessage(message ...any) string {
	if len(message) == 0 {
		return ""
	}

	return internal.FormatMessage(message...) + " - "
}
