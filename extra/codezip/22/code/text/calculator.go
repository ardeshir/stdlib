package main

import (
	"flag"
	"log"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

var (
	equation string
	numberRe = regexp.MustCompile(`-?[1-9][0-9]*(\.[0-9]+)?`)
)

func fn(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func show(l, r, value float64, operand string) {
	log.Printf("pushing %s %s %s => %s", fn(l), operand, fn(r), fn(value))
}

type Stack struct {
	data []string
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

func (s *Stack) PushNumber(num float64) {
	s.Push(fn(num))
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	value, data := s.data[len(s.data)-1], s.data[:len(s.data)-1]
	s.data = data
	return value
}

func (s *Stack) PopNumber() float64 {
	value := s.Pop()
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("failed parsing number: %s", err)
	}
	return num
}

func (s *Stack) PopOperands() (float64, float64) {
	r, l := s.PopNumber(), s.PopNumber()
	return l, r
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")

	flag.StringVar(&equation, "rpn", "1 2 + 3 * 2 / 10 -", "the equation to evaluate")
	flag.Parse()
}

func main() {
	var s scanner.Scanner
	s.Filename = "equation"
	s.Init(strings.NewReader(equation))

	stack := Stack{}
	for {
		// Using Scan() we skip whitespace
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}
		text := s.TokenText()
		switch tok {
		case '+':
			l, r := stack.PopOperands()
			value := l + r
			show(l, r, value, "+")
			stack.PushNumber(value)
		case '-':
			l, r := stack.PopOperands()
			value := l - r
			show(l, r, value, "-")
			stack.PushNumber(value)
		case '*':
			l, r := stack.PopOperands()
			value := l * r
			show(l, r, value, "*")
			stack.PushNumber(value)
		case '/':
			l, r := stack.PopOperands()
			value := l / r
			show(l, r, value, "/")
			stack.PushNumber(value)
		default:
			switch {
			case numberRe.MatchString(text):
				log.Printf("pushing %s", text)
				stack.Push(text)
			}
		}
	}
	log.Printf("=> %s", stack.Pop())
}
