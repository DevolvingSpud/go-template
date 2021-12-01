//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"

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
	mg.Deps(Lint)
	mg.Deps(Security)

	// Download the project's dependencies
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	// Build and install the project
	return sh.Run("go", "install", "./...")
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
	if err := sh.Run("go", "fmt", "./..."); err != nil {
		return err
	}
	return nil
}

// Lint runs various static checkers to ensure you follow The Rules(tm)
func Lint() error {
	fmt.Println("Running linter (go vet)...")
	if err := sh.Run("go", "vet", "./..."); err != nil {
		return err
	}
	return nil
}

// Security runs various static checkers to ensure you minimize security holes
func Security() error {
	fmt.Println("Running gosec...")

	// If gosec is missing, install it
	_, missing := exec.LookPath("gosec")
	if missing != nil {
		fmt.Println("installing gosec...")
		err := sh.Run("go", "install", "github.com/securego/gosec/cmd/gosec@latest")
		if err != nil {
			fmt.Println("Could not install gosec, skipping...")
			return nil
		}
	}

	if err := sh.Run("gosec", "-no-fail", "./..."); err != nil {
		return err
	}
	return nil
}
