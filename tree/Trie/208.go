package main

import (
	"fmt"
)

type Trie struct {
	Child map[rune]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		Child: make(map[rune]*Trie),
		isEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if node, ok := parent.Child[ch]; ok {
			parent = node
		} else {
			newChild := &Trie{
				Child: make(map[rune]*Trie),
				isEnd: false,
			}
			parent.Child[ch] = newChild
			parent = newChild
		}
	}
	parent.isEnd = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if node, ok := parent.Child[ch]; ok {
			parent = node
		} else {
			return false
		}
	}
	return parent.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if node, ok := parent.Child[ch]; ok {
			parent = node
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

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
