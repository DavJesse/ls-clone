package internal_test

import (
	internal "my-ls/internal/ls"
	"testing"
)

// Test non-empty paths
func TestIsValidPath_NonEmptyString(t *testing.T) {
	result := internal.IsValidPath("some/path")
	if !result {
		t.Errorf("Expected true; Got false")
	}
}

// Test empty inputs
func TestIsValidPath_EmptyStringInput(t *testing.T) {
	result := internal.IsValidPath("")
	if result {
		t.Errorf("Expected false; Got true")
	}
}
