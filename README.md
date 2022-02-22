tiny - The Go + github get started project

Goals is to have a complete setup with pipelines in place for starting
my next project.

## Quick start

    $ go install github.com/gregoryv/tiny/cmd/sayhi@latest
	
## Contribute

Cross platform steps are defined in steps.go which pipelines use

Build 

    $ go run -tags steps steps.go build
	
Test

    $ go run -tags steps steps.go test


