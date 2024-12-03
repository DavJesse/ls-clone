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
	"sort"
	"strconv"
	"strings"
	"syscall"
)

func RetrieveFileInfo(path string, includeHidden bool, metaData *MetaData) []FileInfo {
	var ResultList []FileInfo
	var doc FileInfo
	var fileMetaData MetaData
	var linkCount int
	var userID, groupID string
	var blocks int64

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
		fileMetaData, err = RetrieveMetaData(path + "/" + entry.Name())
		if err != nil {
			log.Fatal(err)
		}
		linkCount = fileMetaData.HardLinkCount
		userID = fileMetaData.UserID
		groupID = fileMetaData.GroupID
		blocks += fileMetaData.Block

		if entry.IsDir() {
			// ignore hidden directories
			if IsHidden(entry.Name()) && !includeHidden {
				continue
			}

			doc.RecursiveList = RetrieveFileInfo(path+"/"+entry.Name(), includeHidden, metaData)
			doc.Index = fmt.Sprintf("%v/", strings.ToLower(entry.Name()))
			doc.DocName = fmt.Sprintf("\033[01;34m%v\033[0m/", entry.Name())
			doc.ModTime = entry.ModTime().String()
			doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s \033[01;34m%v\033[0m/", entry.Mode().Perm().String(), linkCount, userID, groupID, entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())

			// Append 'doc' to fileList
			ResultList = append(ResultList, doc)
			doc = FileInfo{}

			// Append result for file types
		} else {
			// ignore hidden files
			if IsHidden(entry.Name()) && !includeHidden {
				continue
			}

			// Add bright-green color to executable files
			// Retain default color for non-executable
			if IsExecutable(entry) {
				doc.DocName = fmt.Sprintf("\033[01;32m%s\033[0m*", entry.Name())
				doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s %v", entry.Mode().Perm().String(), linkCount, userID, groupID, entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())
			} else {
				doc.DocName = entry.Name()
				doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s %v", entry.Mode().Perm().String(), linkCount, userID, groupID, entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())
			}
			doc.Index = fmt.Sprintf("%v", strings.ToLower(entry.Name()))
			doc.ModTime = entry.ModTime().String()

			// Append 'doc' to fileList
			ResultList = append(ResultList, doc)
			doc = FileInfo{}
		}
	}

	// Sort files and directories lexicographically
	// Case sensitivity is NOT taken in cosideration, as ls does
	sort.Sort(Alphabetic(ResultList))

	// Update the global block count in the metadata
	if metaData != nil {
		metaData.Block += blocks
	}

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
	result.Block = stat.Blocks

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

// Check if file on a given path is hidden.
func IsHidden(path string) bool {
	return strings.HasPrefix(path, ".")
}
