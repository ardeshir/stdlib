package main

import (
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf16"
)

var (
	proverb = "Alla är vi barn i början."
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

type UTF16Writer struct {
	out     io.Writer
	bom     bool
	started bool
	buf     *bytes.Buffer
}

func NewUTF16Writer(out io.Writer, bom bool) *UTF16Writer {
	return &UTF16Writer{
		out: out,
		bom: bom,
		buf: new(bytes.Buffer),
	}
}

func (w *UTF16Writer) Write(p []byte) (int, error) {
	if !w.started {
		if w.bom {
			// We're assuming little endian, since that's what Excel wants,
			// but you could easily pass in a endianess.
			_, err := w.out.Write([]byte{'\xff', '\xfe'})
			if err != nil {
				return 0, err
			}
		}
		w.started = true
	}

	_, err := w.buf.Write(p)
	if err != nil {
		return 0, err
	}

	// omg such a hack
	for {
		r, s, err := w.buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				return len(p), nil
			}
			return 0, err
		}

		// The lazy hack
		if r == unicode.ReplacementChar && s == 1 {
			return 0, errors.New("incomplete rune")
		}

		err = binary.Write(w.out, binary.LittleEndian, utf16.Encode([]rune{r}))
		if err != nil {
			return 0, err
		}
	}
}

func main() {
	proverbs := [][]string{
		{"Language", "Proverb"},
		{"sv", "Alla är vi barn i början."},
		{"zh", "读书须用意，一字值千金"},
	}
	csvWriter := csv.NewWriter(NewUTF16Writer(os.Stdout, true))
	csvWriter.Comma = '\t'
	err := csvWriter.WriteAll(proverbs)
	if err != nil {
		log.Fatalf("failed writing: %s", err)
	}
}
