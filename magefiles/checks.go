//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Checks mg.Namespace

// ------------------------------------------------------------
// Targets for the Magefile that do the quality checks.
// ------------------------------------------------------------

// Lint runs various static checkers to ensure you follow The Rules(tm)
func (Checks) Lint() error {
	fmt.Println("Running linter (go vet)...")
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		return err
	}

	if FoundOrInstalled("staticcheck", "honnef.co/go/tools/cmd/staticcheck@latest") {
		fmt.Println("Running linter (staticcheck)...")
		if err := sh.RunV("staticcheck", "-f", "stylish", "./..."); err != nil {
			return err
		}
	}

	return nil
}

// Security runs various static checkers to ensure you minimize security holes
func (Checks) Security() error {
	fmt.Println("Running gosec...")

	// If gosec is missing, install it
	if FoundOrInstalled("gosec", "github.com/securego/gosec/v2/cmd/gosec@latest") {
		if err := sh.RunV("gosec", "-no-fail", "-exclude-generated", "-exclude-dir=magefiles", "./..."); err != nil {
			return err
		}
	}

	return nil
}

// Licenses pulls down any dependent project licenses, checking for "forbidden ones"
func (Checks) Licenses() error {
	fmt.Println("Running go-licenses...")

	// Make the directory for the license files
	err := os.MkdirAll("licenses", os.ModePerm)
	if err != nil {
		return err
	}

	// If go-licenses is missing, install it
	if FoundOrInstalled("go-licenses", "github.com/google/go-licenses@latest") {
		// The header sets the columns for the contents
		csvHeader := "Package,URL,License\n"
		csvContents := ""

		if csvContents, err = sh.Output("go-licenses", "csv", "--ignore=github.com/DevolvingSpud", "./..."); err != nil {
			return err
		}

		// Write out the CSV file with the header row
		if err = os.WriteFile("./licenses/licenses.csv", []byte(csvHeader+csvContents+"\n"), 0666); err != nil {
			return err
		}
	}

	return nil
}
