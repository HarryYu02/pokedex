package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entryMap map[string]cacheEntry
	mu       sync.Mutex
}

func (c *Cache)reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entryMap {
			if time.Since(entry.createdAt) > interval {
				delete(c.entryMap, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entryMap: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache)Add(key string, val []byte) {
	c.mu.Lock()
    defer c.mu.Unlock()

	c.entryMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache)Get(key string) ([]byte, bool) {
	c.mu.Lock()
    defer c.mu.Unlock()

	entry, ok := c.entryMap[key]
	return entry.val, ok
}
