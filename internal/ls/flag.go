// This file will be responsible for parsing and processing the flags provided in the command-line arguments (like -l, -R, -a, etc.).
// It will interpret how each flag modifies the behavior of the output.

package internal

import (
	"errors"
	"log"
	"runtime"
	"strings"
)

// Assigns arguments as valid flags and path
func SortArgs(args []string) (string, string, error) {

	var flag, path string
	var err error
	var valid bool

	// Set path to current directory if no arguments are given
	// For one argument, validate their status as a flag, then as a path
	// If two arguments, validate order and status as flags or paths
	if len(args) == 0 {
		path = "."

		return flag, path, err
	} else if len(args) == 1 {
		// Validate argument as flag
		// Valid flags start with a hyphen
		if strings.HasPrefix(args[0], "-") {
			valid, err = IsValidFlag(args[0])
			if valid {
				flag = args[0]
				path = "." // Set path to current directory

				return flag, path, err
			} else {
				return flag, path, err
			}

			// Validate argument as a path
		} else {
			valid, err = IsValidPath(args[0])
			if valid {
				path = args[0]

				return flag, path, err
			} else {
				return flag, path, err
			}
		}
	} else if len(args) == 2 {
		if strings.HasPrefix(args[0], "-") {
			// Validate first agument as a flag
			valid, err = IsValidFlag(args[0])
			if valid {
				flag = args[0]
			} else {
				return flag, path, err
			}

			// Validate second argument as a path
			valid, err = IsValidPath(args[1])
			if valid {
				path = args[1]
			} else {
				return "", path, err
			}
			return flag, path, err
		}

		// Check for invalid argument formats
		if strings.HasPrefix(args[1], "-") {
			err = errors.New("invalid format: check argument arrangement\nwe recommend: ./run_my_ls.sh [valid flag] [valid path]")
			return flag, path, err
		}
	}

	err = errors.New("too many arguments")
	return flag, path, err
}

func IsValidFlag(arg string) (bool, error) {
	var err error

	// A valid flag is at least two characters long
	if len(arg) < 2 {
		err = errors.New("flag is empty")
		return false, err
	}

	for i, char := range arg {
		// Check if flag starts with '-'
		if i == 0 && char != '-' {
			err = errors.New("illegal character: flag has a leading '-'")
			return false, err
		}

		// Check for non-valid flag characters after '-
		if i != 0 && !(char == 'R' || char == 'l' || char == 'a' || char == 't' || char == 'r') {
			err = errors.New("illegal character: flag has invalid character(s)")
			return false, err
		}
	}
	return true, err
}

func IsValidPath(arg string) (bool, error) {
	var err error
	// A valid path is a non-empty string
	if len(arg) == 0 {
		err = errors.New("path is empty")
		return false, err
	}

	// Check for problematic leading / trailing characters
	if strings.HasPrefix(arg, "-") {
		err = errors.New("illegal character: leading '-'")
		return false, err
	}

	// Identify illegal character based on operrating system
	system := runtime.GOOS
	log.Println("system: ", system)
	if system == "windows" {
		illegalChars := []string{"<", ">", ":", "\"", "/", "|", "?", "*"}

		// Check for illegal characters
		for i := range illegalChars {
			if strings.Contains(arg, illegalChars[i]) {
				err = errors.New("illegal character: path contains " + illegalChars[i])
				return false, err
			}
		}
	} else {
		illegalSep := "\\"
		if strings.Contains(arg, illegalSep) {
			err = errors.New("illegal character: path contains '\\'")
			return false, err
		}
	}

	invalidChars := []string{
		"\x00", // Null (NUL)
		"\x01", // Start of Heading (SOH)
		"\x02", // Start of Text (STX)
		"\x03", // End of Text (ETX)
		"\x04", // End of Transmission (EOT)
		"\x05", // Enquiry (ENQ)
		"\x06", // Acknowledge (ACK)
		"\x07", // Bell (BEL)
		"\x08", // Backspace (BS)
		"\x09", // Horizontal Tab (HT)
		"\x0A", // Line Feed (LF)
		"\x0B", // Vertical Tab (VT)
		"\x0C", // Form Feed (FF)
		"\x0D", // Carriage Return (CR)
		"\x0E", // Shift Out (SO)
		"\x0F", // Shift In (SI)
		"\x10", // Data Link Escape (DLE)
		"\x11", // Device Control 1 (DC1)
		"\x12", // Device Control 2 (DC2)
		"\x13", // Device Control 3 (DC3)
		"\x14", // Device Control 4 (DC4)
		"\x15", // Negative Acknowledge (NAK)
		"\x16", // Synchronous Idle (SYN)
		"\x17", // End of Transmission Block (ETB)
		"\x18", // Cancel (CAN)
		"\x19", // End of Medium (EM)
		"\x1A", // Substitute (SUB)
		"\x1B", // Escape (ESC)
		"\x1C", // File Separator (FS)
		"\x1D", // Group Separator (GS)
		"\x1E", // Record Separator (RS)
		"\x1F", // Unit Separator (US)
	}

	// Check for illegal characters
	for i := range invalidChars {
		if strings.Contains(arg, invalidChars[i]) {
			err = errors.New("illegal character: path contains " + invalidChars[i])
			return false, err
		}
	}

	invalidChars = []string{
		"\n",   // Newline
		"\t",   // Tab
		"\b",   // Backspace
		"\r",   // Carriage Return
		"\033", // ESC
		"\x1B", // ESC
		"\x7F", // DEL
		" ",    // Shitespace
	}
	// Check for more illegal characters
	for i := range invalidChars {
		if strings.Contains(arg, invalidChars[i]) {
			err = errors.New("illegal character: path contains " + invalidChars[i])
			return false, err
		}
	}
	return true, err
}

// Cleans arguments, removing empty strings
func CleanArgs(args []string) []string {
	var result []string

	// Only append non-empty strings
	for i := range args {
		if args[i] != "" {
			result = append(result, args[i])
		}
	}
	return result
}
