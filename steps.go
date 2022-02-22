//go:build steps
// +build steps

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	for _, target := range os.Args[1:] {
		switch target {
		case "build":
			sh("go", "build", filepath.Join(".", "cmd", "sayhi"))

		case "test":
			sh("go", "test", filepath.Join(".", "..."))

		default:
			fmt.Fprint(os.Stderr, "unknown target:", target)
			os.Exit(1)
		}
	}
}

func sh(app string, args ...string) {
	c := exec.Command(app, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		os.Exit(1)
	}
}
