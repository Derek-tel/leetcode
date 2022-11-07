package LRUCache

type LRUNode struct {
	Key, Value int
	prev, next *LRUNode
}
type LRUCache struct {
	Size       int
	Capacity   int
	Table      map[int]*LRUNode
	Head, Tail *LRUNode
}

func initLRUNode(key, value int) *LRUNode {
	return &LRUNode{Key: key, Value: value}
}
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Size:     0,
		Capacity: capacity,
		Table:    make(map[int]*LRUNode),
		Head:     initLRUNode(0, 0),
		Tail:     initLRUNode(0, 0),
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

func (this *LRUCache) removeTail() *LRUNode {
	node := this.Tail.prev
	this.remove(node)
	return node
}

func (this *LRUCache) remove(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev

}
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Table[key]; ok {
		node.Value = value
		this.moveToHead(node)
	} else {
		n := initLRUNode(key, value)
		this.Table[key] = n
		this.addToHead(n)
		this.Size++
		if this.Size > this.Capacity {
			del := this.removeTail()
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
