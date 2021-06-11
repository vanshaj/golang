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

/*
 this can't be created without pointer receiver because if pointer receiver is not used the LinkedList will be duplicated and head of duplicated LinkedList will now point to the
 new element and our actual linked list will never know about the element
*/
func (ll *LinkedList) prepend(n *node) {
	n.next = ll.head
	ll.head = n
	ll.size++
}

/*
this can be created without function pointer receiver (ll LinkedList) because element is appended at the end. Duplicated list struct head will point to the same element our actual head
point to so iterating over them will result to the same list
*/

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

// same will not support (ll LinkedList)
func (ll *LinkedList) DeleteAtStart() {
	deleteNode := ll.head
	ll.head = deleteNode.next
	ll.size--
}

// will support (ll LinkedList)
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

func (ll *LinkedList) DeleteAtPos(pos int) {
	node := ll.head
	counter := 0
	if pos > ll.size-1 {
		fmt.Println("cant delete at position greater than size")
	} else {
		if pos == 0 {
			ll.head = node.next
		} else {
			for pos-1 != counter {
				node = node.next
				counter++
			}
			node.next = node.next.next
		}
		ll.size--
	}
}

func (ll *LinkedList) emptyList() {
	ll.head = &node{}
	ll.size = 0
}

/*
 can use both pointer receiver or normal
*/
func (ll *LinkedList) printList() {
	node := ll.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func main() {
	firstElem := node{1, nil}
	head := &firstElem
	ll := &LinkedList{head, 1}
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
	ll.emptyList()
	ll.printList()
	fmt.Println("size of list is ", ll.size)
}
