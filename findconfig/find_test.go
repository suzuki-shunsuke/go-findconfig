package findconfig_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-findconfig/findconfig"
)

func TestFind(t *testing.T) {
	t.Parallel()
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			p := findconfig.Find(d.wd, findconfig.NewMockExist(d.files...), d.names...)
			if p != d.exp {
				t.Fatalf(`Find() = "%s", wanted "%s"`, p, d.exp)
			}
		})
	}
}

func TestFinds(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		wd    string
		exp   []string
		names []string
		files []string
	}{
		{
			title: "normal",
			wd:    "/home/foo/workspace",
			exp:   []string{"/home/foo/foo.json"},
			names: []string{".foo.json", "foo.json"},
			files: []string{
				"/home/foo/foo.json",
				"/home/foo/bar.txt",
			},
		},
		{
			title: "not found",
			wd:    "/home/foo/workspace",
			exp:   nil,
			names: []string{"foo.json"},
			files: []string{
				"/home/foo/bar.txt",
			},
		},
		{
			title: "multiple",
			wd:    "/home/foo/workspace",
			exp:   []string{"/home/foo/foo.json", "/home/.foo.json", "/foo.json"},
			names: []string{".foo.json", "foo.json"},
			files: []string{
				"/home/foo/foo.json",
				"/home/foo/bar.txt",
				"/home/.foo.json",
				"/foo.json",
			},
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			p := findconfig.Finds(d.wd, findconfig.NewMockExist(d.files...), d.names...)
			if !reflect.DeepEqual(d.exp, p) {
				t.Fatalf(`Finds() = "%+v", wanted "%+v"`, p, d.exp)
			}
		})
	}
}
