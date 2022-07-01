//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default sets the default target for Mage
var Default = Build

// ------------------------------------------------------------
// Targets
//
// * Build (default)
// * Tidy
// * Format
// * Lint
// * Security
// * Test
// ------------------------------------------------------------

// Build Runs go mod download and then installs the binary.
func Build() error {
	fmt.Println("Starting build...")

	// Make everything nice, neat, and proper
	mg.Deps(Tidy)
	mg.Deps(Format)
	mg.Deps(Lint)
	mg.Deps(Security)

	// Download the project's dependencies
	if err := sh.RunV("go", "mod", "download"); err != nil {
		return err
	}

	// Build and Install the project
	fmt.Printf("Running go install...\n")
	if err := sh.RunV("go", "install", "-ldflags="+flags(), "./..."); err != nil {
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

// Lint runs various static checkers to ensure you follow The Rules(tm)
func Lint() error {
	fmt.Println("Running linter (go vet)...")
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		return err
	}

	isInstalled := installIfMissing("staticcheck", "honnef.co/go/tools/cmd/staticcheck@latest")
	if !isInstalled {
		return nil
	}
	fmt.Println("Running linter (staticcheck)...")
	if err := sh.RunV("staticcheck", "-f", "stylish", "./..."); err != nil {
		return err
	}

	return nil
}

// Security runs various static checkers to ensure you minimize security holes
func Security() error {
	fmt.Println("Running gosec...")

	// If gosec is missing, install it
	isInstalled := installIfMissing("gosec", "github.com/securego/gosec/v2/cmd/gosec@latest")
	if !isInstalled {
		return nil
	}

	if err := sh.RunV("gosec", "-no-fail", "./..."); err != nil {
		return err
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

// ------------------------------------------------------------
// Helper Functions
// ------------------------------------------------------------

// installIfMissing checks for existence then installs a file if it's not there
func installIfMissing(executableName, installURL string) (isInstalled bool) {
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

// flags gets all the compile flags to set the version and stuff
func flags() string {
	timestamp := time.Now().Format(time.RFC3339)
	hash := hash()
	tag := tag()
	if tag == "" {
		tag = "dev"
	}
	return fmt.Sprintf(`-X "github.com/DevolvingSpud/template/pkg/template/version.Timestamp=%s" -X "github.com/DevolvingSpud/template/pkg/template/version.CommitHash=%s" -X "github.com/DevolvingSpud/template/pkg/template/version.Tag=%s"`,
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
