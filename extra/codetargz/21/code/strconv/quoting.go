package main

import (
	"log"
	"strconv"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	str := `

    "wat"

`
	log.Println(strconv.Quote(str))
	log.Println(strconv.QuoteRune(7))             // ASCII bell
	log.Println(strconv.QuoteRuneToASCII(0x30f0)) // ヰ
	log.Println(strconv.QuoteToASCII("ヰ"))
	log.Println(strconv.Unquote(`\n\r\t`)) // invalid due to lack of quotes
	log.Println(strconv.Unquote(`"\n\r\t"`))
}
