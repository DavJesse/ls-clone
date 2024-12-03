package internal

type FileInfo struct {
	Index         string
	DocName       string
	DocPerm       string
	RecursiveList []FileInfo
	PlusHidden    string
	ReverseList   string
	ModTime       string
}

type ReverseAlpha []FileInfo
type Alphabetic []FileInfo
type ByTime []FileInfo

type MetaData struct {
	HardLinkCount int
	UserID        string
	GroupID       string
	Block int64
}

type DirFile struct {
	Dir   string
	Files []string
}

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

// Give sort.Sort interface size for sorting
func (f ReverseAlpha) Len() int {
	return len(f)
}

// Give sorting algoriths parameter for sorting
func (f ReverseAlpha) Less(i, j int) bool {
	return f[i].Index > f[j].Index
}

// Handle swapping
func (f ReverseAlpha) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Give sort.Sort interface size for sorting
func (f ByTime) Len() int {
	return len(f)
}

// Give sorting algoriths parameter for sorting
func (f ByTime) Less(i, j int) bool {
	return f[i].ModTime > f[j].ModTime
}

// Handle swapping
func (f ByTime) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
