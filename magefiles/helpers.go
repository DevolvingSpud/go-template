//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/magefile/mage/sh"
)

// ------------------------------------------------------------
// Helper Functions
// ------------------------------------------------------------

// FoundOrInstalled checks for existence then installs a file if it's not there
func FoundOrInstalled(executableName, installURL string) (isInstalled bool) {
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

// GetFlags gets all the compile flags to set the version and stuff
func GetFlags() string {
	timestamp := time.Now().Format(time.RFC3339)
	hash := Hash()
	tag := Tag()
	if tag == "" {
		tag = "dev"
	}
	return fmt.Sprintf(`-X "github.com/DevolvingSpud/template/pkg/version.Timestamp=%s" -X "github.com/DevolvingSpud/template/pkg/version.CommitHash=%s" -X "github.com/DevolvingSpud/template/pkg/version.Version=%s"`,
		timestamp, hash, tag)
}

// Tag returns the git tag for the current branch or "" if none.
func Tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	return s
}

// hash returns the git hash for the current repo or "" if none.
func Hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return hash
}

func buildProject() error {
	// Download the project's dependencies
	if err := sh.RunV("go", "mod", "download"); err != nil {
		return err
	}

	// Build and Install the project
	fmt.Printf("Running go install...\n")
	if err := sh.RunV("go", "install", "-ldflags="+GetFlags(), "./..."); err != nil {
		return err
	}
	fmt.Println("Install complete")
	return nil
}
