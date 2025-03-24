package gols

import (
	"fmt"
	"time"
)

type IndexedEntry struct {
	Index     int
	Path      string
	createdAt time.Time
}

type IndexList map[IndexedEntry]struct{} // a set

// Index reshuffles indexes of all il 's IndexedEntries in ascending order and returns Go int representing the last and highest index.
// The first index starts at 1, unless il is empty.
// Returns 0 if il is empty.
func (il *IndexList) Index() (topIndex int) {
	if len(*il) == 0 {
		return 0
	}
	var index int = 1
	for entry := range *il {
		entry.Index = index
		index++
	}
	return index
}

// String returns a custom string representation of all the content of the IndexList.
// Each newline begins with an Index then a Path representing the IndexedEntry
// Example
// |Index)	Path
// |1)	SomePath
func (il *IndexList) String() string {
	var out string = ""
	out = fmt.Sprintf("|Index)\tPath")
	if len(*il) == 0 {
		out += "-"
		return out
	}
	for entry := range *il {
		out += fmt.Sprintf("|%v)\t%s\n", entry.Index, entry.Path)
	}
	return out
}
// String returns a custom string representation of IndexedEntry.
func (ie IndexedEntry) String() string {
	return ie.Path
}
