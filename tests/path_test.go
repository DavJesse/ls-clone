package tests

import (
	"log"
	internal "my-ls/internal/ls"
	"runtime"
	"testing"
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
	for i := 0; i < 1000000; i++ {
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
	var pointer int

	result := internal.RetrieveFileInfo(".")
	expect := []string{"flag_test.go", "ls_test.go", "path_test.go", "sort_args_test.go"}

	for pointer < len(result) && pointer < len(expect) {
		if result[pointer] != expect[pointer] {
			log.Println(result[pointer])
			t.Errorf("Expected %v, Got %v", expect, result)
			t.FailNow()
		} else {
			pointer++
		}
	}

}

// Test handling of current directory
func TestRetrieveFileInfo_NonCurrentDir(t *testing.T) {
	var pointer int
	var expect []string
	system := runtime.GOOS

	result := internal.RetrieveFileInfo("..\\")
	if system == "windows" {
		expect = []string{".git\\", "LICENSE", "README.md", "cmd\\", "commit.sh", "go.mod", "internal\\", "LICENSE", "push_both.sh", "README.md", "run_my_ls.sh", "tests\\"}
	} else {
		expect = []string{".git/", "LICENSE", "README.md", "cmd/", "commit.sh", "go.mod", "internal/", "LICENSE", "push_both.sh", "README.md", "run_my_ls.sh", "tests/"}
	}

	for pointer < len(result) && pointer < len(expect) {
		if result[pointer] != expect[pointer] {
			log.Println(result[pointer])
			t.Errorf("Expected %v, Got %v", expect[pointer], result[pointer])
			t.FailNow()
		} else {
			pointer++
		}
	}

}
