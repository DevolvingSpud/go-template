//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Preps mg.Namespace

// Tidy runs go mod tidy to update the go.mod and go.sum files
func (Preps) Tidy() error {
	fmt.Println("Running go mod tidy...")
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	return nil
}

// Format prettifies your code in a standard way to prevent arguments over curly braces
func (Preps) Format() error {
	fmt.Println("Running go fmt...")
	if err := sh.RunV("go", "fmt", "./..."); err != nil {
		return err
	}
	return nil
}
