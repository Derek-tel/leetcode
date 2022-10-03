package main

import "fmt"

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*LruListNode
	head, tail *LruListNode
}

type LruListNode struct {
	key, value int
	prev, next *LruListNode
}

func initListNode(key, value int) *LruListNode {
	return &LruListNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    map[int]*LruListNode{},
		head:     initListNode(0, 0),
		tail:     initListNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node := initListNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}

}

func (this *LRUCache) addToHead(node *LruListNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LruListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *LruListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LruListNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(0)
	obj.Put(5, 3)
	param_1 := obj.Get(5)
	fmt.Println(param_1)
}

type Node struct {
	key        int
	value      int
	prev, next *Node
}
type LruCache struct {
	size       int
	capacity   int
	table      map[int]*Node
	head, tail *Node
}

func initNode(key, value int) *Node {
	return &Node{key: key, value: value}
}

func Con(capacity int) LruCache {
	cache := LruCache{
		0, capacity, make(map[int]*Node), initNode(0, 0), initNode(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (lru *LruCache) Get(key int) int {
	if node, ok := lru.table[key]; ok {
		lru.MoveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (lru *LruCache) MoveToHead(node *Node) {
	lru.MoveNode(node)
	lru.AddToHead(node)
}

func (lru *LruCache) MoveNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LruCache) AddToHead(node *Node) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LruCache) RemoveTail() *Node {
	node := lru.tail.prev
	lru.MoveNode(node)
	return node
}

func (lru *LruCache) Put(key, value int) {
	if node, ok := lru.table[key]; ok {
		node.value = value
		lru.MoveToHead(node)
	} else {
		node := initNode(key, value)
		lru.table[key] = node
		lru.AddToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			del := lru.RemoveTail()
			delete(lru.table, del.key)
			lru.size--
		}
	}
}

type LruNode struct {
	key        int
	value      int
	prev, next *LruNode
}

type Lru struct {
	size       int
	capacity   int
	table      map[int]*LruNode
	head, tail *LruNode
}

func initLruNode(key, value int) *LruNode {
	return &LruNode{key: key, value: value}
}
func Construct(capacity int) Lru {
	cache := Lru{
		capacity: capacity,
		table:    map[int]*LruNode{},
		head:     initLruNode(0, 0),
		tail:     initLruNode(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (l *Lru) Get(key int) int {
	if node, ok := l.table[key]; ok {
		l.MoveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (l *Lru) MoveToHead(node *LruNode) {
	l.MoveNode(node)
	l.AddToHead(node)
}

func (l *Lru) MoveNode(node *LruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *Lru) AddToHead(node *LruNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}
func (l *Lru) removeTail() *LruNode {
	node := l.tail.prev
	l.MoveNode(node)
	return node
}

func (l *Lru) Put(key, value int) {
	if node, ok := l.table[key]; ok {
		node.value = value
		l.MoveToHead(node)
	} else {
		n := initLruNode(key, value)
		l.table[key] = n
		l.AddToHead(n)
		l.size++
		for l.size > l.capacity {
			del := l.removeTail()
			delete(l.table, del.key)
			l.size--
		}
	}
}
