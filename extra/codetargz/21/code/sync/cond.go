package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(8)
}

type Status int

const (
	Absent Status = iota
	InProgress
	Complete
)

func getData(key string) []byte {
	// Do some work
	time.Sleep(3 * time.Second)
	return []byte(fmt.Sprintf("getData: %v", key))
}

type CacheEntry struct {
	sync.Mutex
	C      *sync.Cond
	Status Status
	Data   []byte
}

func (ce *CacheEntry) SetComplete(data []byte) {
	ce.Lock()
	defer ce.Unlock()
	ce.Data = data
	ce.Status = Complete
}

func (ce *CacheEntry) Wait() []byte {
	ce.Lock()
	defer ce.Unlock()
	for {
		if ce.Status == Complete {
			break
		}
		ce.C.Wait()
	}
	return ce.Data
}

type Cache struct {
	sync.RWMutex
	statuses map[string]Status
	data     map[string]*CacheEntry
}

func NewCache() *Cache {
	return &Cache{
		statuses: make(map[string]Status),
		data:     make(map[string]*CacheEntry),
	}
}

func (c *Cache) setComplete(key string) {
	c.Lock()
	defer c.Unlock()
	c.statuses[key] = Complete
}

func (c *Cache) setInProgress(key string) (*CacheEntry, bool) {
	c.Lock()
	defer c.Unlock()

	// Check again, maybe another thread go to this first
	// in between the c.RUnlock() and c.Lock()
	if c.statuses[key] != Absent {
		return c.data[key], false
	}

	c.statuses[key] = InProgress
	entry := &CacheEntry{Status: InProgress}
	entry.C = sync.NewCond(entry)
	c.data[key] = entry
	return entry, true
}

func (c *Cache) Get(key string) []byte {
	c.RLock()

	status := c.statuses[key]
	switch status {
	case Absent:
		c.RUnlock() // We'll take a write lock right away.

		entry, ok := c.setInProgress(key)
		if !ok {
			// Missed our chance, just wait.
			return entry.Wait()
		}

		data := getData(key)
		entry.SetComplete(data)
		c.setComplete(key)

		// Wake up everybody, not just a single goroutine
		entry.C.Broadcast()

		return data
	case InProgress:
		entry := c.data[key]
		c.RUnlock()
		return entry.Wait()
	case Complete:
		entry := c.data[key]
		c.RUnlock()
		return entry.Data
	}
	panic("not reached")
}

func main() {
	log.Println("starting")

	c := NewCache()
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			log.Printf("%s", c.Get("Batman"))
			log.Printf("%s", c.Get("Robin"))
			wg.Done()
		}()
	}
	wg.Wait()
	// These print right away, already in cache.
	log.Printf("%s", c.Get("Batman"))
	log.Printf("%s", c.Get("Robin"))

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			log.Printf("%s", c.Get("Captain America"))
			log.Printf("%s", c.Get("Thor"))
			wg.Done()
		}()
	}
	time.Sleep(time.Second)
	// These print right away, already in cache, not blocked by
	// other goroutines trying to read "captain america" and "thor"
	log.Printf("%s", c.Get("Batman"))
	log.Printf("%s", c.Get("Robin"))
	wg.Wait()
}
