//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Tests mg.Namespace

// Test runs go test on the project
func (Tests) Test() error {
	fmt.Println("Running go test...")
	if err := sh.RunV("go", "test", "-v", "./..."); err != nil {
		return err
	}
	return nil
}

// Benchmark runs go test -bench on the project
func (Tests) Benchmark() error {
	fmt.Println("Running go test -bench...")
	if err := sh.RunV("go", "test", "-bench=.", "./..."); err != nil {
		return err
	}
	return nil
}
