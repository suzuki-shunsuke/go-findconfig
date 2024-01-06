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
	s, _ := find(wd, exist, names...)
	return s
}

func find(wd string, exist func(string) bool, names ...string) (string, string) {
	for {
		for _, name := range names {
			p := filepath.Join(wd, name)
			if exist(p) {
				return p, wd
			}
		}

		// Check in the .config subdirectory of working directory to support adapted XDG Base Directory Specification defined here https://github.com/dot-config/dot-config.github.io.
		// Only if no file is found in the working directory.
		configPath := filepath.Join(wd, ".config")
		for _, name := range names {
			p := filepath.Join(configPath, name)
			if exist(p) {
				return p, wd
			}
		}

		parent := filepath.Dir(wd)
		if wd == parent {
			return "", ""
		}
		wd = parent
	}
}

func Finds(wd string, exist func(string) bool, names ...string) []string {
	var files []string
	for {
		p, foundDir := find(wd, exist, names...)
		if p == "" {
			return files
		}
		files = append(files, p)
		parent := filepath.Dir(foundDir)
		if wd == parent {
			return files
		}
		wd = parent
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
