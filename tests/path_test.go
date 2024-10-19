package tests

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
		t.Errorf("Expected False; Got True")
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
	if result {
		t.Errorf("Expected false; Got true")
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

// Test paths with only numeric characters
func TestIsValidPath_NumericOnlyPath(t *testing.T) {
	result := internal.IsValidPath("1234567890")
	if !result {
		t.Errorf("Expected true; Got false")
	}
}

// Test case-sensitive paths
func TestIsValidPath_CaseSensitivePath(t *testing.T) {
	result1 := internal.IsValidPath("Dir")
	result2 := internal.IsValidPath("dir")
	if !result1 || !result2 {
		t.Errorf("Expected true; Got false")
	}
}

// Test leasing and trailing whitespaces
func TestIsValidPath_LeadingAndTrailingWhitespacePath(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"  path  ", true},
		{"  ", true},
		{" ", true},
		{"\t\n", true},
		{"", false},
	}

	for _, tc := range testCases {
		result := internal.IsValidPath(tc.input)
		if result != tc.expected {
			t.Errorf("IsValidPath(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
