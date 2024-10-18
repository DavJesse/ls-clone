// This file will be responsible for parsing and processing the flags provided in the command-line arguments (like -l, -R, -a, etc.).
// It will interpret how each flag modifies the behavior of the output.

package internal

// func SortArgs(args []string) {

// }

func IsValidFlag(arg string) bool {
	// A valid flag is at least two characters long
	if len(arg) < 2 {
		return false
	}

	for i, char := range arg {
		// Check if flag starts with '-'
		if i == 0 && char != '-' {
			return false
		}

		// Check for non-valid flag characters after '-
		if i != 0 && !(char == 'R' || char == 'l' || char == 'a' || char == 't' || char == 'r') {
			return false
		}
	}
	return true
}

func IsValidPath(arg string) bool {
	// // A valid path is a non-empty string
	// return len(arg) > 0

	// for _,
}
