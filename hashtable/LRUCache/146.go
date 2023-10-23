package LRUCache

type LRUNode struct {
	Key        int
	Value      int
	Next, Prev *LRUNode
}

type LRUCache struct {
	Size       int
	Capacity   int
	Table      map[int]*LRUNode
	Head, Tail *LRUNode
}

func initNode(key, value int) *LRUNode {
	return &LRUNode{
		Key:   key,
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Size:     0,
		Capacity: capacity,
		Table:    map[int]*LRUNode{},
		Head:     initNode(0, 0),
		Tail:     initNode(0, 0),
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Table[key]; ok {
		this.MoveToHead(node)
		return node.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Table[key]; ok {
		node.Value = value
		this.MoveToHead(node)
	} else {
		node = initNode(key, value)
		this.AddToHead(node)
		this.Size++
		this.Table[key] = node
		if this.Size > this.Capacity {
			del := this.RemoveFromTail()
			delete(this.Table, del.Key)
			this.Size--
		}
	}
}

func (this *LRUCache) MoveToHead(node *LRUNode) {
	this.Remove(node)
	this.AddToHead(node)
}

func (this *LRUCache) Remove(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToHead(node *LRUNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) RemoveFromTail() *LRUNode {
	node := this.Tail.Prev
	this.Remove(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
