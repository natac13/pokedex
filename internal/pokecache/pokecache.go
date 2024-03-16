package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cache[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now().UTC()
		c.reap(now, interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, k)
		}
	}
}
