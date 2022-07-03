# Nate's Template

A template Go project with Mage, GitHub Actions, and all its add-ons.  This is a good
starting point for a Go project that does the following:

* Provides a pretty comprehensive build script using Mage
* Includes GitHub Actions to build, test, and tag the project
* Has a simple bit of linked-list code mainly to enable:
    * Example unit tests with full coverage
    * Example benchmarks showing difference in performance between slices and pointers
* Has some built-in code-quality checks:
  * `go vet` and `staticcheck` for linting and quality
  * `gosec` for security linting
  * `go-licenses` to build an inventory of dependency licenses

## Building

```shell

# You might have to install Mage
> go install github.com/magefile/mage

# Build and test the program, running fmt, tidy, linters, gosec, and licenses
> mage

# Run it
> template

# Test it
> go test ./...

# or...
> mage Test

# Benchmark it (shows the differences in linked-list performance)
> go test -bench=. ./...

# or ...
> mage Benchmark

```
