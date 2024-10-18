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

// Test single-character paths
func TestIsValidPath_SingleCharacterPath(t *testing.T) {
	result := internal.IsValidPath("a")
	if !result {
		t.Errorf("Expected true; Got false")
	}
}

// Test paths with only white spaces
func TestIsValidPath_WhitespaceOnly(t *testing.T) {
	result := internal.IsValidPath("   ")
	if !result {
		t.Errorf("Expected true; Got false")
	}
}
