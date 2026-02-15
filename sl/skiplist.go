package sl

import (
	"math"
	"math/rand/v2"
)

var (
	maxHeight = 4
	pValue    = 1 / math.E // Saw that somewhere, might be a good choice
)

type Node struct {
	Key    int
	Value  int
	height int
	levels [5]*Node
}

type Skiplist struct {
	head  *Node
	level int // current highest level
}

func NewSkiplist() *Skiplist {
	// head node with max height
	head := &Node{Key: -1, height: maxHeight}
	return &Skiplist{
		head:  head,
		level: 0,
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
	update := [5]*Node{}
	curr := l.head

	for i := maxHeight; i >= 0; i-- {
		for curr.levels[i] != nil && curr.levels[i].Key < key {
			curr = curr.levels[i]
		}
		update[i] = curr
	}
	curr = curr.levels[0]

	// TODO : handle update value case
	if curr == nil || curr.Key != key {
		rHeight := l.rHeight()

		// update current level
		if rHeight > l.level {
			for i := l.level + 1; i <= rHeight; i++ {
				update[i] = l.head
			}
			l.level = rHeight
		}

		// insert new node and update levels
		newNode := Node{Key: key, Value: value, height: rHeight}
		for i := range rHeight + 1 {
			newNode.levels[i] = update[i].levels[i]
			update[i].levels[i] = &newNode
		}
	}
}

// func (l *Skiplist) Find(key int) *Node {

// }

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
