package tests

import (
	internal "my-ls/internal/ls"
	"testing"
)

// Test for zero-length-arguments
func TestSortArgs_NoArguments(t *testing.T) {
	flag, path, err := internal.SortArgs([]string{})
	if flag != "" || path != "." || err != nil {
		if flag != "" {
			t.Errorf("Expected: \"\", Got: %#v", flag)
		}

		if path != "." {
			t.Errorf("Expected: \"\", Got: %#v", path)
		}

		if err != nil {
			t.Errorf("Expected: \"nil\", Got: %#v", err)
		}
	}
}
