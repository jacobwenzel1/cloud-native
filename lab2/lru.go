package lru   

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int               //size of cache
	remaining int               //remaining capacity
	cache     map[string]string //to store cache data
	queue     []string          //to store order of cache access
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, 0)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	// Your code here....
	if ele, ok := lru.cache[key.(string)]; ok { //if cache has data for given key, return the data
		lru.MoveToFront(ele)
		return ele
	}

	return nil //if cache doesn't have data, return null

	// if lru.cache != null{
	// 	lru.queue.MoveToFront(cache)

	// }

	// return null;

}

func (lru *lruCache) Put(key, val interface{}) error {
	// Your code here....
	if ele, ok := lru.cache[key.(string)]; ok { //if cache isn't full, add data to cache
		lru.MoveToFront(ele) //and store reference of cache to the front
		lru.cache[key] = val
	} else {
		if len(lru.queue) >= lru.size { //compares size and queue to check if cache is full
			(func(lru *lruCache) qDel)(ele.(string)) //since full, head of queue(last recently used element) is removed
		}
	}
	// if len(cache) > size {
	// 	 qDel(ele string)
	// }

	// else{

	// }

}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}
