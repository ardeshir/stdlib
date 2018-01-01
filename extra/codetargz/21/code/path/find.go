package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type FilterFunc func(path string, info os.FileInfo, err error) bool
type FilterChain []FilterFunc

var (
	root                 string
	ftype, name          string
	printNewline, print0 bool
	filters              FilterChain

	output = func(s string) {}
)

func init() {
	flag.StringVar(&ftype, "type", "", "f for file, d for directory")
	flag.StringVar(&name, "name", "", "find files/directories that match")
	flag.BoolVar(&printNewline, "print", false, "print elements to stdout with newlines separators")
	flag.BoolVar(&print0, "print0", false, "print elements to stdout with NULL separators")
	flag.Parse()
	root = flag.Arg(0)
	if root == "" {
		root = "."
	}
}

func setupPrinting() {
	if printNewline {
		output = func(s string) { fmt.Println(s) }
	} else if print0 {
		output = func(s string) { fmt.Printf("%s\x00", s) }
	} else {
		output = func(s string) { fmt.Println(s) }
	}
}

func nameFilter(path string, info os.FileInfo, err error) bool {
	matched, err := filepath.Match(name, filepath.Base(path))
	if err != nil {
		fmt.Printf("failed matching: %s", err)
		os.Exit(1)
	}
	return matched
}

func fileFilter(path string, info os.FileInfo, err error) bool {
	return !info.IsDir()
}

func directoryFilter(path string, info os.FileInfo, err error) bool {
	return info.IsDir()
}

func ok(path string, info os.FileInfo, err error) bool {
	return true
}

func setupFilters() {
	switch ftype {
	case "f":
		filters = append(filters, fileFilter)
	case "d":
		filters = append(filters, directoryFilter)
	}

	if name != "" {
		filters = append(filters, nameFilter)
	}

	if len(filters) == 0 {
		filters = append(filters, ok)
	}
}

func walker(path string, info os.FileInfo, err error) error {
	for _, filter := range filters {
		if !filter(path, info, err) {
			return nil
		}
	}
	output(path)
	return nil
}

func main() {
	setupPrinting()
	setupFilters()
	filepath.Walk(root, walker)
}
