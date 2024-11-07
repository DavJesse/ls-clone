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

// Test output for flat directories
func TestUnravelFiles_FlatDirectory(t *testing.T) {
	files := []internal.FileInfo{
		{DocName: "file1.txt"},
		{DocName: "file2.txt"},
		{DocName: "file3.txt"},
	}

	expected := "testDir:\nfile1.txt  file2.txt  file3.txt\n"
	result := internal.UnravelFiles("testDir", "  ", files)

	if result != expected {
		t.Errorf("UnravelFiles returned incorrect result.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

// Test nested directories
func TestUnravelFiles_NestedDirectories(t *testing.T) {
	// Create a mock directory structure
	files := []internal.FileInfo{
		{DocName: "file1.txt"},
		{DocName: "dir1/", RecursiveList: []internal.FileInfo{
			{DocName: "file2.txt"},
			{DocName: "subdir1/", RecursiveList: []internal.FileInfo{
				{DocName: "file3.txt"},
			}},
		}},
		{DocName: "file4.txt"},
	}

	expected := `.:
file1.txt  dir1  file4.txt

./dir1:
file2.txt  subdir1

./dir1/subdir1:
file3.txt
`

	result := internal.UnravelFiles(".", "  ", files)

	if result != expected {
		t.Errorf("UnravelFiles returned unexpected result.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}
