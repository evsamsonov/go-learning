package main

import (
	"fmt"
)

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

type LinkedList struct {
	len  int
	head *LinkedListNode
	tail *LinkedListNode
}

func (l *LinkedList) Add(value int) {
	if l.head == nil {
		n := &LinkedListNode{
			value: value,
		}
		l.head = n
		l.tail = n
		l.len = 1
		return
	}

	n := &LinkedListNode{
		value: value,
	}
	l.tail.next = n
	l.tail = n
	l.len++
	return
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Traverse() {
	if l.head == nil {
		return
	}
	node := l.head
	for {
		fmt.Printf("%d -> ", node.value)
		if node.next == nil {
			fmt.Println()
			break
		}
		node = node.next
	}
}

func (l *LinkedList) Lookup(value int) bool {
	if l.head == nil {
		return false
	}
	node := l.head
	for {
		if node.value == value {
			return true
		}
		if node.next == nil {
			break
		}
		node = node.next
	}
	return false
}

// Remove removes first equal element
func (l *LinkedList) Remove(value int) bool {
	if l.head == nil {
		return false
	}
	var previous *LinkedListNode
	current := l.head
	for {
		if current.value == value {
			if previous == nil {
				l.head = l.head.next
			} else {
				previous.next = current.next
			}
			if current == l.tail {
				l.tail = previous
			}
			l.len--
			return true
		}
		if current.next == nil {
			break
		}
		previous = current
		current = current.next
	}
	return false
}

func main() {
	linkedList := &LinkedList{}
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(4)
	linkedList.Add(3)
	linkedList.Add(10)
	fmt.Println(linkedList.Len())
	fmt.Println(linkedList.Lookup(4))
	fmt.Println(linkedList.Lookup(11))
	linkedList.Traverse()

	fmt.Println()
	linkedList.Remove(2)
	linkedList.Remove(2)
	linkedList.Remove(4)
	linkedList.Remove(3)
	linkedList.Remove(10)
	linkedList.Add(11)
	linkedList.Add(12)
	linkedList.Add(13)
	fmt.Println(linkedList.Len())
	fmt.Println(linkedList.Lookup(4))
	linkedList.Traverse()
}
