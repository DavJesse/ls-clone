//This file will contain the logic for retrieving file and directory metadata.
// It uses Go’s os and syscall packages (since the os/exec package is not allowed).
//Functions here will retrieve file sizes, permissions, timestamps, etc.

package internal

import (
	"errors"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"syscall"
)

func RetrieveFileInfo(path string) []FileInfo {
	var fileList []FileInfo

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
	for i, entry := range entries {
		// ignore git directories
		if strings.Contains(entry.Name(), ".git") {
			continue
		}

		if entry.IsDir() {
			system := runtime.GOOS
			// Append result for windows systems
			// Wrap text in bright blue
			if system == "windows" {
				if i != len(entries)-1 {
					fileList[i].DefaultList += "\033[01;34m" + entry.Name() + "\033[0m" + "\\" + " "
					fileList[i].DetailedList += entry.Mode().Perm().String() + "\033[01;34m" + entry.Name() + "\033[0m" + "\\" + " "
				} else {
					fileList[i].DefaultList += "\033[01;34m" + entry.Name() + "\033[0m" + "\\"
				}

				// Append result for other systems
			} else {

				if i != len(entries)-1 {
					fileList[i].DefaultList += "\033[01;34m" + entry.Name() + "\033[0m" + "/" + " "
				} else {
					fileList[i].DefaultList += "\033[01;34m" + entry.Name() + "\033[0m" + "/"
				}
			}
		} else {
			fileList[i].DefaultList += entry.Name()
		}
	}

	// Sort files and directories lexicographically
	// Case sensitivity is NOT taken in cosideration, as ls does
	sort.Slice(fileList, func(i, j int) bool {
		return strings.ToLower(fileList[i].DefaultList) < strings.ToLower(fileList[j].DefaultList)
	})

	return fileList
}

func RetrieveHardLinkCount(path string) (int, error) {

	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		err = errors.New("couldn't get raw syscall.Stat_t data from" + path)
		return 0, err
	}
	hardLinks := stat.Nlink

	return hardLinks, err
}
