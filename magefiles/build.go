//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

type Build mg.Namespace

// All Runs go mod download and then installs the binary.
func (Build) All() error {
	fmt.Println("Starting full build...")

	// Make everything nice, neat, and proper
	mg.Deps(Preps.Tidy)
	mg.Deps(Preps.Format)

	// Record all licenses in a registry
	mg.Deps(Checks.Licenses)

	// Check for code quality and security issues
	mg.Deps(Checks.Lint)
	mg.Deps(Checks.Security)

	return buildProject()
}

// All Runs go mod download and then installs the binary.
func (Build) CI() error {
	fmt.Println("Starting CI build...")

	// Check for code quality and security issues
	mg.Deps(Checks.Lint)
	mg.Deps(Checks.Security)

	return buildProject()

}
