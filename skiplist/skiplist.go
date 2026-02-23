package skiplist

import (
	"bytes"
	"math/rand/v2"
	"sync"
	"sync/atomic"
)

var (
	maxHeight = 32
	pValue    = 0.5
)

type Node struct {
	Key    []byte
	Value  []byte
	height int
	levels []*Node
}

func NewNode(Key, Value []byte, height int) *Node {
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
	length atomic.Int64
	rand   *rand.Rand
	mu     *sync.RWMutex
	update []*Node // Store and reuse update node links
}

func New() *Skiplist {
	// head node with max height
	head := NewNode(nil, nil, maxHeight+1)
	update := make([]*Node, maxHeight+1)
	randSrc := rand.NewChaCha8([32]byte{byte(42)})
	return &Skiplist{
		head:   head,
		level:  0,
		rand:   rand.New(randSrc),
		mu:     &sync.RWMutex{},
		update: update,
	}
}

func (l *Skiplist) rHeight() int {
	h := 0
	for h < maxHeight && (l.rand.Float64() < pValue) {
		h++
	}
	return h
}

// Insert adds a key-value pair to the skiplist.
// If the key already exists, it updates the value (O(log(n)))
func (l *Skiplist) Insert(key, value []byte) {
	l.mu.Lock()
	defer l.mu.Unlock()

	curr := l.head

	for i := maxHeight; i >= 0; i-- {
		for curr.levels[i] != nil && bytes.Compare(curr.levels[i].Key, key) < 0 {
			curr = curr.levels[i]
		}
		l.update[i] = curr
	}
	curr = curr.levels[0]

	if curr != nil && bytes.Equal(curr.Key, key) {
		curr.Value = value
		return
	}

	if curr == nil || !bytes.Equal(curr.Key, key) {
		rHeight := l.rHeight()

		if rHeight > l.level {
			for i := l.level + 1; i <= rHeight; i++ {
				l.update[i] = l.head
			}
			l.level = rHeight
		}

		newNode := NewNode(key, value, rHeight)
		for i := range rHeight + 1 {
			newNode.levels[i] = l.update[i].levels[i]
			l.update[i].levels[i] = newNode
		}
		l.length.Add(1)
	}
}

// Search returns the node with the given key, or nil if not found (O(log(n)))
func (l *Skiplist) Search(key []byte) *Node {
	l.mu.RLock()
	defer l.mu.RUnlock()

	curr := l.head
	for i := l.level; i >= 0; i-- {
		for curr.levels[i] != nil && bytes.Compare(curr.levels[i].Key, key) < 0 {
			curr = curr.levels[i]
		}
	}
	curr = curr.levels[0]

	if curr != nil && bytes.Equal(curr.Key, key) {
		return curr
	}

	return nil
}

// returns all keys in the skiplist in sorted order (O(n))
func (l *Skiplist) Keys() [][]byte {
	l.mu.RLock()
	defer l.mu.RUnlock()

	res := [][]byte{}
	curr := l.head.levels[0]
	for curr != nil {
		res = append(res, curr.Key)
		curr = curr.levels[0]
	}
	return res
}

func (l *Skiplist) Length() int64 {
	return l.length.Load()
}
