// This file will handle the formatting and output of the files to the console.
// This is where the -l flag’s logic will reside (which mimics the exact output of ls -l),
// handling file permissions, user, group, size, modification time, etc.
package internal

import "fmt"

func UnravelFiles(files []FileInfo) {
	//var relPath string
	for i := range files {
		fmt.Println(files[i].DocName)
		if len(files[i].RecursiveList) > 0 {
			UnravelFiles(files[i].RecursiveList)
		}
	}
}
