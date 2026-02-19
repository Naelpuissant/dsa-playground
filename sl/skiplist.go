package sl

import (
	"math/rand/v2"
)

var (
	// 	maxHeight = 48
	// 	pValue    = 1 / math.E // Saw that somewhere, might be a good choice
	maxHeight = 32
	pValue    = 0.5
)

type Node struct {
	Key    int
	Value  int
	height int
	levels []*Node
}

func NewNode(Key, Value, height int) *Node {
	return &Node{
		Key:    Key,
		Value:  Value,
		height: height,
		levels: make([]*Node, height+1),
	}
}

type Skiplist struct {
	head   *Node
	level  int // current highest level
	length int

	update []*Node // Store and reuse update node links
}

func NewSkiplist() *Skiplist {
	// head node with max height
	head := NewNode(-1, 0, maxHeight+1)
	update := make([]*Node, maxHeight+1)
	return &Skiplist{
		head:   head,
		level:  0,
		update: update,
	}
}

func (l *Skiplist) rHeight() int {
	h := 0
	for h < maxHeight && (rand.Float64() < pValue) {
		h++
	}
	return h
}

func (l *Skiplist) Insert(key, value int) {
	// reach insert point
	curr := l.head

	for i := maxHeight; i >= 0; i-- {
		for curr.levels[i] != nil && curr.levels[i].Key < key {
			curr = curr.levels[i]
		}
		l.update[i] = curr
	}
	curr = curr.levels[0]

	// if key exists, update value
	if curr != nil && curr.Key == key {
		curr.Value = value
		return
	}

	if curr == nil || curr.Key != key {
		rHeight := l.rHeight()

		// update current level
		if rHeight > l.level {
			for i := l.level + 1; i <= rHeight; i++ {
				l.update[i] = l.head
			}
			l.level = rHeight
		}

		// insert new node and update levels
		newNode := NewNode(key, value, rHeight)
		for i := range rHeight + 1 {
			newNode.levels[i] = l.update[i].levels[i]
			l.update[i].levels[i] = newNode
		}
		l.length++
	}
}

func (l *Skiplist) Search(key int) *Node {
	curr := l.head
	for i := l.level; i >= 0; i-- {
		for curr.levels[i] != nil && curr.levels[i].Key < key {
			curr = curr.levels[i]
		}
	}
	curr = curr.levels[0]

	if curr != nil && curr.Key == key {
		return curr
	}

	return nil
}

func (l *Skiplist) Keys() []int {
	res := []int{}
	// Start after head
	curr := l.head.levels[0]
	for curr != nil {
		res = append(res, curr.Key)
		curr = curr.levels[0]
	}
	return res
}

func (l *Skiplist) Length() int {
	return l.length
}
