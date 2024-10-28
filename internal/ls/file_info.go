// This file will contain the logic for retrieving file and directory metadata.
// It uses Go’s os and syscall packages (since the os/exec package is not allowed).
// Functions here will retrieve file sizes, permissions, timestamps, etc.
package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

func RetrieveFileInfo(path string, includeHidden bool) []FileInfo {
	var ResultList []FileInfo
	var doc FileInfo
	var fileMetaData MetaData
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

	// Append root directory to fileList
	doc = AppendRootDir(entries, doc, includeHidden)
	ResultList = append(ResultList, doc)

	// Retrieve directory/file name and append to fileList
	// For directories, we add '/' or '\' depending on opperating system
	for _, entry := range entries {
		if system != "windows" {
			fileMetaData, err = RetrieveMetaData(path + "/" + entry.Name())
			if err != nil {
				log.Fatal(err)
			}
			linkCount = fileMetaData.HardLinkCount
			userID = fileMetaData.UserID
			groupID = fileMetaData.GroupID
		}

		if entry.IsDir() {
			doc.RecursiveList = fmt.Sprintf("\033[01;34m%s\033[0m/\n", entry.Name())
			err := filepath.WalkDir(filepath.Join(path, entry.Name()), func(subPath string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				// Skip hidden files and directories if not included
				if !includeHidden && !IsHidden(d.Name()) {
					if d.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}

				// Skip the root directory itself
				if subPath == filepath.Join(path, entry.Name()) {
					return nil
				}

				// relPath, _ := filepath.Rel(path, subPath)
				if d.IsDir() {
					doc.RecursiveList += fmt.Sprintf("\033[01;34m%s\033[0m:\n", d.Name())
				} else {
					doc.RecursiveList += d.Name() + " "
				}

				return nil
			})

			if err != nil {
				log.Printf("Error walking directory %s: %v\n", entry.Name(), err)
			}
			// Append result for windows systems
			// Wrap text in bright blue
			if system == "windows" {
				// ignore hidden directories
				if entry.Name()[0] == '.' && !includeHidden {
					continue
				}
				doc.Index = fmt.Sprintf("%v\\", strings.ToLower(entry.Name()))
				doc.DocName = "\033[01;34m" + entry.Name() + "\033[0m" + "\\"
				doc.ModTime = entry.ModTime().String()
				doc.DocPerm = fmt.Sprintf("%v '-' '-' '-' %d %s \033[01;34m%v\033[0m//", entry.Mode().Perm().String(), entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())

				// Append 'doc' to fileList
				ResultList = append(ResultList, doc)

				// Append result for other systems
			} else {
				// ignore hidden directories
				if !IsHidden(entry.Name()) && !includeHidden {
					continue
				}

				doc.Index = fmt.Sprintf("%v/", strings.ToLower(entry.Name()))
				doc.DocName = fmt.Sprintf("\033[01;34m%v\033[0m/", entry.Name())
				doc.ModTime = entry.ModTime().String()
				doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s \033[01;34m%v\033[0m/", entry.Mode().Perm().String(), linkCount, userID, groupID, entry.Size(), entry.ModTime().Format("Jan 02 15:04"), entry.Name())

				// Append 'doc' to fileList
				ResultList = append(ResultList, doc)

			}

			// Append result for file types
		} else {
			// ignore hidden files
			if entry.Name()[0] == '.' && !includeHidden {
				continue
			}

			doc.RecursiveList = ""

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

// Check if file on a given path is hidden.
func IsHidden(path string) bool {
	if strings.HasPrefix(filepath.Base(path), ".") {
		if len(path) == 2 && path[1] != '.' {
			return true
		}
	}
	return false
}

// Updates doc with the contents of the root directory.
func AppendRootDir(fileList []fs.FileInfo, doc FileInfo, includeHidden bool) FileInfo {
	doc.Index = "."

	if includeHidden {
		doc.RecursiveList = ". .. "
	}
	for i := range fileList {
		if i != 0 {
			doc.RecursiveList += fileList[i].Name()
		} else {
			doc.RecursiveList += fileList[i].Name() + " "
		}
	}
	files := strings.Fields(doc.RecursiveList)
	log.Println(files)
	sort.Strings(files)
	doc.RecursiveList = strings.Join(files, " ")
	return doc
}
