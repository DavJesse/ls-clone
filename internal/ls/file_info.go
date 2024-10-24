// This file will contain the logic for retrieving file and directory metadata.
// It uses Go’s os and syscall packages (since the os/exec package is not allowed).
// Functions here will retrieve file sizes, permissions, timestamps, etc.
package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

func RetrieveFileInfo(path string) []FileInfo {
	var ResultList []FileInfo
	var doc FileInfo
	// var fileMetaData MetaData
	var linkCount int
	var userID, groupID string

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
			fileMetaData, err := RetrieveMetaData(path + "/" + entry.Name())
			if err != nil {
				log.Fatal(err)
			}
			linkCount = fileMetaData.HardLinkCount
			userID = fileMetaData.UserID
			groupID = fileMetaData.GroupID
		}
		// ignore git directories
		if strings.Contains(entry.Name(), ".git") {
			continue
		}

		if entry.IsDir() {
			// Append result for windows systems
			// Wrap text in bright blue
			if system == "windows" {
				doc.Index = fmt.Sprintf("%v\\", strings.ToLower(entry.Name()))
				doc.DocName = "\033[01;34m" + entry.Name() + "\033[0m" + "\\"
				doc.ModTime = entry.ModTime().String()
				doc.DocPerm = fmt.Sprintf("%v '-' '-' '-' %d %s \033[01;34m%v\033[0m//\n", entry.Mode().Perm().String(), entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())

				// Append 'doc' to fileList
				ResultList = append(ResultList, doc)

				// Append result for other systems
			} else {
				doc.Index = fmt.Sprintf("%v/", strings.ToLower(entry.Name()))
				doc.DocName = fmt.Sprintf("\033[01;34m%v\033[0m/", entry.Name())
				doc.ModTime = entry.ModTime().String()
				doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s \033[01;34m%v\033[0m/", entry.Mode().Perm().String(), linkCount, userID, groupID, entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())

				// Append 'doc' to fileList
				ResultList = append(ResultList, doc)

			}

			// Append result for file types
		} else {
			doc.Index = fmt.Sprintf("%v", strings.ToLower(entry.Name()))
			doc.ModTime = entry.ModTime().String()
			if IsExecutable(entry) {
				doc.DocName = fmt.Sprintf("\033[01;32m%s\033[0m*", entry.Name())
				doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s %v", entry.Mode().Perm().String(), linkCount, userID, groupID, entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())
			} else {
				doc.DocName = entry.Name()
				doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s %v", entry.Mode().Perm().String(), linkCount, userID, groupID, entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())
			}

			// Append 'doc' to fileList
			ResultList = append(ResultList, doc)
		}
	}

	// Sort files and directories lexicographically
	// Case sensitivity is NOT taken in cosideration, as ls does
	sort.Sort(Alphabetic(ResultList))

	return ResultList
}

func RetrieveMetaData(path string) (MetaData, error) {
	var result MetaData

	info, err := os.Lstat(path)
	if err != nil {
		return result, err
	}

	// Extract metadata from syscall.Stat_t
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		err = errors.New("couldn't get raw syscall.Stat_t data from" + path)
		return result, err
	}
	result.HardLinkCount = int(stat.Nlink)
	groupID := strconv.Itoa(int(stat.Gid))
	userID := strconv.Itoa(int(stat.Uid))

	// Extract user
	u, err1 := user.LookupId(userID)
	if err1 != nil {
		return result, err1
	}

	// Extract group
	g, err2 := user.LookupGroupId(groupID)
	if err2 != nil {
		return result, err2
	}

	result.UserID = u.Username
	result.GroupID = g.Name

	return result, err
}

func IsExecutable(fileInfo os.FileInfo) bool {
	mode := fileInfo.Mode()
	return mode&0o100 != 0 || mode&0o010 != 0 || mode&0o001 != 0
}
