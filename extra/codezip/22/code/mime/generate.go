package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func Must(err error) {
	if err != nil {
		log.Fatalf("WriteField failed: %s", err)
	}
}

func WriteFile(w io.Writer, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatalf("failed writing file: %s", err)
	}
}

func Generate(w io.Writer) string {
	wr := multipart.NewWriter(w)
	defer wr.Close()
	Must(wr.WriteField("book", "Go, The Standard Library"))
	Must(wr.WriteField("chapter", "mime"))
	Must(wr.WriteField("examples", "2"))
	ff, err := wr.CreateFormFile("uploaded", "generate.go")
	if err != nil {
		log.Fatalf("failed creating form file: %s", err)
	}
	WriteFile(ff, "generate.go")
	return wr.Boundary()
}

func Parse(r io.Reader, boundary string) {
	rd := multipart.NewReader(r, boundary)
	form, err := rd.ReadForm(1024 * 1024 * 1024)
	if err != nil {
		log.Fatalf("failed reading form: %s", err)
	}

	for name, value := range form.Value {
		log.Printf("got form data %s: %s", name, value)
	}

	for name, fhs := range form.File {
		for _, fh := range fhs {
			log.Printf("got form file %s: %s", name, fh.Filename)
		}
	}
}

func main() {
	var buffer bytes.Buffer
	boundary := Generate(&buffer)
	log.Println(buffer.String())
	Parse(&buffer, boundary)
}
