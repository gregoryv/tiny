package tiny

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	var buf bytes.Buffer
	Greet(&buf)
	if got := buf.String(); got != "Hello, world!" {
		t.Error(got)
	}
}
