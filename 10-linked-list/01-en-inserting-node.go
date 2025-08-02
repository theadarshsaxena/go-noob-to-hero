package main

import "fmt"

type LinkedList struct {
	header *Node
	length int
}

type Node struct {
	data int
	next *Node
}

func (ll *LinkedList) insertAtHead(data int) {
	newNode := &Node{data:data}
	newNode.next = ll.header
	ll.header = newNode
	ll.length++
}

func (ll *LinkedList) insertAtEnd(data int) {
	newNode := &Node{data: data}
	if ll.header == nil {
		ll.header = newNode
	} else {
		var p *Node = ll.header
		for p.next != nil {
			p = p.next
		}
		p.next = newNode
	}
}

func (ll LinkedList) traverse() {
	var p *Node = ll.header
	for p!=nil {
		fmt.Println(p.data)
		p = p.next
	}
}

// Insert at the start of the linked list

// func main() {
// 	ll := &linkedList{} // Initializing the linked list
// 	ll.insertAtHead(10)
// 	ll.insertAtHead(9)
// 	ll.insertAtEnd(11)
// 	ll.traverse()
// }
