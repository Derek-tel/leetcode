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

func initNode(x, y int) *LRUNode {
	return &LRUNode{
		Key:   x,
		Value: y,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Capacity: capacity,
		Table:    make(map[int]*LRUNode),
		Head:     initNode(0, 0),
		Tail:     initNode(0, 0),
	}
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head
	return lru
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
		tag := &LRUNode{
			Key:   key,
			Value: value,
		}
		this.AddToHead(tag)
		this.Size++
		this.Table[key] = tag
		if this.Size > this.Capacity {
			del := this.removeTail()
			delete(this.Table, del.Key)
			this.Size--
		}
	}
}

func (this *LRUCache) MoveToHead(node *LRUNode) {
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveNode(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToHead(node *LRUNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) removeTail() *LRUNode {
	node := this.Tail.Prev
	this.RemoveNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
