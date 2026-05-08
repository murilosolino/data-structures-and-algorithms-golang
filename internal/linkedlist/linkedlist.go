package linkedlist

import "fmt"

type Node struct {
	Value int
	next  *Node
}

func newNode(value int) *Node {
	return &Node{
		Value: value,
		next:  nil,
	}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

func (l *LinkedList) AddToEnd(value int) {
	newNode := newNode(value)
	if l.tail != nil {
		l.tail.next = newNode
		l.tail = newNode
		return
	}
	l.head = newNode
	l.tail = newNode
}

func (l *LinkedList) AddToFront(value int) {
	newNode := newNode(value)

	if l.head != nil {
		newNode.next = l.head
		l.head = newNode
		return
	}

	l.head = newNode
	l.tail = newNode
}

func (l *LinkedList) RemoveToFront() *Node {
	if l.head != nil {
		removed := l.head
		l.head = l.head.next
		return removed
	}
	return &Node{}
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) GetTail() *Node {
	return l.tail
}

func (l *LinkedList) PrintLinkedList() {
	var pointer Node = *l.head
	for {
		fmt.Printf("%d -> ", pointer.Value)
		if pointer.next == nil {
			fmt.Printf("NULL \n")
			break
		}
		pointer = *pointer.next
	}
}

func (l *LinkedList) SearchValueInLinkedList(target int) int {
	var pointer Node = *l.head
	index := 0
	for {
		if pointer.Value == target {
			return index
		}

		if pointer.next == nil {
			return -1
		}

		pointer = *pointer.next
		index++
	}
}
