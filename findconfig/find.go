package findconfig

import (
	"os"
	"path/filepath"
)

// Find finds file from wd to the root directory recursively.
// If no file is found, an empty string is returned.
// exist returns whether a file is found.
// names is a list of file names.
func Find(wd string, exist func(string) bool, names ...string) string {
	for {
		for _, name := range names {
			p := filepath.Join(wd, name)
			if exist(p) {
				return p
			}
		}
		if wd == "/" || wd == "" {
			return ""
		}
		wd = filepath.Dir(wd)
	}
}

func Finds(wd string, exist func(string) bool, names ...string) []string {
	var files []string
	for {
		p := Find(wd, exist, names...)
		if p == "" {
			return files
		}
		files = append(files, p)
		if wd == "/" || wd == "" {
			return files
		}
		wd = filepath.Dir(filepath.Dir(p))
	}
}

// Exist returns whether a file is found.
func Exist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// NewMockExist returns a mock function of Find's exist.
func NewMockExist(files ...string) func(p string) bool {
	return func(p string) bool {
		for _, f := range files {
			if f == p {
				return true
			}
		}
		return false
	}
}
