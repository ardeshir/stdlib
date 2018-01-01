package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

var files = []string{"write_zip.go", "read_zip.go"}

func addFile(filename string, zw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed opening %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zw.Create(filename)
	if err != nil {
		msg := "failed creating entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	// Not checking how many bytes copied,
	// since we don't know the file size without doing more work
	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed writing %s to zip: %s", filename, err)
	}

	return nil
}

func main() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("go.zip", flags, 0644)
	if err != nil {
		log.Fatalf("failed opening zip for writing: %s", err)
	}
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	for _, filename := range files {
		if err := addFile(filename, zw); err != nil {
			log.Fatalf("failed adding file %s to zip: %s", filename, err)
		}
	}
}
