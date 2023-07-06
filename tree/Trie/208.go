package main

import (
	"fmt"
)

type Trie struct {
	IsEnd    bool
	Children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		IsEnd:    false,
		Children: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{Children: make(map[rune]*Trie)}
			parent.Children[ch] = newChild
			parent = newChild
		}
	}
	parent.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
		} else {
			return false
		}
	}
	return parent.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.Children[ch]; ok {
			parent = child
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
