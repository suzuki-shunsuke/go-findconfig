package findconfig_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-findconfig/findconfig"
)

func TestFind(t *testing.T) {
	data := []struct {
		title string
		wd    string
		exp   string
		names []string
		files []string
	}{
		{
			title: "normal",
			wd:    "/home/foo/workspace",
			exp:   "/home/foo/foo.json",
			names: []string{".foo.json", "foo.json"},
			files: []string{
				"/home/foo/foo.json",
				"/home/foo/bar.txt",
			},
		},
		{
			title: "not found",
			wd:    "/home/foo/workspace",
			exp:   "",
			names: []string{"foo.json"},
			files: []string{
				"/home/foo/bar.txt",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			p := findconfig.Find(d.wd, findconfig.NewMockExist(d.files...), d.names...)
			if p != d.exp {
				t.Fatalf(`Find() = "%s", wanted "%s"`, p, d.exp)
			}
		})
	}
}
