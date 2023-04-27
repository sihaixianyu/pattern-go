package strategy

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     uint
	maxCapacity  uint
}

func NewCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)

	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(k, v string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}

	c.capacity += 1
	c.storage[k] = v
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity -= 1
}
