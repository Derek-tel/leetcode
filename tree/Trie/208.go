package Trie

type Trie struct {
	Children map[rune]*Trie
	IsEnd    bool
}

func Constructor() Trie {
	return Trie{
		Children: make(map[rune]*Trie),
		IsEnd:    false,
	}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if node, ok := parent.Children[ch]; ok {
			parent = node
		} else {
			temp := &Trie{
				Children: make(map[rune]*Trie),
				IsEnd:    false,
			}
			parent.Children[ch] = temp
			parent = temp
		}
	}
	parent.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if node, ok := parent.Children[ch]; ok {
			parent = node
		} else {
			return false
		}
	}
	return parent.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if node, ok := parent.Children[ch]; ok {
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
