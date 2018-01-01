package main

import (
	"log"
	"math"
	"strconv"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	var data []byte
	data = strconv.AppendBool(data, true)
	log.Printf("%s", data)

	data = append(data, ',', ' ')
	data = strconv.AppendFloat(data, math.Pi, 'e', 2, 32)
	log.Printf("%s", data)

	data = append(data, ',', ' ')
	data = strconv.AppendInt(data, 42, 8)
	log.Printf("%s", data)

	data = append(data, ',', ' ')
	data = strconv.AppendQuote(data, `bat"man`)
	log.Printf("%s", data)

	data = append(data, ',', ' ')
	data = strconv.AppendQuoteRune(data, 0x30f0)
	log.Printf("%s", data)

	data = append(data, ',', ' ')
	data = strconv.AppendQuoteRuneToASCII(data, 0x30f0)
	log.Printf("%s", data)

	data = append(data, ',', ' ')
	data = strconv.AppendQuoteToASCII(data, "ヰヱ")
	log.Printf("%s", data)

	data = append(data, ',', ' ')
	data = strconv.AppendUint(data, 10, 2)
	log.Printf("%s", data)
}
