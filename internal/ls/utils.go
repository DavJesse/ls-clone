package internal

type FileInfo struct {
	DocName       string
	DocPerm       string
	RecursiveList string
	PlusHidden    string
	ReverseList   string
	ModTime       string
}

type FileList []FileInfo

// Give sort.Sort interface size for sorting
func (f FileList) Len() int {
	return len(f)
}

// Give sorting algoriths parameter for sorting
func (f FileList) Less(i, j int) bool {
	return f[i].DocName < f[j].DocName
}

// Handle swapping
func (f FileList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
