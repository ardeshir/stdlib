// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"sort"
)

type earthMass float64
type au float64

type Planet struct {
	name     string
	mass     earthMass
	distance au
}

type By func(p1, p2 *Planet) bool

func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by,
	}
	sort.Sort(ps)
}

type planetSorter struct {
	planets []Planet
	by      By
}

func (s *planetSorter) Len() int {
	return len(s.planets)
}

func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}

func main() {
	// Closures that order the Planet structure.
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	decreasingDistance := func(p1, p2 *Planet) bool {
		return !distance(p1, p2)
	}

	// Sort the planets by the various criteria.
	By(name).Sort(planets)
	log.Println("By name:", planets)

	By(mass).Sort(planets)
	log.Println("By mass:", planets)

	By(distance).Sort(planets)
	log.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)
	log.Println("By decreasing distance:", planets)
}
