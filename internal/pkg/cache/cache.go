package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var (
	ServiceCache Cache
)

type Cache interface {
	// Get returns the value associated with the key.
	Get(key string) (interface{}, bool)
	// Set sets the value associated with the key.
	Set(key string, value interface{})
	// Set sets the value associated with the key and expiry.
	SetWithExpiry(key string, value interface{}, expiry time.Duration)
	// Delete deletes the value associated with the key.
	Delete(key string)
}

type LocalCache struct {
	cache *gocache.Cache
}

func LoadCache() Cache {
	// Create a cache which purges expired items every 1 minutes
	// replace with a distributed cache like redis for production
	ServiceCache = &LocalCache{
		cache: gocache.New(gocache.NoExpiration, 1*time.Minute),
	}
	return ServiceCache
}

func (c *LocalCache) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *LocalCache) Set(key string, value interface{}) {
	c.cache.Set(key, value, gocache.NoExpiration)
}

func (c *LocalCache) SetWithExpiry(key string, value interface{}, expiry time.Duration) {
	c.cache.Set(key, value, expiry)
}

func (c *LocalCache) Delete(key string) {
	c.cache.Delete(key)
}
