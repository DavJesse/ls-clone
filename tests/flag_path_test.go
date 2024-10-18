package internal_test

import (
	internal "my-ls/internal/ls"
	"testing"
)

// Test non-empty paths
func TestIsValidPath_NonEmptyString(t *testing.T) {
	result := internal.IsValidPath("some/path")
	if !result {
		t.Errorf("Expected IsValidPath to return true for a non-empty string, but got false")
	}
}
