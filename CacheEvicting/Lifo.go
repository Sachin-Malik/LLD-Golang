package CacheEvicting

import "fmt"

type Lifo struct {
}

func (lifo *Lifo) evict(cache *Cache) string{
	len := len(cache.Data);
	lastValue := cache.Data[len-1];
	cache.Data = cache.Data[0:len-1];
	fmt.Printf("Removed %s\n",lastValue)
	return lastValue;
}
