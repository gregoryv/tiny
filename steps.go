//go:build steps
// +build steps

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	for _, target := range os.Args[1:] {
		switch target {
		case "build":
			sh("go", "build", "./cmd/sayhi")

		case "test":
			sh("go", "test", "./...")

		case "dist":
			os.MkdirAll("dist", 0755)
			notes := "dist/release_notes.txt"
			os.WriteFile(notes, releaseNotes(), 0644)

		case "clear":
			os.RemoveAll("dist")

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

// reelaseNotes returns top changelog section
func releaseNotes() []byte {
	h2 := []byte("## [")
	from := bytes.Index(changelog, h2)
	to := bytes.Index(changelog[from+len(h2):], h2)
	if to == -1 { // only one
		return changelog[from:]
	}
	return changelog[from : to+from]
}

// Use the changelog for your release notes

//go:embed changelog.md
var changelog []byte
