// This file will contain the logic for retrieving file and directory metadata.
// It uses Go’s os and syscall packages (since the os/exec package is not allowed).
// Functions here will retrieve file sizes, permissions, timestamps, etc.
package internal

import (
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
)

func RetrieveFileInfo(path string) FileList {
	var ResultList FileList
	var doc FileInfo
	var linkCount string
	system := runtime.GOOS

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
		if system != "windows" {
			//hardLinks, err := RetrieveHardLinkCount(path + "/" + entry.Name())
			if err != nil {
				log.Fatal(err)
			}
			//	linkCount = strconv.Itoa(int(hardLinks))
		}
		// ignore git directories
		if strings.Contains(entry.Name(), ".git") {
			continue
		}

		if entry.IsDir() {
			// Append result for windows systems
			// Wrap text in bright blue
			if system == "windows" {

				doc.DocName = "\033[01;34m" + entry.Name() + "\033[0m" + "\\"
				doc.DocPerm = entry.Mode().Perm().String() + " " + "not available" + " " + "\033[01;34m" + entry.Name() + "\\\n"

				// Append 'doc' to fileList
				ResultList = append(ResultList, doc)

				// Append result for other systems
			} else {

				doc.DocName = "\033[01;34m" + entry.Name() + "\033[0m" + "/"
				doc.DocPerm = entry.Mode().Perm().String() + " " + linkCount + " " + "\033[01;34m" + entry.Name()

				// Append 'doc' to fileList
				ResultList = append(ResultList, doc)

			}

			// Append result for file types
		} else {

			doc.DocName = entry.Name()
			doc.DocPerm = entry.Mode().Perm().String() + " " + linkCount + " " + entry.Name()

			// Append 'doc' to fileList
			ResultList = append(ResultList, doc)
		}
	}

	// Sort files and directories lexicographically
	// Case sensitivity is NOT taken in cosideration, as ls does
	sort.Sort(ResultList)

	return ResultList
}

// func RetrieveHardLinkCount(path string) (uint64, error) {
// 	info, err := os.Lstat(path)
// 	if err != nil {
// 		return 0, err
// 	}

// 	stat, ok := info.Sys().(*syscall.Stat_t)
// 	if !ok {
// 		err = errors.New("couldn't get raw syscall.Stat_t data from" + path)
// 		return 0, err
// 	}
// 	hardLinks := stat.Nlink

// 	return hardLinks, err
// }
