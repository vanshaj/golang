package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	size int
}

func (ll *LinkedList) prepend(n *node) {
	n.next = ll.head
	ll.head = n
	ll.size++
}

func (ll LinkedList) printList() {
	node := ll.head
	for ll.size != 0 {
		fmt.Println(node.data)
		node = node.next
		ll.size--
	}
}

func main() {
	firstElem := node{1, nil}
	head := &firstElem
	ll := LinkedList{head, 1}
	node1 := node{2, nil}
	ll.prepend(&node1)
	ll.prepend(&node{3, nil})
	ll.prepend(&node{4, nil})
	node1.data = 6
	ll.printList()
}
