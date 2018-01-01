package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	p          string
	walk       string
	ignore     string
	ignoreList []string
)

type Walker struct {
	NumDirs  int
	NumFiles int
}

func (w *Walker) Visit(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		base := filepath.Base(path)
		for _, dir := range ignoreList {
			if base == dir {
				return filepath.SkipDir
			}
		}
		w.NumDirs++
	} else {
		w.NumFiles++
	}
	return nil
}

func init() {
	flag.StringVar(&p, "path", "./foo/../baz.gif", "the path to examine")
	flag.StringVar(&walk, "walk", "..", "the path to walk")
	flag.StringVar(&ignore, "ignore", ".git:.hg", "directories to ignore")
	flag.Parse()

	ignoreList = filepath.SplitList(ignore)
}

func main() {
	log.Printf("p: %s", p)

	abs, err := filepath.Abs(p)
	log.Printf("Abs(p): %s, %v", abs, err)
	log.Printf("Base(p): %s", filepath.Base(p))
	log.Printf("Clean(p): %s", filepath.Clean(p))
	log.Printf("Dir(p): %s", filepath.Dir(p))

	sym, err := filepath.EvalSymlinks(p)
	log.Printf("EvalSymlinks(p): %s, %v", sym, err)
	log.Printf("Ext(p): %s", filepath.Ext(p))
	log.Printf("FromSlash(p): %s", filepath.FromSlash(p))

	glob, err := filepath.Glob("*.go")
	log.Printf("Glob(\"*.go\"): %s, %v", glob, err)
	log.Printf("IsAbs(p): %t", filepath.IsAbs(p))
	log.Printf("Join(\"/fizz/bin\", p): %s", filepath.Join("/fizz/bin", p))

	matched, err := filepath.Match("/*/bin/*.gif", p)
	log.Printf("Match(\"/*/bin/*.gif\", p): %t, %v", matched, err)

	matched, err = filepath.Match("/*/bin/*.gif", filepath.Join("/fizz/bin", p))
	log.Printf("Match(\"/*/bin/*.gif\", Join(\"/fizz/bin\", p)): %t, %v", matched, err)

	rel, err := filepath.Rel("/batman", "/path/file.go")
	log.Printf("Rel(\"/batman\", \"/path/file.go\"): %s, %v", rel, err)

	dir, file := filepath.Split(p)
	log.Printf("Split(p): %s, %s", dir, file)

	list := filepath.SplitList("/foo.go:/bar.go:/baz.go")
	log.Printf("SplitList(\"/foo.go:/bar.go:/baz.go\"): %s", list)

	var w Walker
	filepath.Walk("..", (&w).Visit)
	log.Printf("found %d directories and %d files", w.NumDirs, w.NumFiles)
}
