package strategy

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type FIFO struct{}

func (f *FIFO) evict(c *Cache)  {
	fmt.Println("Evicting by FIFO algorithm...")
}

type LRU struct{}

func (l *LRU) evict(c *Cache)  {
	fmt.Println("Evicting by LRU algorithm...")
}

type LFU struct{}

func (l *LFU) evict(c *Cache)  {
	fmt.Println("Evicting by LFU algorithm...")
}
