package main

type Node struct {
	key        int
	value      int
	prev, next *Node
}

type LRU struct {
	size       int
	cap        int
	cache      map[int]*Node
	head, tail *Node
}

func initNode(k, v int) *Node {
	return &Node{key: k, value: v}
}

func construct(c int) LRU {
	lru := LRU{
		cap:   c,
		cache: make(map[int]*Node),
		head:  initNode(0, 0),
		tail:  initNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

func (l *LRU) get(k int) int {
	node, ok := l.cache[k]
	if ok {
		//move to head
		l.moveToHead(node)
		return node.value
	}
	return -1
}

func (l *LRU) put(k, v int) {
	node, ok := l.cache[k]
	if ok { //can touch  remove and update
		l.moveToHead(node)
		node.value = v
	} else {
		node = initNode(k, v)
		l.addToHead(node)
		l.size++
		if l.size > l.cap {

		}
	}
}

func (l *LRU) moveToHead(node *Node) {
	l.remove(node)
	l.addToHead(node)
}

func (l *LRU) addToHead(node *Node) {
	node.next = l.head.next
	l.head.next = node
}

func (l *LRU) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRU) removeTail() *Node {
	node := l.tail.prev
	l.remove(node)
	return node
}
