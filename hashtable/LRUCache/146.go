package LRUCache

type LRUNode struct {
	Key, Value int
	Next, Prev *LRUNode
}

type LRUCache struct {
	Capacity   int
	Size       int
	Table      map[int]*LRUNode
	Head, Tail *LRUNode
}

func initNode(key, value int) *LRUNode {
	return &LRUNode{Key: key, Value: value}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Size:     0,
		Table:    make(map[int]*LRUNode),
		Head:     initNode(0, 0),
		Tail:     initNode(0, 0),
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head

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

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Table[key]; ok {
		node.Value = value
		this.moveToHead(node)
	} else {
		top := initNode(key, value)
		this.addToHead(top)
		this.Size++
		this.Table[key] = top
		if this.Size > this.Capacity {
			del := this.removeFromTail()
			delete(this.Table, del.Key)
			this.Size--
		}
	}
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.remove(node)
	this.addToHead(node)
}

func (this *LRUCache) remove(node *LRUNode) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) removeFromTail() *LRUNode {
	node := this.Tail.Prev
	this.remove(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
