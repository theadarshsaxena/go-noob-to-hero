package main

// https://leetcode.com/problems/delete-node-in-a-linked-list/

// Delete the given node
// Idea is to assign the value of the next node to the current node and delete the next node
func deleteNode(node *Node) {
	p := node
	p.data = p.next.data
	p.next = p.next.next
}

func main() {
	ll := &LinkedList{}
	ll.insertAtEnd(1)
	ll.insertAtEnd(2)
	ll.insertAtEnd(3)
	// need to delete 2
	deleteNode(ll.header.next)
	ll.traverse()
}