package internal

type FileInfo struct {
	Index         string
	DocName       string
	DocPerm       string
	RecursiveList string
	PlusHidden    string
	ReverseList   string
	ModTime       string
}

type Alphabetic []FileInfo

// Give sort.Sort interface size for sorting
func (f Alphabetic) Len() int {
	return len(f)
}

// Give sorting algoriths parameter for sorting
func (f Alphabetic) Less(i, j int) bool {
	return f[i].Index < f[j].Index
}

// Handle swapping
func (f Alphabetic) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
