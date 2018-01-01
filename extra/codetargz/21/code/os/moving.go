package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	flag.Parse()
}

func RunInDir(dir string, f func(string)) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("failed getting absolute directory path: %s", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %s", err)
	}

	os.Chdir(dir)
	defer os.Chdir(cwd)
	f(dir)
}

func Dir() []os.FileInfo {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}
	return files
}

func main() {
	f := func(cwd string) {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatalf("failed reading directory: %s", err)
		}
		log.Printf("found %d files in %s", len(files), cwd)
	}

	RunInDir(".", f)
	RunInDir("..", f)
	RunInDir("../..", f)
	RunInDir("../log", f)
	RunInDir("../../..", f)
}
