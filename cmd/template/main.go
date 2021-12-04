package main

import (
	"fmt"
)

// Node is a singly linked list node, with a pointer to the next node.
type Node struct {
	value int
	next  *Node
}

var (
	// start is the first node in the singly linked list
	start *Node
)

const (
	// initSize is the length of singly-linked list to create
	initSize = 65535
)

func init() {
	start = initSinglyLinkedList(initSize)
}

func main() {

	fmt.Println("Hello")

	middle := findMiddleWithSlice(start)
	fmt.Printf("middle value from slice is %v\n", middle)

	middle = findMiddleWithPointer(start)
	fmt.Printf("middle value from pointer is %v\n", middle)

}

// findMiddleWithSlice finds the middle element by creating a slice of Nodes, then
// using len() to find the center value
func findMiddleWithSlice(start *Node) (middle int) {
	if start == nil {
		return -1
	}

	nodeSlice := make([]*Node, 0, 1)
	current := start

	for current != nil {
		nodeSlice = append(nodeSlice, current)
		current = current.next
	}

	return nodeSlice[(len(nodeSlice))/2].value
}

// findMiddleWithPointer finds the middle element by using slow/fast pointers
func findMiddleWithPointer(start *Node) (middle int) {
	if start == nil {
		return -1
	}

	slow := start
	fast := start

	counter := 0
	for fast != nil {
		fast = fast.next
		counter++

		if counter%2 == 0 {
			slow = slow.next
		}
	}

	return slow.value
}

// initSinglyLinkedList just creates a list of the passed-in size.
func initSinglyLinkedList(size int) (start *Node) {
	if size == 0 {
		return nil
	}

	start = &Node{
		value: 0,
		next:  nil,
	}

	currentNode := start

	for i := 1; i < size; i++ {
		currentNode.next = &Node{
			value: i,
			next:  nil,
		}
		currentNode = currentNode.next
	}

	return start

}
