package main

import (
	"log"
	"strconv"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func parseBools(strings ...string) {
	for _, s := range strings {
		b, err := strconv.ParseBool(s)
		log.Printf("%t, %s", b, err)
	}
}

func printBool(bools ...bool) {
	for _, b := range bools {
		log.Println(strconv.FormatBool(b))
	}
}

func parseFloats(bitSize int, strings ...string) {
	for _, s := range strings {
		f, err := strconv.ParseFloat(s, bitSize)
		log.Printf("bitSize: %d, %#v => %f, %s", bitSize, s, f, err)
	}
}

func printFloat(f float64, fmt byte, prec, bitSize int) {
	s := strconv.FormatFloat(f, fmt, prec, bitSize)
	lfmt := "fmt: %q, prec: %2d, bitSize: %d => %s"
	log.Printf(lfmt, fmt, prec, bitSize, s)
}

var bitSizes = []int{32, 64}
var formats = []byte("efg")
var precisions = []int{5, 10, 15}

func printFloats(fs ...float64) {
	for _, f := range fs {
		for _, fmt := range formats {
			for _, prec := range precisions {
				for _, bitSize := range bitSizes {
					printFloat(f, fmt, prec, bitSize)
				}
			}
		}
	}
}

func parseInts(base, bitSize int, ss ...string) {
	for _, s := range ss {
		i, err := strconv.ParseInt(s, base, bitSize)
		fmt := "base: %2d, bitSize: %2d, %#v => %d, %s"
		log.Printf(fmt, base, bitSize, s, i, err)
	}
}

func printInts(base int, is ...int64) {
	for _, i := range is {
		s := strconv.FormatInt(i, base)
		log.Printf("base: %2d, %d => %#v", base, i, s)
	}
}

func DemoBool() {
	log.Println("DemoBool")

	parseBools("true", "1", "f", "wat")
	printBool(true, false)
}

func DemoFloat() {
	log.Println("DemoFloat")

	parseFloats(32, "1.0", "-1.5", "1e10", "wat", "4e38")
	parseFloats(64, "4e38")

	printFloats(1.1234567, 4e38)
}

func DemoInt() {
	log.Println("DemoInt")

	big := "1010101010101010101010101010101010101010"
	parseInts(2, 32, "101101010", "10", "8", big)
	parseInts(2, 64, big)
	parseInts(8, 8, "12345", "7")
	parseInts(10, 32, "12345", "7")
	parseInts(16, 32, "abcdef")
	// Detect base based on prefix
	parseInts(0, 32, "0xff", "0644", "255")

	printInts(2, 100)
	printInts(3, 100)
	printInts(4, 100)
	printInts(5, -100)
	printInts(10, 100)
	printInts(16, 1250)
}

func main() {
	DemoBool()
	DemoFloat()
	DemoInt()
}
