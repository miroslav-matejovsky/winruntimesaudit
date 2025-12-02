//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Tidy cleans up go.mod and go.sum files.
func Tidy() error {
	if err := sh.RunWith(nil, "go", "mod", "tidy"); err != nil {
		return err
	}
	fmt.Println("tidy done")
	return nil
}

// Vet runs go vet on all packages.
func Vet() error {
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		return err
	}
	fmt.Println("vet done")
	return nil
}

// Fmt formats Go source files with gofmt.
func Fmt() error {
	if err := sh.RunV("gofmt", "-s", "-w", "."); err != nil {
		return err
	}
	fmt.Println("fmt done")
	return nil
}

// Lint runs golangci-lint on all packages.
func Lint() error {
	fmt.Println("linting...")
	if err := sh.RunV("golangci-lint", "run", "./..."); err != nil {
		return err
	}
	fmt.Println("linting done")
	return nil
}

// Test runs all tests and outputs results in testdox format using gotestsum.
func Test() error {
	return sh.RunV("gotestsum", "--format", "testdox", "./...")
}

// All runs all quality checks and tests.
func All() error {
	mg.Deps(Tidy, Fmt, Vet, Lint)
	if err := Test(); err != nil {
		return err
	}
	return nil
}

// Publish publishes the module to the Go registry by tagging and pushing the version.
// Requires VERSION environment variable to be set (e.g., VERSION=v1.0.0).
func Publish(version string) error {
	mg.Deps(All)

	// Tag the current commit
	if err := sh.RunV("git", "tag", version); err != nil {
		return fmt.Errorf("failed to create git tag: %w", err)
	}

	// Push the tag to the remote repository
	if err := sh.RunV("git", "push", "origin", version); err != nil {
		return fmt.Errorf("failed to push git tag: %w", err)
	}

	fmt.Printf("Published version %s as a git tag\n", version)
	return nil
}
