package assert

import (
	"io"
	"os"
	"sync"
)

type predicate func() bool

var (
	writer io.Writer = os.Stdout
	once   sync.Once
)

// DefaultWriter sets the output writer for the assert package.
// It can only be called once; subsequent calls will be ignored.
func DefaultWriter(w io.Writer) {
	once.Do(func() {
		writer = w
	})
}
