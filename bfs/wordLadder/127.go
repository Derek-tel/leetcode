package main

import (
	"fmt"
	"math"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	fmt.Printf("%+v\n", graph)
	fmt.Printf("%+v\n", wordId)

	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	const inF int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inF
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inF {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}

	return 0
}

func length(begin string, end string, list []string) int {
	wordId := make(map[string]int)
	graph := [][]int{}
	addWord := func(s string) int {
		id, has := wordId[s]
		if !has {
			id = len(wordId)
			wordId[s] = id
			graph = append(graph, []int{})
		}
		return id
	}

	addEdge := func(word string) int {
		idA := addWord(word)
		s := []byte(word)
		for i, i2 := range s {
			s[i] = '*'
			idB := addWord(string(s))
			graph[idA] = append(graph[idA], idB)
			graph[idB] = append(graph[idB], idA)
			s[i] = i2
		}
		return idA
	}

	for _, s := range list {
		addEdge(s)
	}
	beginId := addEdge(begin)
	endId, has := wordId[end]
	if !has {
		return 0
	}

	const inf = math.MaxInt64

	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[endId]/2 + 1
		}
		for _, i2 := range graph[top] {
			if dist[i2] == inf {
				queue = append(queue, i2)
				dist[i2] = dist[top] + 1
			}
		}
	}

	return 0
}

func test(beginWord string, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := [][]int{}
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}

	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			extId := addWord(string(wordByte))
			graph[id] = append(graph[id], extId)
			graph[extId] = append(graph[extId], id)
			wordByte[i] = v
		}
		return id
	}

	for _, s := range wordList {
		addEdge(s)
	}
	beginId := addEdge(beginWord)
	endId, ok := wordId[endWord]
	if !ok {
		return 0
	}

	const inf = math.MaxInt
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}

	for len(queue) > 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if top == endId {
			return dist[top]/2 - 1
		}
		for i := 0; i < len(graph[top]); i++ {
			if dist[i] == inf {
				dist[i] = dist[top] + 1
				queue = append(queue, dist[i])
			}
		}
	}
	return 0
}

func get(beginWord, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := [][]int{}
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			extId := addWord(string(wordByte))
			graph[id] = append(graph[id], extId)
			graph[extId] = append(graph[extId], id)
			wordByte[i] = v
		}
		return id
	}
	for _, v := range wordList {
		addEdge(v)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	const inf = math.MaxInt
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[endId]/2 - 1
		}

		for _, v := range graph[top] {
			if dist[v] == inf {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func five(beginWord, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := [][]int{}
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			extId := addWord(string(wordByte))
			graph[id] = append(graph[id], extId)
			graph[extId] = append(graph[extId], id)
			wordByte[i] = v
		}
		return id
	}
	for _, v := range wordList {
		addEdge(v)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	const inf = math.MaxInt
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[endId]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == inf {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func six(beginWord, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			extId := addWord(string(wordByte))
			graph[id] = append(graph[id], extId)
			graph[extId] = append(graph[extId], id)
			wordByte[i] = v
		}
		return id
	}
	for _, v := range wordList {
		addEdge(v)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	inf := math.MaxInt
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == inf {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}

		}
	}
	return 0
}

func seven(beginWord, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			extId := addWord(string(wordByte))
			graph[id] = append(graph[id], extId)
			graph[extId] = append(graph[extId], id)
			wordByte[i] = v
		}
		return id
	}
	for _, s := range wordList {
		addEdge(s)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	inf := math.MaxInt
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == inf {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func eight(beginWord, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			extId := addWord(string(wordByte))
			graph[id] = append(graph[id], extId)
			graph[extId] = append(graph[extId], id)
			wordByte[i] = v
		}
		return id
	}
	for _, s := range wordList {
		addEdge(s)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	inf := math.MaxInt
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == inf {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func nine(beginWord, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			tag := addWord(string(wordByte))
			graph[id] = append(graph[id], tag)
			graph[tag] = append(graph[tag], id)
			wordByte[i] = v
		}
		return id
	}

	for _, l := range wordList {
		addEdge(l)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	beginId := addEdge(beginWord)

	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}

	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, id := range graph[top] {
			if dist[id] == math.MaxInt {
				queue = append(queue, id)
				dist[id] = dist[top] + 1
			}
		}
	}
	return 0
}

func ten(beginWord, endWord string, wordList []string) int {
	var wordId = make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, ch := range wordByte {
			wordByte[i] = '*'
			tag := addWord(string(wordByte))
			graph[id] = append(graph[id], tag)
			graph[tag] = append(graph[tag], id)
			wordByte[i] = ch
		}
		return id
	}

	for _, s := range wordList {
		addEdge(s)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	beginId := addEdge(beginWord)

	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == math.MaxInt {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func eleven(beginWord, endWord string, wordList []string) int {
	var wordId = make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, ch := range wordByte {
			wordByte[i] = '*'
			maskId := addWord(string(wordByte))
			graph[id] = append(graph[id], maskId)
			graph[maskId] = append(graph[maskId], id)
			wordByte[i] = ch
		}
		return id
	}

	for _, w := range wordList {
		addEdge(w)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	beginId := addEdge(beginWord)

	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}

	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == math.MaxInt {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func twelve(beginWord, endWord string, wordList []string) int {
	var wordId = make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, ch := range wordByte {
			wordByte[i] = '*'
			maskId := addWord(string(wordByte))
			graph[id] = append(graph[id], maskId)
			graph[maskId] = append(graph[maskId], id)
			wordByte[i] = ch
		}
		return id
	}
	for _, w := range wordList {
		addEdge(w)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	beginId := addEdge(beginWord)
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == math.MaxInt {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func thirteen(beginWord, endWord string, wordList []string) int {
	var wordId = make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, ch := range wordByte {
			wordByte[i] = '*'
			maskId := addWord(string(wordByte))
			graph[id] = append(graph[id], maskId)
			graph[maskId] = append(graph[maskId], id)
			wordByte[i] = ch
		}
		return id
	}
	for _, w := range wordList {
		addEdge(w)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	beginId := addEdge(beginWord)
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == math.MaxInt {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func fourteen(beginWord, endWord string, wordList []string) int {
	var wordId = make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, ch := range wordByte {
			wordByte[i] = '*'
			maskId := addWord(string(wordByte))
			graph[id] = append(graph[id], maskId)
			graph[maskId] = append(graph[maskId], id)
			wordByte[i] = ch
		}
		return id
	}
	for _, word := range wordList {
		addEdge(word)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	beginId := addEdge(beginWord)
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == math.MaxInt {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func fifteen(beginWord, endWord string, wordList []string) int {
	var wordId = make(map[string]int)
	var graph [][]int
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		wordByte := []byte(word)
		for i, v := range wordByte {
			wordByte[i] = '*'
			maskId := addWord(string(wordByte))
			graph[id] = append(graph[id], maskId)
			graph[maskId] = append(graph[maskId], id)
			wordByte[i] = v
		}
		return id
	}
	for _, word := range wordList {
		addEdge(word)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	beginId := addEdge(beginWord)
	dist := make([]int, len(wordId))
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == endId {
			return dist[top]/2 + 1
		}
		for _, v := range graph[top] {
			if dist[v] == math.MaxInt {
				queue = append(queue, v)
				dist[v] = dist[top] + 1
			}
		}
	}
	return 0
}

func main() {

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	res := ladderLength(beginWord, endWord, wordList)

	fmt.Println(res)
}
