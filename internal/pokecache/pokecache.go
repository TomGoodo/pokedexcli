package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries: map[string]cacheEntry{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.entries[key]
	if ok == false {
		return []byte{}, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, cacheentry := range c.entries {
			if time.Since(cacheentry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
