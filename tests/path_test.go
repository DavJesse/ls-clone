package tests

import (
	"log"
	"os"
	"strings"
	"testing"

	internal "my-ls/internal/ls"
)

// Test non-empty paths
func TestIsValidPath_NonEmptyString(t *testing.T) {
	result, err := internal.IsValidPath("some/path")
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
	result, err := internal.IsValidPath("#$%^&*()")
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

	result = internal.RetrieveFileInfo(".", false, true)
	log.Printf("%#v", result)
	expect = []internal.FileInfo{
		{DocName: "display_test.go"},
		{DocName: "flag_test.go"},
		{DocName: "ls_test.go"},
		{DocName: "path_test.go"},
		{DocName: "sort_args_test.go"},
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

	result = internal.RetrieveFileInfo("../", false, true)
	expect = []internal.FileInfo{
		{DocName: "\033[01;34mcmd\033[0m"},
		{DocName: "\033[01;32mcommit.sh\033[0m*"},
		{DocName: "go.mod"},
		{DocName: "\033[01;34minternal\033[0m"},
		{DocName: "LICENSE"},
		{DocName: "\033[01;32mpush_both.sh\033[0m*"},
		{DocName: "README.md"},
		{DocName: "\033[01;32mrun_my_ls.sh\033[0m*"},
		{DocName: "\033[01;34mtests\033[0m"},
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

// Test correctness
func TestIsExecutable(t *testing.T) {
	file, _ := os.Open("../")

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
