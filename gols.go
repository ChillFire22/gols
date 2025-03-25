package gols

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"
)

const (
	INITIAL_SLICE_SIZE = 128
)

type IndexedEntry struct {
	Index     int
	Path      string
	createdAt time.Time
}

type (
	IndexList []IndexedEntry // faster for reading by index, which is my case
	IndexMap  map[int]IndexedEntry
)

func GetPath(index int, list IndexList, imap IndexMap) (path string) {
	if index < 0 {
		return ""
	}
	switch {
	case len(list) != 0:
		return (list[index]).Path
	case len(imap) != 0:
		return (imap[index]).Path
	default:
		return ""
	}
}

func ListToMap(list IndexList) IndexMap {
	imap := make(IndexMap, len(list))
	for i, e := range list {
		imap[i] = e
	}
	return imap
}

func MapToList(imap IndexMap) IndexList {
	list := make(IndexList, len(imap))
	for i, e := range imap {
		list[i] = e
	}
	return list
}

// Returns an IndexList populated with IndexedEntry-s created from files
// The underlying size of the returned IndexList is len(files)
func NewList(files []string) IndexList {
	flen := len(files)
	list := make(IndexList, flen)
	for i, path := range files {
		entry := NewIndexedEntry(path, i)
		list[i] = entry
	}
	return list
}

func NewMap(files []string) IndexMap {
	flen := len(files)
	imap := make(IndexMap, flen)
	for i, path := range files {
		entry := NewIndexedEntry(path, i)
		imap[i] = entry
	}
	return imap
}

func NewIndexedEntry(path string, index int) IndexedEntry {
	return IndexedEntry{
		Index:     index,
		Path:      path,
		createdAt: time.Now(),
	}
}

func isPathGood(dirpath string) (bool, error) {
	// path := filepath.Clean(dirpath)
	return true, nil
}

func getPaths(dirpath string) ([]string, error) {
	files := make([]string, INITIAL_SLICE_SIZE)
	var i int = 0

	err := filepath.WalkDir(dirpath, func(path string, e fs.DirEntry, err error) error {
		if e.Type().IsRegular() || e.IsDir() {
			files[i] = path
			i++
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

// String returns a custom string representation of all the content of list.
// Each newline represents an entry in the list
func (list IndexList) String() string {
	var out string = ""
	out = fmt.Sprintf("|Index)\tPath")
	if len(list) == 0 {
		out += "-"
		return out
	}
	for _, entry := range list {
		out += fmt.Sprintf("|%v)\t%s\n", entry.Index, entry.Path)
	}
	return out
}

// String returns a custom string representation of entry.
func (entry IndexedEntry) String() string {
	return entry.Path
}
