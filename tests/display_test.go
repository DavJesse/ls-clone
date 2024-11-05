package tests

import (
	"testing"

	internal "my-ls/internal/ls"
)

// Test empty directory inputs
func TestUnravelFiles_EmptyDirectory(t *testing.T) {
	dirName := "emptyDir"
	indent := ""
	files := []internal.FileInfo{}

	result := internal.UnravelFiles(dirName, indent, files)
	expected := "emptyDir:\n"

	if result != expected {
		t.Errorf("UnravelFiles(%q, %q, %v) = %q; want %q", dirName, indent, files, result, expected)
	}
}
