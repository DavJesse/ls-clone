package tests

import (
	"os"
	"runtime"
	"strings"
	"testing"

	internal "my-ls/internal/ls"
)

// Test non-empty paths
func TestIsValidPath_NonEmptyString(t *testing.T) {
	var result bool
	var err error
	system := runtime.GOOS

	// Test path based on operating system
	if system == "windows" {
		result, err = internal.IsValidPath("some\\path")
	} else {
		result, err = internal.IsValidPath("some/path")
	}
	if !result {
		t.Errorf("Expected true; Got false")
		t.Errorf("Reason: %v", err)
	}
}

// Test empty inputs
func TestIsValidPath_EmptyStringInput(t *testing.T) {
	result, _ := internal.IsValidPath("")
	if result {
		t.Errorf("Expected False; Got True")
	}
}

// Test single-character paths
func TestIsValidPath_SingleCharacterPath(t *testing.T) {
	result, err := internal.IsValidPath("a")
	if !result {
		t.Errorf("Expected true; Got false")
		t.Errorf("Reason: %v", err)
	}
}

// Test paths with only white spaces
func TestIsValidPath_WhitespaceOnly(t *testing.T) {
	result, _ := internal.IsValidPath("   ")
	if result {
		t.Errorf("Expected false; Got true")
	}
}

// Test unicode-character-paths
func TestIsValidPath_UnicodeCharacterPath(t *testing.T) {
	result, err := internal.IsValidPath("こんにちは")
	if !result {
		t.Errorf("Expected true; Got false")
		t.Errorf("Reason: %v", err)
	}
}

// Test very long paths
func TestIsValidPath_VeryLongPath(t *testing.T) {
	longPath := "a"
	for i := 0; i < 100000; i++ {
		longPath += "a"
	}
	result, err := internal.IsValidPath(longPath)
	if !result {
		t.Errorf("Expected true; Got false")
		t.Errorf("Reason: %v", err)
	}
}

// Test special-character-paths
func TestIsValidPath_SpecialCharacterPath(t *testing.T) {
	var result bool
	var err error
	system := runtime.GOOS

	// Test special characters based on opperating system
	if system == "windows" {
		result, err = internal.IsValidPath("#$%^&()")
	} else {
		result, err = internal.IsValidPath("#$%^&*()")
	}
	if !result {
		t.Errorf("Expected true; Got false")
		t.Errorf("Reason: %v", err)
	}
}

// Test paths with only numeric characters
func TestIsValidPath_NumericOnlyPath(t *testing.T) {
	result, err := internal.IsValidPath("1234567890")
	if !result {
		t.Errorf("Expected true; Got false")
		t.Errorf("Reason: %v", err)
	}
}

// Test case-sensitive paths
func TestIsValidPath_CaseSensitivePath(t *testing.T) {
	result1, err1 := internal.IsValidPath("Dir")
	result2, err2 := internal.IsValidPath("dir")
	if !result1 || !result2 {
		t.Errorf("Expected true; Got false")
		t.Errorf("Reason 1: %v", err1)
		t.Errorf("Reason 2: %v", err2)
	}
}

