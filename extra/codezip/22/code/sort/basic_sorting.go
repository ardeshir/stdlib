package main

import (
	"log"
	"sort"
)

type Question struct {
	Q, A           string
	PositionOnExam int
}

type Exam []Question

func (e Exam) Len() int {
	return len(e)
}

func (e Exam) Less(i, j int) bool {
	return e[i].PositionOnExam < e[j].PositionOnExam
}

func (e Exam) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func sortInts() {
	i := []int{5, 2, 9, 8, 7}
	log.Println(i, sort.IntsAreSorted(i))
	sort.Ints(i)
	log.Println(i, sort.IntsAreSorted(i))
}

func sortStrings() {
	s := []string{"robin", "batman", "thor", "loki", "captain america"}
	log.Println(s, sort.StringsAreSorted(s))
	sort.Strings(s)
	log.Println(s, sort.StringsAreSorted(s))
}

func sortFloats() {
	f := []float64{1.5, 2.3, 0.5, 0.4}
	log.Println(f, sort.Float64sAreSorted(f))
	sort.Float64s(f)
	log.Println(f, sort.Float64sAreSorted(f))
}

func sortCustomCollection() {
	exam := Exam{
		{Q: "How much wood...", A: "A lot", PositionOnExam: 4},
		{Q: "When did WWII start?", A: "1939", PositionOnExam: 5},
		{Q: "What color is the sky?", A: "Blue", PositionOnExam: 2},
		{Q: "Who builds the iPhone?", A: "Apple", PositionOnExam: 1},
		{Q: "Why is Go awesome?", A: "Lots of reasons", PositionOnExam: 3},
	}
	log.Println(exam, sort.IsSorted(exam))
	sort.Sort(exam)
	log.Println(exam, sort.IsSorted(exam))
}

func main() {
	sortInts()
	sortStrings()
	sortFloats()
	sortCustomCollection()
}
