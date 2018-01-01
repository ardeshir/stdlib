package main

import (
	"flag"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	flag.Parse()
}

func must(err error) {
	if err != nil {
		log.Fatalf("failed operation: %s", err)
	}
}

func DemoMkdir() {
	must(os.MkdirAll("foo/bar/baz", 0755))
	must(os.Mkdir("example", 0755))
}

func CleanupDir() {
	must(os.RemoveAll("foo"))
	must(os.Remove("example"))
}

func DemoLink() {
	must(os.Symlink("Makefile", "Makefile-symlink"))
	must(os.Link("Makefile", "Makefile-link"))
}

func CleanupLink() {
	must(os.Remove("Makefile-symlink"))
	must(os.Remove("Makefile-link"))
}

func DemoRename() {
	must(os.Rename("Makefile", "makefile"))
}

func CleanupRename() {
	must(os.Rename("makefile", "Makefile"))
}

func DemoTruncate() {
	// Look at the size of Makefile after this
	// Content hasn't changed, but it's magically 1kb
	must(os.Truncate("Makefile", 1024))
}

func CleanupTruncate() {
	must(os.Truncate("Makefile", 315))
}

func main() {
	DemoMkdir()
	CleanupDir()
	DemoLink()
	CleanupLink()
	DemoRename()
	CleanupRename()
	DemoTruncate()
	CleanupTruncate()
}
