package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func printFile(file *zip.File) error {
	frc, err := file.Open()
	if err != nil {
		msg := "failed opening zip entry %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer frc.Close()

	fmt.Fprintf(os.Stdout, "Contents of %s:\n", file.Name)

	copied, err := io.Copy(os.Stdout, frc)
	if err != nil {
		msg := "failed reading zip entry %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	if uint64(copied) != file.UncompressedSize64 {
		msg := "read %d bytes of %s but expected to read %d bytes"
		return fmt.Errorf(msg, copied, file.UncompressedSize64)
	}

	fmt.Println()

	return nil
}

func main() {
	rc, err := zip.OpenReader("go.zip")
	if err != nil {
		msg := "failed opening archive, run `go run write_zip.go` first: %s"
		log.Fatalf(msg, err)
	}
	defer rc.Close()

	for _, file := range rc.File {
		if err := printFile(file); err != nil {
			log.Fatalf("failed reading %s from zip: %s", file.Name, err)
		}
	}
}
