package CacheEvicting

import "fmt"


type Fifo struct {
}

func (fifo *Fifo) evict(cache *Cache) string {
	// just remove the first element
	value:= cache.Data[0];
	cache.Data = cache.Data[1:];
	fmt.Printf("Removed %s\n",value )
	return value;
}