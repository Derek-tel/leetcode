package LRUCache

type LRUNode struct {
	Key, Value int
	Next, Prev *LRUNode
}
type LRUCache struct {
	Size       int
	Capacity   int
	Table      map[int]*LRUNode
	Head, Tail *LRUNode
}

func initLRUNode(key, value int) *LRUNode {
	return &LRUNode{
		Key:   key,
		Value: value,
	}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Size:     0,
		Capacity: capacity,
		Table:    make(map[int]*LRUNode),
		Head:     initLRUNode(0, 0),
		Tail:     initLRUNode(0, 0),
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
		tmp := initLRUNode(key, value)
		this.AddToHead(tmp)
		this.Table[key] = tmp
		this.Size++
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
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
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
