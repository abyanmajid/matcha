package main

import (
	"fmt"
	"time"

	"github.com/abyanmajid/matcha/cache"
)

// EXPECTED OUTPUT:
//
// Found key1: value1
// Found key2: value2
// key1 not found or expired
// key3 not found or expired
// Cache metrics: {2 2 3}

func main() {
	c := cache.NewRedisCache("localhost:6379", "", 0)

	c.Set("key1", "value1", 5*time.Second)
	c.Set("key2", "value2", 5*time.Second)
	c.Set("key3", "value3", 5*time.Second)

	time.Sleep(2 * time.Second)

	val, found := c.Get("key1")
	if found {
		fmt.Println("Found key1:", val)
	} else {
		fmt.Println("key1 not found or expired")
	}

	val, found = c.Get("key2")
	if found {
		fmt.Println("Found key2:", val)
	} else {
		fmt.Println("key2 not found or expired")
	}

	time.Sleep(4 * time.Second)

	val, found = c.Get("key1")
	if found {
		fmt.Println("Found key1:", val)
	} else {
		fmt.Println("key1 not found or expired")
	}

	c.Set("key4", "value4", 5*time.Second)

	val, found = c.Get("key3")
	if found {
		fmt.Println("Found key3:", val)
	} else {
		fmt.Println("key3 not found or expired")
	}

	metrics := c.RedisMetrics()
	fmt.Println("Redis cache metrics:", metrics)

	c.Stop()
}
