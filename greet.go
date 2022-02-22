package tiny

import (
	"fmt"
	"io"
)

// Greet writes a greeting to the given writer.
func Greet(w io.Writer) {
	fmt.Fprint(w, "Hello, world!")
}
