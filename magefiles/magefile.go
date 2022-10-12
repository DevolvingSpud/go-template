//go:build mage

package main

// ------------------------------------------------------------
// The Magefile is a Go program that builds the project,
// using shell commands (that target both Windows and Linux)
//
// You can find more about it at https://github.com/magefile/mage
//
// Targets
//
// * Build (default)
// * Tidy
// * Format
// * Licenses
// * Lint
// * Security
// * Test
// * Benchmark
// ------------------------------------------------------------

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default sets the default target for Mage
var Default = Build

// Build Runs go mod download and then installs the binary.
func Build() error {
	fmt.Println("Starting build...")

	// Make everything nice, neat, and proper
	mg.Deps(Tidy)
	mg.Deps(Format)

	// Record all licenses in a registry
	mg.Deps(Licenses)

	// Check for code quality and security issues
	mg.Deps(Lint)
	mg.Deps(Security)

	// Download the project's dependencies
	if err := sh.RunV("go", "mod", "download"); err != nil {
		return err
	}

	// Build and Install the project
	fmt.Printf("Running go install...\n")
	if err := sh.RunV("go", "install", "-ldflags="+getFlags(), "./..."); err != nil {
		return err
	}
	fmt.Println("Install complete")
	return nil
}

// Tidy runs go mod tidy to update the go.mod and go.sum files
func Tidy() error {
	fmt.Println("Running go mod tidy...")
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	return nil
}

// Format prettifies your code in a standard way to prevent arguments over curly braces
func Format() error {
	fmt.Println("Running go fmt...")
	if err := sh.RunV("go", "fmt", "./..."); err != nil {
		return err
	}
	return nil
}

// Licenses pulls down any dependent project licenses, checking for "forbidden ones"
func Licenses() error {
	fmt.Println("Running go-licenses...")

	// Make the directory for the license files
	err := os.MkdirAll("licenses", os.ModePerm)
	if err != nil {
		return err
	}

	// If go-licenses is missing, install it
	if foundOrInstalled("go-licenses", "github.com/google/go-licenses@latest") {
		// The header sets the columns for the contents
		csvHeader := "Package,URL,License\n"
		csvContents := ""

		if csvContents, err = sh.Output("go-licenses", "csv", "--ignore=github.com/DevolvingSpud", "./..."); err != nil {
			return err
		}

		// Write out the CSV file with the header row
		err = os.WriteFile("./licenses/licenses.csv", []byte(csvHeader+csvContents+"\n"), 0666)
	}

	return nil
}

// Test the project
func Test() error {
	fmt.Println("Running go test...")
	if err := sh.RunV("go", "test", "-v", "./..."); err != nil {
		return err
	}
	return nil
}

// Test the project
func Benchmark() error {
	fmt.Println("Running go test -bench...")
	if err := sh.RunV("go", "test", "-bench=.", "./..."); err != nil {
		return err
	}
	return nil
}

// ------------------------------------------------------------
// Targets for the Magefile that do the quality checks.
// ------------------------------------------------------------

// Lint runs various static checkers to ensure you follow The Rules(tm)
func Lint() error {
	fmt.Println("Running linter (go vet)...")
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		return err
	}

	if foundOrInstalled("staticcheck", "honnef.co/go/tools/cmd/staticcheck@latest") {
		fmt.Println("Running linter (staticcheck)...")
		if err := sh.RunV("staticcheck", "-f", "stylish", "./..."); err != nil {
			return err
		}
	}

	return nil
}

// Security runs various static checkers to ensure you minimize security holes
func Security() error {
	fmt.Println("Running gosec...")

	// If gosec is missing, install it
	if foundOrInstalled("gosec", "github.com/securego/gosec/v2/cmd/gosec@latest") {
		if err := sh.RunV("gosec", "-no-fail", "-exclude-generated", "-exclude-dir=magefiles", "./..."); err != nil {
			return err
		}
	}

	return nil
}

// ------------------------------------------------------------
// Helper Functions
// ------------------------------------------------------------

// foundOrInstalled checks for existence then installs a file if it's not there
func foundOrInstalled(executableName, installURL string) (isInstalled bool) {
	_, missing := exec.LookPath(executableName)
	if missing != nil {
		fmt.Printf("installing %v...\n", executableName)
		err := sh.Run("go", "install", installURL)
		if err != nil {
			fmt.Printf("Could not install %v, skipping...\n", executableName)
			return false
		}
		fmt.Printf("%v installed...\n", executableName)
	}
	return true
}

// getFlags gets all the compile flags to set the version and stuff
func getFlags() string {
	timestamp := time.Now().Format(time.RFC3339)
	hash := hash()
	tag := tag()
	if tag == "" {
		tag = "dev"
	}
	return fmt.Sprintf(`-X "github.com/DevolvingSpud/template/pkg/version.Timestamp=%s" -X "github.com/DevolvingSpud/template/pkg/version.CommitHash=%s" -X "github.com/DevolvingSpud/template/pkg/version.Version=%s"`,
		timestamp, hash, tag)
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	return s
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return hash
}
