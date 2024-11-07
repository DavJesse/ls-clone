// This handles the sorting functionality.
// Sorting is necessary for flags like -t (sort by modification time) and -r (reverse sorting).
// By separating this logic, we can ensure flexibility and make it easier to test the sorting independently.

package internal

func SortByEmptyDir(files []FileInfo) []FileInfo {
	if len(files) <= 1 {
		return files
	}

	var emptDir, nonEmptDir []FileInfo
	for i := range files {
		if files[i].RecursiveList == nil {
			emptDir = append(emptDir, files[i])
		} else {
			nonEmptDir = append(nonEmptDir, files[i])
		}
	}
	return append(emptDir, nonEmptDir...)

}
