package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}

type Cache struct {
	cache			map[string]cacheEntry
	mux 			*sync.Mutex
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}