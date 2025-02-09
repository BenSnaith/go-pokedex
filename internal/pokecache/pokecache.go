package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheTable map[string]cacheEntry
	mu         *sync.Mutex
}

type cacheEntry struct {
	// time cache entry was created
	createdAt time.Time
	// raw cached data
	val []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheTable: make(map[string]cacheEntry),
		mu:         &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheTable[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cacheTable[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cacheTable {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheTable, k)
		}
	}
}
