package main

import (
	"container/list"
	"fmt"
)

// LRUCache implements a simple Least Recently Used cache.
// It combines a map (for O(1) lookups) with a doubly-linked list (for O(1) evictions).
type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	eviction *list.List
}

// entry is the data stored in the linked list
type entry struct {
	key   int
	value string
}

// NewLRUCache creates a new LRUCache with the given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		eviction: list.New(),
	}
}

// Get retrieves a value from the cache.
func (c *LRUCache) Get(key int) (string, bool) {
	if element, ok := c.items[key]; ok {
		// Move to front (most recently used)
		c.eviction.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return "", false
}

// Put adds a value to the cache or updates an existing one.
func (c *LRUCache) Put(key int, value string) {
	if element, ok := c.items[key]; ok {
		// Update existing
		c.eviction.MoveToFront(element)
		element.Value.(*entry).value = value
		return
	}

	// Add new
	ent := &entry{key, value}
	element := c.eviction.PushFront(ent)
	c.items[key] = element

	// Evict if over capacity
	if c.eviction.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *LRUCache) removeOldest() {
	element := c.eviction.Back()
	if element != nil {
		c.eviction.Remove(element)
		ent := element.Value.(*entry)
		delete(c.items, ent.key)
	}
}

func (c *LRUCache) Print() {
	fmt.Print("Cache State (MRU -> LRU): ")
	for e := c.eviction.Front(); e != nil; e = e.Next() {
		ent := e.Value.(*entry)
		fmt.Printf("[%d:%s] ", ent.key, ent.value)
	}
	fmt.Println()
}

func main() {
	fmt.Println("--- LRU Cache Demo ---")
	cache := NewLRUCache(3)

	cache.Put(1, "Alice")
	cache.Put(2, "Bob")
	cache.Put(3, "Charlie")
	cache.Print() // [3:Charlie] [2:Bob] [1:Alice]

	fmt.Println("Getting key 1...")
	if val, ok := cache.Get(1); ok {
		fmt.Printf("Found: %s\n", val)
	}
	cache.Print() // [1:Alice] [3:Charlie] [2:Bob]

	fmt.Println("Adding key 4 (should evict Bob)...")
	cache.Put(4, "Dave")
	cache.Print() // [4:Dave] [1:Alice] [3:Charlie]

	fmt.Println("Updating key 3...")
	cache.Put(3, "Charlie-Updated")
	cache.Print() // [3:Charlie-Updated] [4:Dave] [1:Alice]
}
