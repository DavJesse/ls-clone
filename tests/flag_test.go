package tests

import (
	internal "my-ls/internal/ls"
	"testing"
)

func TestIsValidFlag_EmptyString(t *testing.T) {
	result := internal.IsValidFlag("")
	if result != false {
		t.Errorf("Expected false; Got %v", result)
	}
}

