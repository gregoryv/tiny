[tiny](https://pkg.go.dev/github.com/gregoryv/tiny) - The Go + github getting started project

- Use basic repository layout with ./cmd
  See [Organizing Go Module](https://go.dev/doc/modules/layout), for
  alternatives.
- Use Go for [continuous integration](./ci.go) tasks
- Automated [build](.github/workflows/build.yml) with github actions
- Release with [release](.github/workflows/release.yml) flow and [goreleaser](.goreleaser.yaml)
- Including [MIT license](license.txt)

## Quickstart

    git clone git@github.com:gregoryv/tiny.git
	go run ci.go build test

