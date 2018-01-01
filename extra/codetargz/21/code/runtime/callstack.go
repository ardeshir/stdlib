package main

import (
	"log"
	"runtime"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func PrintStack() {
	stack := make([]byte, 1024)
	i := runtime.Stack(stack, false)
	log.Printf("%s", stack[0:i])
}

func C() {
	for i := 0; i < 6; i++ {
		log.Println(runtime.Caller(i))
	}
}

func B() {
	C()
}

func A() {
	B()
}

func main() {
	PrintStack()
	A()
}
