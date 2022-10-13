//go:build mage
// +build mage

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

// Default sets the default target for Mage
var Default = Build.All