// Test leasing and trailing whitespaces
func TestIsValidPath_LeadingAndTrailingWhitespacePath(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"  path  ", false},
		{"  ", false},
		{" ", false},
		{"\t\n", false},
		{"", false},
	}

	for _, tc := range testCases {
		result, _ := internal.IsValidPath(tc.input)
		if result != tc.expected {
			t.Errorf("IsValidPath(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

// Test handling of current directory
func TestRetrieveFileInfo_CurrentDir(t *testing.T) {
	var expect []internal.FileInfo
	var result []internal.FileInfo
	var point int
	system := runtime.GOOS

	if system == "windows" {
		result = internal.RetrieveFileInfo(".", false)
		expect = []internal.FileInfo{
			{DocName: "flag_test.go"},
			{DocName: "ls_test.go"},
			{DocName: "path_test.go"},
			{DocName: "sort_args_test.go"},
		}

	} else {
		result = internal.RetrieveFileInfo(".", false)
		expect = []internal.FileInfo{
			{DocName: "flag_test.go"},
			{DocName: "ls_test.go"},
			{DocName: "path_test.go"},
			{DocName: "sort_args_test.go"},
		}
	}

	for point < len(result) && point < len(expect) {
		if result[point].DocName == expect[point].DocName {
			point++
		} else {
			t.Errorf("Expected %v, Got %v", expect[point].DocName, result[point].DocName)
			t.FailNow()
		}
	}
}

// Test handling of non current directory
func TestRetrieveFileInfo_NonCurrentDir(t *testing.T) {
	var expect []internal.FileInfo
	var result []internal.FileInfo
	var point int
	system := runtime.GOOS

	if system == "windows" {
		result = internal.RetrieveFileInfo("..\\", false)
		expect = []internal.FileInfo{
			{DocName: "\033[01;34mcmd\033[0m\\"},
			{DocName: "commit.sh"},
			{DocName: "go.mod"},
			{DocName: "\033[01;34minternal\033[0m\\"},
			{DocName: "LICENSE"},
			{DocName: "push_both.sh"},
			{DocName: "README.md"},
			{DocName: "run_my_ls.sh"},
			{DocName: "\033[01;34mtests\033[0m\\"},
		}

	} else {
		result = internal.RetrieveFileInfo("../", false)
		expect = []internal.FileInfo{
			{DocName: "\033[01;34mcmd\033[0m/"},
			{DocName: "\033[01;32mcommit.sh\033[0m*"},
			{DocName: "go.mod"},
			{DocName: "\033[01;34minternal\033[0m/"},
			{DocName: "LICENSE"},
			{DocName: "\033[01;32mpush_both.sh\033[0m*"},
			{DocName: "README.md"},
			{DocName: "\033[01;32mrun_my_ls.sh\033[0m*"},
			{DocName: "\033[01;34mtests\033[0m/"},
		}
	}

	for point < len(result) && point < len(expect) {
		if result[point].DocName == expect[point].DocName {
			point++
		} else {
			t.Errorf("Expected %v, Got %v", expect[point].DocName, result[point].DocName)
			t.FailNow()
		}
	}
}

// test only empty string inputs and one populated slot on indices 0 or 1
func TestCleanArgs(t *testing.T) {
	testCases := []struct {
		input  []string
		expect []string
	}{
		{[]string{"", "", ""}, nil},
		{[]string{"-l", ""}, []string{"-l"}},
		{[]string{"", "."}, []string{"."}},
	}

	for _, tc := range testCases {
		var pointer int
		output := internal.CleanArgs(tc.input)
		for pointer < len(output) && pointer < len(tc.expect) {
			if output[pointer] != tc.expect[pointer] {
				t.Errorf("Expected %v, Got %v", tc.expect[pointer], output[pointer])
				t.FailNow()
			}
			pointer++
		}
	}
}

// Test file input
// func TestRetrieveHardLinkCount_FileInput(t *testing.T) {
// 	// Create a temporary file
// 	tmpfile, err := os.CreateTemp("", "testfile")
// 	if err != nil {
// 		t.Fatalf("Failed to create temporary file: %v", err)
// 	}
// 	defer os.Remove(tmpfile.Name())

// 	// Get the hard link count
// 	count, err := internal.RetrieveHardLinkCount(tmpfile.Name())
// 	if err != nil {
// 		t.Fatalf("RetrieveHardLinkCount failed: %v", err)
// 	}

// 	// Check if the count is 1 (as a regular file typically has 1 hard link)
// 	if count != 1 {
// 		t.Errorf("Expected hard link count of 1, got %d", count)
// 	}
// }

// // Test directory input
// func TestRetrieveHardLinkCount_DirInput(t *testing.T) {
// 	// Create a temporary directory
// 	tempDir, err := os.MkdirTemp("", "testdir")
// 	if err != nil {
// 		t.Fatalf("Failed to create temporary directory: %v", err)
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Get the hard link count
// 	count, err := internal.RetrieveHardLinkCount(tempDir)
// 	if err != nil {
// 		t.Fatalf("RetrieveHardLinkCount failed: %v", err)
// 	}

// 	// On most systems, a new directory should have 2 hard links
// 	// (one for "." and one for its entry in the parent directory)
// 	expectedCount := 2
// 	if count != expectedCount {
// 		t.Errorf("Expected hard link count of %d, but got %d", expectedCount, count)
// 	}
// }

// // Test Symbolic link handling
// func TestRetrieveHardLinkCount_SymbolicLink(t *testing.T) {
// 	// Create a temporary directory
// 	tempDir, err := os.MkdirTemp("", "test")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp directory: %v", err)
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Create a file
// 	filePath := filepath.Join(tempDir, "testfile.txt")
// 	if err := os.WriteFile(filePath, []byte("test content"), 0o644); err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	// Create a symbolic link to the file
// 	linkPath := filepath.Join(tempDir, "testlink")
// 	if err := os.Symlink(filePath, linkPath); err != nil {
// 		t.Fatalf("Failed to create symbolic link: %v", err)
// 	}

// 	// Test RetrieveHardLinkCount on the symbolic link
// 	count, err := internal.RetrieveHardLinkCount(linkPath)
// 	if err != nil {
// 		t.Fatalf("RetrieveHardLinkCount failed: %v", err)
// 	}

// 	// The symbolic link itself should have 1 hard link
// 	if count != 1 {
// 		t.Errorf("Expected 1 hard link for symbolic link, got %d", count)
// 	}
// }

// // Test Handling of multiple links
// func TestRetrieveHardLinkCount_MultipleLinks(t *testing.T) {
// 	// Create a temporary file
// 	tmpfile, err := os.CreateTemp("", "testfile")
// 	if err != nil {
// 		t.Fatalf("Failed to create temporary file: %v", err)
// 	}
// 	defer os.Remove(tmpfile.Name())

// 	// Create a hard link to the temporary file
// 	hardlinkPath := tmpfile.Name() + "_hardlink"
// 	err = os.Link(tmpfile.Name(), hardlinkPath)
// 	if err != nil {
// 		t.Fatalf("Failed to create hard link: %v", err)
// 	}
// 	defer os.Remove(hardlinkPath)

// 	// Test RetrieveHardLinkCount
// 	count, err := internal.RetrieveHardLinkCount(tmpfile.Name())
// 	if err != nil {
// 		t.Fatalf("RetrieveHardLinkCount failed: %v", err)
// 	}

// 	expectedCount := 2 // Original file + 1 hard link
// 	if count != expectedCount {
// 		t.Errorf("Expected hard link count to be %d, but got %d", expectedCount, count)
// 	}
// }

// Test correctness
func TestIsExecutable(t *testing.T) {
	var file *os.File
	system := runtime.GOOS

	// Define path based on os
	if system == "windows" {
		file, _ = os.Open("..\\")
	} else {
		file, _ = os.Open("../")
	}
	defer file.Close()

	// All '.sh' files in this directory are executable
	// Fail test when detection fails
	entries, _ := file.Readdir(-1)
	for _, entry := range entries {
		if !entry.IsDir() && strings.Contains(entry.Mode().Perm().String(), "x") {
			if !internal.IsExecutable(entry) {
				t.Errorf("Expected true, Got false")
				t.FailNow()
			}
		}
	}
}
