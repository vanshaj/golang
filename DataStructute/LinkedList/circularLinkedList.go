package main

import "fmt"

type node struct {
	data int
	next *node
}

type circularLinkedList struct {
	head *node
}

func (cl *circularLinkedList) addAtBeg(data int) {
	if cl.head == nil {
		n := node{data, nil}
		cl.head = &n
		n.next = &n
	} else {
		n := node{data, cl.head}
		ptr := cl.head
		for ptr.next != cl.head {
			ptr = ptr.next
		}
		cl.head = &n
		ptr.next = &n
	}
}

func (cl *circularLinkedList) addAtEnd(data int) {
	if cl.head == nil {
		n := node{data, nil}
		cl.head = &n
		n.next = &n
	} else {
		n := node{data, cl.head}
		ptr := cl.head
		for ptr.next != cl.head {
			ptr = ptr.next
		}
		ptr.next = &n
	}
}

func (cl *circularLinkedList) printList() {
	first := cl.head
	if first != nil {
		fmt.Println(first.data)
		ptr := first.next
		for ptr != first {
			fmt.Println(ptr.data)
			ptr = ptr.next
		}
	}
}

func main() {
	cl := &circularLinkedList{nil}
	cl.addAtBeg(3)
	cl.addAtBeg(4)
	cl.addAtBeg(5)
	cl.printList()
	fmt.Println("----------------------")
	cl.addAtEnd(2)
	cl.addAtEnd(1)
	cl.printList()
}
