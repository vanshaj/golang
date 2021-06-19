package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type doubleLinkedList struct {
	head *node
}

func (dl *doubleLinkedList) insertAtBeg(data int) {
	if dl.head == nil {
		dl.head = &node{data, nil, nil}
	} else {
		n1 := node{data, nil, nil}
		n1.prev = nil
		n1.next = dl.head
		dl.head = &n1
	}

}
func (dl *doubleLinkedList) insertAtEnd(data int) {
	if dl.head == nil {
		dl.head = &node{data, nil, nil}
	} else {
		nod := dl.head
		for nod.next != nil {
			nod = nod.next
		}
		n1 := node{data, nod, nil}
		nod.next = &n1
	}
}

func (dl *doubleLinkedList) printData() {
	nod := dl.head
	for nod != nil {
		fmt.Println(nod.data)
		nod = nod.next
	}
}

func main() {
	dl := &doubleLinkedList{nil}
	dl.insertAtBeg(1)
	dl.insertAtBeg(2)
	dl.insertAtBeg(3)
	dl.printData()
	fmt.Println("----------------------")
	dl.insertAtEnd(0)
	dl.insertAtEnd(4)
	dl.insertAtBeg(5)
	dl.printData()

}
