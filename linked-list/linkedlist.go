package main

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data int) {
	n := &Node{
		Next: nil,
		Data: data,
	}
	if l.Head == nil {
		l.Head = n
		return
	}
	//1->2->3->4->5->
	curr := l.Head

	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
}

func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("->%d", curr.Data)
		curr = curr.Next
	}
	println()
}

func (l *LinkedList) Delete(data int) {
	prev := l.Head
	curr := l.Head

	for curr != nil {
		if curr.Data == data {

			prev.Next = curr.Next

		}
		prev = curr
		curr = curr.Next
	}
}
func main() {
	l := NewLinkedList()
	l.Insert(1)
	l.Insert(6)
	l.Insert(8)
	l.Insert(12)
	l.Print()
	l.Delete(12)
	l.Print()
}
