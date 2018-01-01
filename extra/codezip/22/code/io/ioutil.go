package main

import (
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func DemoReadAll() {
	file, err := os.Open("ioutil.go")
	if err != nil {
		log.Panicf("failed opening file: %s", err)
	}
	defer file.Close()
	log.Println(`reading file "ioutil.go"`)
	data, err := ioutil.ReadAll(file)
	log.Printf("read %d bytes with err %v", len(data), err)
}

func DemoReadDir() {
	entries, err := ioutil.ReadDir(".")
	if err != nil {
		log.Panicf("failed reading directory: %s", err)
	}
	log.Printf("found %d files in the current directory", len(entries))
}

func DemoReadFile() {
	data, err := ioutil.ReadFile("ioutil.go")
	log.Printf("read %d bytes with err %v", len(data), err)
}

func main() {
	DemoReadAll()
	DemoReadDir()
	DemoReadFile()
}
