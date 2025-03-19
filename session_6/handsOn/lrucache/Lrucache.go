package lrucache

import "fmt"

type LRUCache struct {
	mp    map[int]int
	order []int
	cap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		mp:    make(map[int]int),
		order: make([]int, 0, capacity),
		cap:   capacity,
	}

}
func (this *LRUCache) Remove_key_order(key int) {
	for ind, val := range this.order {
		if val == key {
			this.order = append(this.order[:ind], this.order[ind+1:]...)
			return
		}
	}
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("Get operation:", key)
	val, found := this.mp[key]
	if found {
		this.Remove_key_order(key)
		this.order = append(this.order, key)
		fmt.Println(key, "-", val)
		return val

	}
	fmt.Println("Not found")
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println("Put operation:{Key:", key, ",val:", value, "}")
	_, found := this.mp[key]
	if found {
		this.Remove_key_order(key)
		this.order = append(this.order, key)
		this.mp[key] = value

	} else {

		if len(this.order) == this.cap {
			lru := this.order[0]
			delete(this.mp, lru)
			this.order = this.order[1:]

		}
		this.mp[key] = value
		this.order = append(this.order, key)

	}
	fmt.Println("cache memory:", this.mp)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
