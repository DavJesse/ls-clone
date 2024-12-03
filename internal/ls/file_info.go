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
	"sort"
	"strconv"
	"strings"
	"syscall"
)

func RetrieveFileInfo(path string, includeHidden, rootIncluded bool) []FileInfo {
	var ResultList []FileInfo
	var doc FileInfo
	var fileMetaData MetaData
	var linkCount int
	var userID, groupID string

	// Open directory/file for reading
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if !rootIncluded {
		doc.RecursiveList = append(doc.RecursiveList, retrieveRootInfo(path, includeHidden)...)
		ResultList = append(ResultList, doc)
		doc = FileInfo{}
		rootIncluded = true
	}

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

		if entry.IsDir() {
			// ignore hidden directories
			if IsHidden(entry.Name()) && !includeHidden {
				continue
			}

			doc.RecursiveList = RetrieveFileInfo(path+"/"+entry.Name(), includeHidden, rootIncluded)
			doc.Index = fmt.Sprintf("%v/", strings.ToLower(entry.Name()))
			doc.DocName = fmt.Sprintf("\033[01;34m%v\033[0m", entry.Name())
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

	return ResultList
}

func retrieveRootInfo(path string, includeHidden bool) []FileInfo {
	var ResultList []FileInfo
	var doc FileInfo
	var linkCount int
	var userID, groupID string

	dirInfo, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer dirInfo.Close()

	files, err := dirInfo.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// Skip hidden files and directories
		if IsHidden(file.Name()) && !includeHidden {
			continue
		}

		metaData, err := RetrieveMetaData(path + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		// Retrieve root metadata
		linkCount = metaData.HardLinkCount
		userID = metaData.UserID
		groupID = metaData.GroupID

		if file.IsDir() {
			doc.DocName = file.Name()
		} else {
			doc.DocName = file.Name()
		}
		doc.Index = strings.ToLower(file.Name())
		doc.ModTime = file.ModTime().String()
		doc.DocPerm = fmt.Sprintf("%v %d %v %v %d %s %v", file.Mode().Perm().String(), linkCount, userID, groupID, file.Size(), file.ModTime().Format("Jan 02 15:04"), file.Name())
		ResultList = append(ResultList, doc)
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

func AddColor(file string, kind string) string {
	reset := "\033[0m"
	colorLibrary := map[string]string{
		"dir":             "\033[01;34m",
		"sym":             "\033[01;36m",
		"pipe":            "\033[33m",
		"soc":             "\033[01;35m",
		"exec":            "\033[01;32m",
		"bloc spec":       "\033[01;33m",
		"char spec":       "\033[01;33m",
		"set uid":         "\033[01;37;41m",
		"set gid":         "\033[01;30;43m",
		"dir stic bit":    "\033[01;34;42m",
		"dir nonstic bit": "\033[01;34;43m",
		"comp":            "\033[01;31mm",
		"img":             "\033[01;35m",
		"vid":             "\033[01;36m",
		"orph sym":        "\033[01;31m",
	}

	code, exists := colorLibrary[kind]

	if exists {
		return fmt.Sprintf("%s%s%s", code, file, reset)
	} else {
		return file
	}
}

func RemoveColor(file string) string {
	var start, end int

	if !ContainsColor(file) {
		return file
	}

	for i, v := range file {
		// mark character after first'm' as the start of file name
		if v == 'm' && start == 0 {
			start = i + 1
		}

		// mark character after last 'm' as the end of file name
		if v == '\x1b' && start != 0 {
			end = i
			break
		}
	}

	return file[start:end]
}

func ContainsColor(file string) bool {
	return strings.Contains(file, "\x1b")
}

func IsSymLink(file fs.FileInfo) bool {
	return file.Mode()&os.ModeSymlink != 0
}

func IsPipe(file fs.FileInfo) bool {
	return (file.Mode() & syscall.S_IFMT) == syscall.S_IFIFO
}

func IsSocket(file fs.FileInfo) bool {
	return (file.Mode() & syscall.S_IFMT) == syscall.S_IFSOCK
}

func IsBlockSpecial(file fs.FileInfo) bool {
	return (file.Mode() & syscall.S_IFMT) == syscall.S_IFBLK
}

func IsCharacterSpecial(file fs.FileInfo) bool {
	return (file.Mode() & syscall.S_IFMT) == syscall.S_IFCHR
}

func IsSetUserIDSet(file fs.FileInfo) bool {
	return file.Mode()&os.ModeSetuid != 0
}

func IsSetGroupIDSet(file fs.FileInfo) bool {
	return file.Mode()&os.ModeSetgid != 0
}

func IsStickyBitSet(file fs.FileInfo) bool {
	return file.IsDir() && file.Mode()&os.ModeSticky != 0
}

func IsStickyBitNotSet(file fs.FileInfo) bool {
	return file.IsDir() && file.Mode()&os.ModeSticky == 0
}

func IsCompressed(file fs.FileInfo) bool {
	extensions := []string{".zip", ".gz", ".bz2", ".7z", ".xz", ".tar.gz"}

	// Check file extension
	for i := range extensions {
		if strings.HasSuffix(file.Name(), extensions[i]) {
			return true
		}
	}
	return false
}

func IsImage(file fs.FileInfo) bool {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", "tif", ".ico", "webp"}

	// Check file extension
	for i := range extensions {
		if strings.HasSuffix(file.Name(), extensions[i]) {
			return true
		}
	}
	return false
}
