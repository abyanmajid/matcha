package cache

import (
	"sync"
	"time"
)

// CacheItem represents an item stored in the cache.
// It contains the cached value and its expiration time in Unix nanoseconds.
type CacheItem struct {
	Value      interface{}
	Expiration int64
}

// Cache is an in-memory cache with TTL (Time-to-Live) and size limits.
// It is safe for concurrent use by multiple goroutines.
type Cache struct {
	items    map[string]CacheItem
	mu       sync.RWMutex
	maxSize  int
	metrics  Metrics
	stopChan chan struct{}
}

// Metrics tracks cache performance metrics, including hits, misses, and evictions.
type Metrics struct {
	Hits      int
	Misses    int
	Evictions int
}

// NewMemoryCache creates and initializes a new in-memory cache with a specified maximum size.
// It also starts a background goroutine to periodically evict expired items.
func NewMemoryCache(maxSize int) *Cache {
	c := &Cache{
		items:    make(map[string]CacheItem), // Initialize the cache map.
		maxSize:  maxSize,                    // Set the maximum cache size.
		metrics:  Metrics{},                  // Initialize metrics.
		stopChan: make(chan struct{}),        // Initialize the stop channel.
	}
	go c.startEvictionWorker()
	return c
}

// Set adds an item to the cache with a specified key, value, and TTL (Time-to-Live).
// If the cache is full, the oldest item is evicted to make space.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.evictExpired()

	if len(c.items) >= c.maxSize {
		c.evictOldest()
	}

	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
}

// Get retrieves an item from the cache by its key.
// It returns the value and a boolean indicating whether the item was found and not expired.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		c.metrics.Misses++
		return nil, false
	}

	if time.Now().UnixNano() > item.Expiration {
		c.metrics.Misses++
		return nil, false
	}

	c.metrics.Hits++
	return item.Value, true
}

// Delete removes an item from the cache by its key.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// evictExpired removes all expired items from the cache.
// This function is called internally to clean up the cache.
func (c *Cache) evictExpired() {
	now := time.Now().UnixNano()
	for key, item := range c.items {
		if now > item.Expiration {
			delete(c.items, key)
			c.metrics.Evictions++
		}
	}
}

// evictOldest removes the oldest item from the cache.
// This function is called internally when the cache is full.
func (c *Cache) evictOldest() {
	var oldestKey string
	var oldestExpiration int64 = time.Now().UnixNano()

	for key, item := range c.items {
		if item.Expiration < oldestExpiration {
			oldestKey = key
			oldestExpiration = item.Expiration
		}
	}

	if oldestKey != "" {
		delete(c.items, oldestKey)
		c.metrics.Evictions++
	}
}

// startEvictionWorker runs a background goroutine to periodically evict expired items.
// It checks for expired items every minute until the cache is stopped.
func (c *Cache) startEvictionWorker() {
	ticker := time.NewTicker(1 * time.Minute) // Create a ticker that fires every minute.
	defer ticker.Stop()                       // Ensure the ticker is stopped when done.

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			c.evictExpired()
			c.mu.Unlock()
		case <-c.stopChan:
			return
		}
	}
}

// Stop cleans up the cache and stops the background eviction worker.
func (c *Cache) Stop() {
	close(c.stopChan)
}

// Metrics returns the current cache performance metrics.
func (c *Cache) Metrics() Metrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.metrics
}
