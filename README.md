# Nate's Template

A template Go project with Mage and all its add-ons

## Building

```shell

# You might have to install Mage
> go install github.com/magefile/mage

# Build the program, running fmt, tidy, linters, and gosec
> mage

# Run it
> template

# Test it
> go test ./...

# Benchmark it
> go test -bench=. ./...

```
