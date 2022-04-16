tiny - The Go + github getting started project

## Automated build and test

Using github actions in this project, but there are other options.

## Automated relase on push tags

With goreleaser you get a simple start of releasing precombiled binaries

## Crossplatform steps for compiling and testing

When you start collaborate with people on different platforms and your
build process involves more that the basic `go build`, if you've
written helper shell scripts you may need to update or replace those.
This project uses a special `steps.go` file, so no other dependency
than go is needed, except for any subcommands you use inside it.

## Enabled godoc documentation

The godoc service does not provide documentation on executable,
ie. main, packages. Thus you should always place commands in
cmd/FOLDERS and keep the root of your package for the domain logic.


