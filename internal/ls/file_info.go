//This file will contain the logic for retrieving file and directory metadata.
// It uses Go’s os and syscall packages (since the os/exec package is not allowed).
//Functions here will retrieve file sizes, permissions, timestamps, etc.

package internal

import (
	"log"
	"os"
	"runtime"
)

func RetrieveFileInfo(path string) []string {
	var fileList []string

	// Open directory/file for reading
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entries, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve directory/file name and append to fileList
	// For directories, we add '/' or '\' depending on opperating system
	for _, entry := range entries {

		if entry.IsDir() {
			system := runtime.GOOS
			if system == "windows" {
				fileList = append(fileList, entry.Name()+"\\")
			} else {
				fileList = append(fileList, entry.Name()+"/")
			}
		} else {
			fileList = append(fileList, entry.Name())
		}
	}

	return fileList
}
