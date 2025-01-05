package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	cacheEntryMap map[string]cacheEntry
	interval      time.Duration
	sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheEntryMap: map[string]cacheEntry{},
		interval:      interval,
	}
	go cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()

	t0 := time.Now()
	cE := cacheEntry{
		createdAt: t0,
		value:     val,
	}

	c.cacheEntryMap[key] = cE

}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()

	val, ok := c.cacheEntryMap[key]
	if ok != true {
		return nil, false
	}
	return val.value, true

}

func (c Cache) reapLoop() {
	//TODO:
	interval := c.interval

	tk := time.NewTicker(interval)
	defer tk.Stop()

	cacheMap := c.cacheEntryMap

	for {
		select {
		case <-tk.C:

			for key, cEntry := range cacheMap {
				if time.Since(cEntry.createdAt) > interval {
					delete(cacheMap, key)
				}
			}
		}
	}

}
