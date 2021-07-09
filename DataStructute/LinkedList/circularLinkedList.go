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

func (cl *circularLinkedList) addAtPost(pos int, data int) {
	if pos == 0 {
		n := node{data, nil}
		cl.head = &n
		n.next = &n
	} else {
		startNode := cl.head
		n := node{data, nil}
		counter := 0
		for counter != pos-1 {
			startNode = startNode.next
			counter++
		}
		n.next = startNode.next
		startNode.next = &n

	}
}

func (cl *circularLinkedList) deleteAtEnd(){
	ptr := cl.head
	if(ptr.next == cl.head){
		cl.head = nil
	}else{
		for ptr.next.next != cl.head{
			ptr = ptr.next
		}
		ptr.next.next = nil
		ptr.next = cl.head
	}

}

func (cl *circularLinkedList) deleteAtStart(){
	ptr := cl.head
	if(ptr.next == cl.head){
		cl.head = nil
	}else{
		for ptr.next != cl.head{
			ptr = ptr.next
		}
		start := cl.head
		cl.head = start.next
		start.next = nil
		ptr.next = cl.head
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
	fmt.Println("----------------------")
	cl.addAtPost(2, 2)
	cl.addAtPost(3, 3)
	cl.printList()
	fmt.Println("----------------------")
	cl.deleteAtEnd()
	cl.deleteAtEnd()
	cl.printList()
	fmt.Println("----------------------")
	cl.deleteAtStart()
	cl.deleteAtStart()
	cl.printList()

}
