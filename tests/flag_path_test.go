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

// Test unicode-character-paths
func TestIsValidPath_UnicodeCharacterPath(t *testing.T) {
	result := internal.IsValidPath("こんにちは")
	if !result {
		t.Errorf("Expected true; Got false")
	}
}

// Test very long paths
func TestIsValidPath_VeryLongPath(t *testing.T) {
	longPath := "a"
	for i := 0; i < 1000000; i++ {
		longPath += "a"
	}
	result := internal.IsValidPath(longPath)
	if !result {
		t.Errorf("Expected true; Got false")
	}
}

// Test special-character-paths
func TestIsValidPath_SpecialCharacterPath(t *testing.T) {
	result := internal.IsValidPath("#$%^&*()")
	if !result {
		t.Errorf("Expected true; Got false")
	}
}
