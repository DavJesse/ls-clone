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
		result.WriteString(indent)
		result.WriteString(file.DocName)

		if i < len(files)-1 {
			result.WriteString("\n")
		}
		if files[i].RecursiveList != nil || len(files[i].RecursiveList) > 0 {

			result.WriteString("\n")
			subResult := UnravelFiles(dirName+file.DocName, indent, file.RecursiveList)
			result.WriteString(subResult)

			if i < len(files)-1 {
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
