package main

import (
	"flag"
	"fmt"
	"log"
)

type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("%+d@%+d", p.X, p.Y)
}

func (p *Point) Set(s string) error {
	_, err := fmt.Sscanf(s, "%d@%d", &p.X, &p.Y)
	return err
}

var point Point

func init() {
	flag.Var(&point, "point", "point as X@Y")
}

func main() {
	flag.Parse()
	log.Printf("%#v", point)
}
