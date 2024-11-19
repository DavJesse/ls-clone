// This file will handle the formatting and output of the files to the console.
// This is where the -l flag’s logic will reside (which mimics the exact output of ls -l),
// handling file permissions, user, group, size, modification time, etc.
package internal

import (
	"strings"
)

func UnravelFiles(dirName, indent string, files []FileInfo) string {

	var result strings.Builder
	result.WriteString(dirName + ":\n")

	for i, file := range files {
		result.WriteString(file.DocName) // Append the file name

		if i < len(files)-1 {
			result.WriteString(indent) // Add indentation for subsequent files
		}

		if i == len(files)-1 {
			result.WriteString("\n") // Add a newline for the last file
		}

	}

	for _, file := range files {
		if dirName == "." {
			dirName += "/"
		}
		// Recursively handle subdirectories if present
		if file.RecursiveList != nil {
			result.WriteString("\n") // Add a newline for subdirectories
			subResult := UnravelFiles(dirName+file.DocName, indent, file.RecursiveList)
			result.WriteString(subResult)
		}

	}

	return result.String()
}
