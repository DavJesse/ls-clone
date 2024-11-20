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
file1.txt  dir1/  file4.txt

./dir1/:
file2.txt  subdir1/

./dir1/subdir1/:
file3.txt
`

	result := internal.UnravelFiles(".", "  ", files)

	if result != expected {
		t.Errorf("UnravelFiles returned unexpected result.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestAddColor(t *testing.T) {
	subject := []string{"non", "exec", "img", "sym", "set gid"}
	expect := []string{
		"non",
		"\033[01;32mexec\033[0m",
		"\033[01;35mimg\033[0m",
		"\033[01;36msym\033[0m",
		"\033[01;30;43mset gid\033[0m",
	}

	var point int
	for point < len(subject) && point < len(expect) {
		if internal.AddColor(subject[point], subject[point]) != expect[point] {
			t.Errorf("Expected: %s, Got: %s", expect[point], internal.AddColor(subject[point], subject[point]))
			t.Errorf("TestAddColor failed at index: %d", point)
			t.FailNow()
		}
		point++
	}
}
