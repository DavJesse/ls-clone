// This handles the sorting functionality.
// Sorting is necessary for flags like -t (sort by modification time) and -r (reverse sorting).
// By separating this logic, we can ensure flexibility and make it easier to test the sorting independently.

package internal

func SortByEmptyDir(files []FileInfo) []FileInfo {
	point1 := 0
	point2 := point1 + 1

	for point1 < len(files) && point2 < len(files) {
		if files[point1].RecursiveList == nil && files[point2].RecursiveList == nil {
			point1 += 2
		} else if files[point1].RecursiveList == nil && files[point2].RecursiveList != nil {
			point1++
		} else if files[point1].RecursiveList != nil && files[point2].RecursiveList == nil {
			files[point1], files[point2] = files[point2], files[point1]
			point1++
		} else if files[point1].RecursiveList != nil && files[point2].RecursiveList != nil {
			point1 += 2
		}
	}
	return files
}
