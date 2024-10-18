package tests

import (
	internal "my-ls/internal/ls"
	"testing"
)

// Test empty string input
func TestIsValidFlag_EmptyString(t *testing.T) {
	result := internal.IsValidFlag("")
	if result != false {
		t.Errorf("Expected false; Got %v", result)
	}
}

// Test single-character input
func TestIsValidFlag_SingleCharacterInput(t *testing.T) {
	result := internal.IsValidFlag("l")
    if result != false {
        t.Errorf("Expected false; Got %v", result)
    }
}