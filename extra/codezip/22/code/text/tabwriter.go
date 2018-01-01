package main

import (
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	data := [][]string{
		{"Continent", "Country", "Nationality"},
		{"North America", "Canada", "Canadian"},
		{"Europe", "France", "French"},
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	defer writer.Flush() // Make sure to Flush the writer when you're done

	for _, tuple := range data {
		writer.Write([]byte(strings.Join(tuple, "\t")))
		writer.Write([]byte{'\n'})
	}
}
