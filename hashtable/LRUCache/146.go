package LRUCache

type LRUNode struct {
	Key, Value int
	prev, next *LRUNode
}

func initNode(key, value int) *LRUNode {
	return &LRUNode{Key: key, Value: value}
}

type LRUCache struct {
	Size       int
	Capacity   int
	Table      map[int]*LRUNode
	Head, Tail *LRUNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Size:     0,
		Capacity: capacity,
		Table:    make(map[int]*LRUNode),
		Head:     initNode(0, 0),
		Tail:     initNode(0, 0),
	}
	cache.Head.next = cache.Tail
	cache.Tail.prev = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Table[key]; ok {
		this.moveToHead(node)
		return node.Value
	} else {
		return -1
	}
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.remove(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.prev = this.Head
	node.next = this.Head.next
	this.Head.next.prev = node
	this.Head.next = node
}

func (this *LRUCache) removeFromTail() *LRUNode {
	node := this.Tail.prev
	this.remove(node)
	return node
}

func (this *LRUCache) remove(node *LRUNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.Table[key]; ok {
		node.Value = value
		this.moveToHead(node)
	} else {
		n := initNode(key, value)
		this.addToHead(n)
		this.Table[key] = n
		this.Size++
		if this.Size > this.Capacity {
			del := this.removeFromTail()
			delete(this.Table, del.Key)
			this.Size--
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
