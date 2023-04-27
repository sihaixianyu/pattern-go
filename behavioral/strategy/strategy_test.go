package strategy

import (
	"testing"
)

func TestStragety(t *testing.T) {
	lfu := &LFU{}
	cache := NewCache(lfu)

	cache.Add("a", "1")
    cache.Add("b", "2")
	cache.Add("c", "3")

	lru := &LRU{}
	cache.SetEvictionAlgo(lru)
	cache.Add("d", "4")

	fifo := &FIFO{}
	cache.SetEvictionAlgo(fifo)
	cache.Add("e", "5")
}
