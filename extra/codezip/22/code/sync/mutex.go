package main

import (
	"log"
	"runtime"
	"sync"
)

// Regular Mutex
type Lockable struct {
	m sync.Mutex
	n int
}

func (l *Lockable) Set(i int) {
	l.m.Lock()
	defer l.m.Unlock()
	l.n = i
}

func (l *Lockable) Get() int {
	l.m.Lock()
	defer l.m.Unlock()
	return l.n
}

// RWMutex
type RWLockable struct {
	m sync.RWMutex
	n int
}

func (l *RWLockable) Set(i int) {
	l.m.Lock()
	defer l.m.Unlock()
	l.n = i
}

func (l *RWLockable) Get() int {
	l.m.RLock()
	defer l.m.RUnlock()
	return l.n
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
	runtime.GOMAXPROCS(8)
}

func main() {
	l := &Lockable{}
	l.Set(10)
	log.Println(l.Get())

	rwl := &RWLockable{}
	rwl.Set(5)
	log.Println(rwl.Get())
}
