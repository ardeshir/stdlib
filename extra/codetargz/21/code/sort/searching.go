package main

import (
	"log"
	"sort"
)

func searchInts(needle int) {
	haystack := []int{1, 4, 7, 9, 10, 66}
	n := len(haystack)
	index := sort.Search(n, func(i int) bool {
		return haystack[i] >= needle
	})
	if index == n {
		log.Printf("didn't find %d", needle)
	} else {
		log.Printf("maybe found %d at index %d", needle, index)
	}
}

func main() {
	searchInts(9)
	searchInts(11)
	searchInts(70)
}
