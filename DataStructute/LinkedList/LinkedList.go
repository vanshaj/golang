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

func (ll *LinkedList) apppend(n *node) {
	node := ll.head
	for {
		if node.next != nil {
			node = node.next
			continue
		}
		node.next = n
		break
	}
	ll.size++
}

func (ll *LinkedList) addAtPosition(n *node, pos int) {
	node := ll.head
	counter := 0
	if pos == 0 {
		n.next = node
		ll.head = n
	} else {
		for counter < ll.size {
			if counter == pos-1 {
				n.next = node.next
				node.next = n
				break
			}
			node = node.next
			counter++
		}
	}
	ll.size++
}

func (ll *LinkedList) DeleteAtStart() {
	deleteNode := ll.head
	ll.head = deleteNode.next
	ll.size--
}

func (ll *LinkedList) DeleteAtEnd() {
	node := ll.head
	for {
		if node.next.next == nil {
			node.next = nil
			ll.size--
			break
		} else {
			node = node.next
		}

	}
}

func (ll LinkedList) printList() {
	node := ll.head
	for node != nil {
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
	ll.apppend(&node{5, nil})
	ll.addAtPosition(&node{6, nil}, 1)
	ll.printList()
	fmt.Println("------------")
	ll.DeleteAtStart()
	ll.printList()
	fmt.Println("------------")
	ll.DeleteAtEnd()
	ll.printList()
}
