package btree

const (
	M    int = 4
	MinM int = M/2 - 1
)

type BTree struct {
	parent   *BTree
	KeyNum   int
	key      [M + 1]int
	Children [M + 1]*BTree
}

func (t *BTree) Search(value int) (bool, int, *BTree) {
	node := &BTree{}
	var i int
	if t == nil {
		return false, 0, nil
	}

	for t != nil {
		for i = t.KeyNum; i > 0 && value <= t.key[i]; i-- {
			if value == t.key[i] {
				return true, i, t
			}
		}
		if t.Children[i] == nil {
			node = t
		}
		t = t.Children[i]
	}
	return false, i, node
}

func (t *BTree) split() *BTree {
	node := &BTree{}
	parent := t.parent
	if parent == nil {
		parent = &BTree{}
	}
	mid := t.KeyNum/2 + 1
	node.KeyNum = M - mid
	t.KeyNum = mid - 1
	j := 1
	k := mid + 1
	for ; k <= M; k++ {
		node.key[k] = t.key[k]
		node.Children[j-1] = t.Children[k-1]
		j = j + 1
	}

	node.Children[j-1] = t.Children[k-1]
	node.parent = parent
	t.parent = parent

	k = parent.KeyNum
	for ; t.key[mid] < t.key[k]; k-- {
		parent.key[k+1] = parent.key[k]
		parent.Children[k+1] = parent.Children[k]
	}
	parent.key[k+1] = t.key[mid]
	parent.Children[k] = t
	parent.Children[k+1] = node
	parent.KeyNum++

	if parent.KeyNum > M {
		return parent.split()
	}
	return parent
}

func (t *BTree) Insert(value int) *BTree {
	var i int
	ok, _, node := t.Search(value)
	if !ok {
		node.key[0] = value
		for i = node.KeyNum; i > 0 && value < node.key[i]; i-- {
			node.key[i+1] = node.key[i]
		}
		node.key[i+1] = value
		node.KeyNum++
		if node.KeyNum < M {
			return t
		} else {
			parent := node.split()
			for parent.parent != nil {
				parent = parent.parent
			}
			return parent
		}
	}
	return t
}

func (t *BTree) Delete(value int) *BTree {
	ok, i, node := t.Search(value)
	if ok {
		t = node.DeleteNode(value, i)
	}
	return t
}

func (t *BTree) DeleteNode(value, i int) *BTree {
	if t.Children[i] != nil { //非终端节点
		valueTemp, nodeTemp := t.FindAfterMinNode(i)
		t.key[i] = *valueTemp
		nodeTemp.DeleteNode(*valueTemp, 1)
	} else {
		for k := i; k < t.KeyNum; k++ {
			t.key[k] = t.key[k+1]
			t.Children[k] = t.Children[k+1]
		}
		t.KeyNum--
		if t.KeyNum < M/2-1 && t.parent != nil {
			ok, t := t.Restore()
			if !ok {
				t = t.MergeNode()
			}
		}
	}
	for t.parent != nil {
		t = t.parent
	}
	return t
}

func (t *BTree) Restore() (bool, *BTree) {
	parent := t.parent
	j := 0
	for ; parent.Children[j] != t; j++ {

	}
	if j > 0 {
		b := parent.Children[j-1]
		if b.KeyNum > M/2-1 {
			for k := t.KeyNum; k > 0; k-- {
				t.key[k+1] = t.key[k]
			}
			t.key[1] = parent.key[j]
			parent.key[j] = b.key[b.KeyNum]
			t.KeyNum++
			b.KeyNum--
			return true, parent
		}
	}
	if j < parent.KeyNum {
		b := parent.Children[j+1]
		if b.KeyNum > M/2-1 {
			t.key[t.KeyNum+1] = parent.key[j+1]
			parent.key[j+1] = b.key[1]
			for k := 1; k <= b.KeyNum; k++ {
				b.key[k] = b.key[k+1]
			}
			t.KeyNum++
			b.KeyNum--
			return true, parent
		}
	}
	return false, t //没有调整成功 需要合并
}

func (t *BTree) FindAfterMinNode(i int) (*int, *BTree) {
	leaf := t
	if leaf == nil {
		return nil, nil
	} else {
		leaf = leaf.Children[i]
		for leaf.Children[0] != nil {
			leaf = leaf.Children[0]
		}
	}
	return &leaf.key[1], leaf
}

func (t *BTree) MergeNode() *BTree {
	b := &BTree{}
	parent := t.parent
	j := 0
	for ; parent.Children[j] != t; j++ {
	}
	if j > 0 { //合并左兄弟
		b = parent.Children[j-1]
		b.KeyNum++
		b.key[b.KeyNum] = parent.key[j]
		for k := 1; k < t.KeyNum; k++ {
			b.KeyNum++
			b.key[b.KeyNum] = t.key[k]
		}
		parent.KeyNum--
		for k := j; k <= parent.KeyNum; k++ {
			parent.key[k] = parent.key[k+1]
			parent.Children[k] = parent.Children[k+1]
		}
	} else { //合并右兄弟
		b = parent.Children[j+1]
		t.KeyNum++
		t.key[t.KeyNum] = parent.key[j]
		for k := 1; k < t.KeyNum; k++ {
			t.KeyNum++
			t.key[b.KeyNum] = b.key[k]
		}
		parent.KeyNum--
		for k := j; k <= parent.KeyNum; k++ {
			parent.key[k] = parent.key[k+1]
			parent.Children[k] = parent.Children[k+1]
		}
	}
	return parent
}
