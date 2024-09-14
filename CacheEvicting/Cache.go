package CacheEvicting

type EvictionStrategy interface {
	evict(cache *Cache)string
}

type Cache struct {
	Capacity int
	Data []string
	EvictStrategy EvictionStrategy
}

func (cache *Cache) addInCache(value string){
	 if(len(cache.Data)==cache.Capacity){
	 	(cache.EvictStrategy.evict(cache));
	 }
	 cache.Data = append(cache.Data,value);
}

func NewCacheInstance(cap int,strategy EvictionStrategy) *Cache{
	return &Cache{Capacity: cap, Data: []string{},EvictStrategy: strategy};
}

func (cache *Cache) setEvictionStrategy(strat EvictionStrategy){
	cache.EvictStrategy = strat;
}


