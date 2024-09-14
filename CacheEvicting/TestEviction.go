package CacheEvicting

import "fmt"

func TestEviction(){
	lifoStrategy := &Lifo{}
	cache := NewCacheInstance(3,lifoStrategy);

	cache.addInCache("sachin");
	cache.addInCache("Malik");
	cache.addInCache("Nitin");
	fmt.Printf("All the data is %v\n",cache.Data)
	// the next insert will cause an eviction
	// will be removing "Nitin"
	cache.addInCache("sachin_again");
	// changing the strategy
	fifoStrategy := &Fifo{}
	cache.setEvictionStrategy(fifoStrategy);
	// adding now will remove the first element
	cache.addInCache("Golang Rules");
	fmt.Printf("All the data is %v\n",cache.Data)

}