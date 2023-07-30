package main

import (
	"fmt"
)

type Trie struct {
	Child map[rune]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{
		IsEnd: true,
		Child: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, v := range word {
		if node, ok := parent.Child[v]; ok {
			parent = node
		} else {
			newChild := &Trie{IsEnd: true, Child: make(map[rune]*Trie)}
			parent.Child[v] = newChild
			parent = newChild
		}
	}
	parent.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, v := range word {
		if node, ok := parent.Child[v]; ok {
			parent = node
		} else {
			return false
		}
	}
	return parent.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, v := range prefix {
		if node, ok := parent.Child[v]; ok {
			parent = node
		} else {
			return false
		}
	}
	return true
}

func main() {
	obj := Constructor()
	obj.Insert("apple")
	param_2 := obj.Search("word")
	param_3 := obj.StartsWith("app")
	fmt.Println(param_2, param_3)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
