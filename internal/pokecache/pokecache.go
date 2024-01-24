package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache: make(map[string]cacheEntry),
		mutex: &sync.RWMutex{},
	}
	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, found := c.cache[key]
	return entry.val, found
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for now := range ticker.C {
		cache.reap(now, interval)
	}
}

func (cache *Cache) reap(now time.Time, interval time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for key, value := range cache.cache {
		if now.Sub(value.createdAt) > interval {
			delete(cache.cache, key)
		}
	}
}
