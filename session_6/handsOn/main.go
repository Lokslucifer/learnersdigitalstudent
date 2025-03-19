package main

import (
	lrucache "handsOn/Lrucache"
)

func main() {
	cache := lrucache.Constructor(2)
	cache.Put(1, 1)

	cache.Put(2, 2)
	cache.Get(1)    // Output: 1
	cache.Put(3, 3) // Removes key 2 (least recently used)
	cache.Get(2)    // Output: -1 (not found)
	cache.Put(4, 4) // Removes key 1 (least recently used)
	cache.Get(1)    // Output: -1 (not found)
	cache.Get(3)    // Output: 3
	cache.Get(4)    // Output: 4

}
