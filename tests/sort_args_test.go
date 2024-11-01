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
