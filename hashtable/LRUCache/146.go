package LRUCache

type LRUNode struct {
	Key, Value int
	Next, Prev *LRUNode
}
type LRUCache struct {
	Size     int
	Capacity int
	Table    map[int]*LRUNode
	Head     *LRUNode
	Tail     *LRUNode
}

func InitNode(key, value int) *LRUNode {
	return &LRUNode{
		Key:   key,
		Value: value,
	}
}
func Constructor(capacity int) LRUCache {
	node := LRUCache{
		Size:     0,
		Capacity: capacity,
		Table:    make(map[int]*LRUNode),
		Head:     InitNode(0, 0),
		Tail:     InitNode(0, 0),
	}
	node.Head.Next = node.Tail
	node.Tail.Prev = node.Head
	return node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Table[key]; ok {
		this.MoveToHead(node)
		return node.Value
	} else {
		return -1
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

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Table[key]; ok {
		node.Value = value
		this.MoveToHead(node)
	} else {
		temp := InitNode(key, value)
		this.AddToHead(temp)
		this.Table[key] = temp
		this.Size++
		if this.Size > this.Capacity {
			del := this.RemoveFromTail()
			delete(this.Table, del.Key)
			this.Size--
		}
	}
}

func (this *LRUCache) RemoveFromTail() *LRUNode {
	tail := this.Tail.Prev
	this.Remove(tail)
	return tail
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
