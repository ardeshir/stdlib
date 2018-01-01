package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"os"
)

type Part struct {
	*multipart.Part
	Body []byte
}

func (p *Part) Reader() io.Reader {
	return bytes.NewReader(p.Body)
}

func ReadMultipartFile(path, boundary string) (parts []*Part) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening %s: %s", path, err)
	}
	defer file.Close()
	return ReadMultipart(file, boundary)
}

func ReadMultipart(r io.Reader, boundary string) (parts []*Part) {
	mr := multipart.NewReader(r, boundary)
	for {
		part, err := mr.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed reading part: %s", err)
		}
		body, err := ioutil.ReadAll(part)
		if err != nil {
			log.Fatalf("failed reading part: %s", err)
		}
		parts = append(parts, &Part{part, body})
		part.Close()
	}
	return parts
}

func DecodeBody(r io.Reader, encoding string) []byte {
	switch encoding {
	case "base64":
		dec := base64.NewDecoder(base64.StdEncoding, r)
		data, err := ioutil.ReadAll(dec)
		if err != nil {
			log.Fatalf("failed decoding: %s", err)
		}
		return data
	default:
		log.Fatalf("can't decode %s", encoding)
	}
	panic("not reached")
}

func DumpParts(parts []*Part, prefix string) {
	log.Printf("found %d parts", len(parts))
	for i, part := range parts {
		ctype := part.Header.Get("Content-Type")
		log.Printf(prefix+"part %d has Content-Type: %s", i+1, ctype)
		mtype, params, err := mime.ParseMediaType(ctype)
		if err != nil {
			log.Fatalf("failed parsing media type %s: %s", ctype, err)
		}
		switch mtype {
		case "text/plain", "text/html":
			log.Printf(prefix+"content: %s", part.Body)
		case "application/octet-stream":
			body := DecodeBody(part.Reader(), part.Header.Get("Content-Transfer-Encoding"))
			log.Printf(prefix+"decoded attachment with contents: %s", body)
		case "multipart/alternative":
			altParts := ReadMultipart(part.Reader(), params["boundary"])
			DumpParts(altParts, prefix+"\t")
		}
	}
}

func main() {
	parts := ReadMultipartFile("body", "047d7bae420e40e18a04e7e1ead4")
	DumpParts(parts, "")
}
