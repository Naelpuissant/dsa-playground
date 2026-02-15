package oll

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type OrderedLinkedList struct {
	head *Node
}

func NewOrderedLinkedList() *OrderedLinkedList {
	return &OrderedLinkedList{
		head: nil,
	}
}

func (l *OrderedLinkedList) Insert(key, value int) {
	newNode := &Node{Key: key, Value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	prev := (*Node)(nil)
	curr := l.head
	for curr != nil {
		if newNode.Key < curr.Key {
			newNode.Next = curr
			if prev == nil {
				l.head = newNode
			} else {
				prev.Next = newNode
			}
			return
		}
		prev = curr
		curr = curr.Next
	}
	prev.Next = newNode
}

func (l *OrderedLinkedList) Find(key int) *Node {
	curr := l.head
	for curr != nil {
		if curr.Key == key {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

func (l *OrderedLinkedList) Delete(key int) {
	prev := (*Node)(nil)
	curr := l.head
	for curr != nil {
		if curr.Key == key {
			if prev != nil {
				prev.Next = curr.Next
			} else {
				l.head = curr.Next
			}
			return
		}
		prev = curr
		curr = curr.Next
	}
}

func (l *OrderedLinkedList) Keys() []int {
	res := []int{}
	curr := l.head
	for curr != nil {
		res = append(res, curr.Key)
		curr = curr.Next
	}
	return res
}
