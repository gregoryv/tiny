//go:build ci

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Similar to a make it should normalize How to accomplish
	// specific steps. Developers can use this locally but the
	// intention is to implement steps here and trigger them from
	// whatever pipeline stack you are using.
	for _, target := range os.Args[1:] {
		switch target {
		case "build":
			sh("go", "build", "./cmd/sayhi")

		case "test":
			sh("go", "test", "./...")

		case "dist":
			// one example where a step is more complicated than a
			// simple, cat | grep | sed.  The release notes in this
			// project are the same as the top section in our
			// changelog file.
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

// releaseNotes returns top changelog section
func releaseNotes() []byte {
	h2 := []byte("## [")
	from := bytes.Index(changelog, h2)
	to := bytes.Index(changelog[from+len(h2):], h2) + len(h2)
	notes := changelog[from:]
	if to > 0 {
		notes = changelog[from : from+to]
	}
	return bytes.TrimSpace(notes)
}

//go:embed changelog.md
var changelog []byte
