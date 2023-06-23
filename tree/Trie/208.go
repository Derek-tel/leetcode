package main

import (
	"fmt"
)

type Trie struct {
	IsEnd    bool
	Children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{IsEnd: false, Children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, letter := range word {
		if child, ok := parent.Children[letter]; ok {
			parent = child
		} else {
			newChildren := &Trie{Children: make(map[rune]*Trie)}
			parent.Children[letter] = newChildren
			parent = newChildren
		}
	}
	parent.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, letter := range word {
		if child, ok := parent.Children[letter]; ok {
			parent = child
		} else {
			return false
		}
	}

	return parent.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, letter := range prefix {
		if child, ok := parent.Children[letter]; ok {
			parent = child
		} else {
			return false
		}
	}

	return true
}

type TrieTree struct {
	IsEnd    bool
	Children map[rune]*TrieTree
}

func Construct() TrieTree {
	return TrieTree{IsEnd: false, Children: make(map[rune]*TrieTree)}
}

func (t *TrieTree) insert(word string) {
	parent := t
	for _, letter := range word {
		if child, ok := parent.Children[letter]; ok {
			parent = child
		} else {
			newChild := &TrieTree{IsEnd: false, Children: make(map[rune]*TrieTree)}
			parent.Children[letter] = newChild
			parent = newChild
		}
	}
	parent.IsEnd = true
}

func (t *TrieTree) Search(word string) bool {
	parent := t
	for _, letter := range word {
		if child, ok := parent.Children[letter]; ok {
			parent = child
		} else {
			return false
		}
	}
	return parent.IsEnd
}

func (t *TrieTree) StartWith(word string) bool {
	parent := t
	for _, letter := range word {
		if child, ok := parent.Children[letter]; ok {
			parent = child
		} else {
			return false
		}
	}
	return true
}

type three struct {
	IsEnd    bool
	Children map[rune]*three
}

func ConstructorThree() three {
	return three{IsEnd: false, Children: make(map[rune]*three)}
}

func (this *three) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
		} else {
			newChild := &three{IsEnd: false, Children: make(map[rune]*three)}
			parent.Children[ch] = newChild
			parent = newChild
		}
	}
	parent.IsEnd = true
}

func (this *three) Search(word string) bool {
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

func (this *three) StartsWith(prefix string) bool {
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

type four struct {
	IsEnd    bool
	Children map[rune]*four
}

func ConstructorFour() four {
	return four{IsEnd: false, Children: make(map[rune]*four)}
}

func (this *four) Insert(word string) {
	parent := this
	for _, v := range word {
		if child, ok := parent.Children[v]; ok {
			parent = child
		} else {
			newChild := &four{IsEnd: false, Children: make(map[rune]*four)}
			this.Children[v] = newChild
			parent = newChild
		}
	}
	parent.IsEnd = true
}

func (this *four) Search(word string) bool {
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

func (this *four) StartWith(prefix string) bool {
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
