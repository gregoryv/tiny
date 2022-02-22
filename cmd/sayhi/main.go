// Command sayhi greets the world
package main

import (
	"os"

	"github.com/gregoryv/tiny"
)

func main() {
	tiny.Greet(os.Stdout)
}
