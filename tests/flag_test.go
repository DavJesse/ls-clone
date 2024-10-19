package tests

import (
	internal "my-ls/internal/ls"
	"testing"
)

// Test empty string input
func TestIsValidFlag_EmptyString(t *testing.T) {
	result, _ := internal.IsValidFlag("")
	if result {
		t.Errorf("Expected false; Got %v", result)
	}
}

// Test single-character input
func TestIsValidFlag_SingleCharacterInput(t *testing.T) {
	result, _ := internal.IsValidFlag("l")
	if result {
		t.Errorf("Expected false; Got %v", result)
	}
}

// Test non-hyphenated input
func TestIsValidFlag_NoHyphenInput(t *testing.T) {
	result, _ := internal.IsValidFlag("a")
	if result {
		t.Errorf("Expected false; Got %v", result)
	}
}

// Test invalid flag character
func TestIsValidFlag_InvalidFlagCharacter(t *testing.T) {
	result, _ := internal.IsValidFlag("-m")
	if result {
		t.Errorf("Expected false; Got %v", result)
	}
}

// Test valid flag characters
func TestIsValidFlag_ValidFlagCharacter(t *testing.T) {
	result, err := internal.IsValidFlag("-Raltr")
	if !result {
		t.Errorf("Expected true; Got %v", result)
		t.Errorf("Reason: %v", err)
	}
}

// Test Reapeated Valid Flag Character
func TestIsValidFlag_ReapetedValidFlagCharacter(t *testing.T) {
	result, err := internal.IsValidFlag("-llllll")
	if !result {
		t.Errorf("Expected true; Got %v", result)
		t.Errorf("Reason: %v", err)
	}
}

// Test valid flag characters with spaces
func TestIsValid_ValidFlagWithSpaces(t *testing.T) {
	flag := "- l a"
	result, _ := internal.IsValidFlag(flag)
	if result {
		t.Errorf("Expected false; Got %v", result)
	}
}

// Test multiple-hyphens input
func TestIsValid_MultipleHyphens(t *testing.T) {
	result, _ := internal.IsValidFlag("--all")
	if result {
		t.Errorf("Expected false; Got %v", result)
	}
}
