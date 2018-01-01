package main

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoReaddir() {
	f, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer f.Close()

	fileInfos, err := f.Readdir(0)
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}

	for _, finfo := range fileInfos {
		log.Printf("Name: %s, Size: %db", finfo.Name(), finfo.Size())
	}
}

func DemoReaddirnames() {
	f, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer f.Close()

	names, err := f.Readdirnames(0)
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}

	for _, name := range names {
		log.Println(name)
	}
}

func main() {
	DemoReaddir()
	DemoReaddirnames()
}
