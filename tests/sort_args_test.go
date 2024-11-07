package tests

import (
	"strings"
	"testing"

	internal "my-ls/internal/ls"
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

// Test for valid one-argument inputs
func TestSortArgs_OneValidArgument(t *testing.T) {
	flag, path, err := internal.SortArgs([]string{"-l"})
	if strings.Contains(flag, "-l") || path != "" || err != nil {
		if flag != "-l" {
			t.Errorf("Expected: '-l', Got: '%v'", flag)
		}

		if path != "." {
			t.Errorf("Expected: '', Got: '%v'", path)
		}

		if err != nil {
			t.Errorf("Expected: 'nil', Got: '%v'", err)
		}
	}
}

// Test for invalid one-argument inputs
func TestSortArgs_OneInValidArgument(t *testing.T) {
	flag, path, err := internal.SortArgs([]string{"-m"})
	if flag != "" || path != "" || err == nil {
		if flag != "" {
			t.Errorf("Expected: '', Got: '%v'", flag)
		}

		if path != "" {
			t.Errorf("Expected: '', Got: '%v'", path)
		}

		if err == nil {
			t.Errorf("Expected: \"illegal character: leading '-'\", Got: %#v", err)
		}
	}
}

// Test valid two-argument input
func TestSortArgs_TwoArgument(t *testing.T) {
	type result struct {
		Input []string
		Flag  string
		Path  string
		Err   bool
	}

	testCases := []result{
		{[]string{"-lRa", "directory/file"}, "-lRa", "directory/file", true},
		{[]string{"directory/file", "-lRa"}, "", "", false},
		{[]string{"-lRa", "directory\\file"}, "", "", false},
	}

	for _, tc := range testCases {
		var errStatus bool
		flag, path, err := internal.SortArgs(tc.Input)
		if err == nil {
			errStatus = true
		} else {
			errStatus = false
		}

		// Compare flag, path, and error outputs
		if flag != tc.Flag || path != tc.Path || errStatus != tc.Err {
			if flag != tc.Flag {
				t.Errorf("Expected: %#v, Got: %#v", tc.Flag, flag)
			}

			if path != tc.Path {
				t.Errorf("Expected: %#v, Got: '%#v'", tc.Path, path)
			}

			if errStatus != tc.Err {
				t.Errorf("Expected: %v, Got: %v", tc.Err, errStatus)
			}
		}

	}
}

// Test more than two arguments
func TestSortArgs_ExcessiveArguments(t *testing.T) {
	flag, path, err := internal.SortArgs([]string{"-m", "directory/file", "-l"})
	if flag != "" || path != "" || err == nil {
		if flag != "" {
			t.Errorf("Expected: '', Got: '%v'", flag)
		}

		if path != "" {
			t.Errorf("Expected: '', Got: '%v'", path)
		}

		if err == nil {
			t.Errorf("Expected: \"error\", Got: %v", err)
		}
	}
}

// Test empty directory
func TestSortByEmptyDir_EmptyArray(t *testing.T) {
	input := []internal.FileInfo{}
	result := internal.SortByEmptyDir(input)

	if len(result) != 0 {
		t.Errorf("Expected empty array, got array of length %d", len(result))
	}
}

// Test single element in directory
func TestSortByEmptyDir_SingleElement(t *testing.T) {
	input := []internal.FileInfo{
		{DocName: "file1.txt", RecursiveList: nil},
	}

	result := internal.SortByEmptyDir(input)

	if len(result) != 1 {
		t.Errorf("Expected result length 1, got %d", len(result))
	}

	if result[0].DocName != "file1.txt" {
		t.Errorf("Expected DocName 'file1.txt', got '%s'", result[0].DocName)
	}

	if result[0].RecursiveList != nil {
		t.Errorf("Expected RecursiveList to be nil, got non-nil")
	}
}

// Test directories with nil Recursive lists
func TestSortByEmptyDir_AllNil(t *testing.T) {
	input := []internal.FileInfo{
		{DocName: "file1.txt", RecursiveList: nil},
		{DocName: "file2.txt", RecursiveList: nil},
		{DocName: "file3.txt", RecursiveList: nil},
		{DocName: "file4.txt", RecursiveList: nil},
	}

	result := internal.SortByEmptyDir(input)

	if len(result) != len(input) {
		t.Errorf("Expected result length %d, got %d", len(input), len(result))
	}

	for i, file := range result {
		if file.DocName != input[i].DocName {
			t.Errorf("Expected file at index %d to be %s, got %s", i, input[i].DocName, file.DocName)
		}
		if file.RecursiveList != nil {
			t.Errorf("Expected RecursiveList to be nil for file %s", file.DocName)
		}
	}
}

// Test directories with non-nil Recursive lists
func TestSortByEmptyDir_AllNonNilRecursiveList(t *testing.T) {
	input := []internal.FileInfo{
		{DocName: "dir1", RecursiveList: []internal.FileInfo{{DocName: "file1"}}},
		{DocName: "dir2", RecursiveList: []internal.FileInfo{{DocName: "file2"}}},
		{DocName: "dir3", RecursiveList: []internal.FileInfo{{DocName: "file3"}}},
	}

	result := internal.SortByEmptyDir(input)

	for i, file := range result {
		if file.RecursiveList == nil {
			t.Errorf("Expected non-nil RecursiveList at index %d, but got nil", i)
		}
		if file.DocName != input[i].DocName {
			t.Errorf("Expected file name %s at index %d, but got %s", input[i].DocName, i, file.DocName)
		}
	}

	if len(result) != len(input) {
		t.Errorf("Expected result length %d, but got %d", len(input), len(result))
	}
}

// Test order of output when directories have nil Recursive lists
func TestSortByEmptyDir_AllNilRecursiveList(t *testing.T) {
	input := []internal.FileInfo{
		{DocName: "file1.txt", RecursiveList: nil},
		{DocName: "file2.txt", RecursiveList: nil},
		{DocName: "file3.txt", RecursiveList: nil},
	}
	expected := []internal.FileInfo{
		{DocName: "file1.txt", RecursiveList: nil},
		{DocName: "file2.txt", RecursiveList: nil},
		{DocName: "file3.txt", RecursiveList: nil},
	}

	result := internal.SortByEmptyDir(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, but got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i].DocName != expected[i].DocName {
			t.Errorf("Expected DocName %s at index %d, but got %s", expected[i].DocName, i, result[i].DocName)
		}
		if !(result[i].RecursiveList == nil && expected[i].RecursiveList == nil) {
			t.Errorf("Expected RecursiveList %v at index %d, but got %v", expected[i].RecursiveList, i, result[i].RecursiveList)
		}
	}
}
