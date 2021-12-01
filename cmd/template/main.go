package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func main() {

	fmt.Println("Hello")

	start := initSinglyLinkedList(65535)

	middle := FindMiddleWithSlice(start)
	fmt.Printf("middle value from slice is %v\n", middle)

	middle = FindMiddleWithPointer(start)
	fmt.Printf("middle value from pointer is %v\n", middle)

}

// FindMiddleWithSlice finds the middle element by creating a slice of Nodes, then
// using len() to find the center value
func FindMiddleWithSlice(start *Node) (middle int) {
	if start == nil {
		return -1
	}

	nodeSlice := make([]*Node, 0, 65535)
	current := start

	for current != nil {
		nodeSlice = append(nodeSlice, current)
		current = current.next
	}

	return nodeSlice[(len(nodeSlice)+1)/2].value
}

// FindMiddleWithPointer finds the middle element by using slow/fast pointers
func FindMiddleWithPointer(start *Node) (middle int) {
	if start == nil {
		return -1
	}

	slow := start
	fast := start

	counter := 0
	for fast != nil {
		fast = fast.next
		counter++

		if counter%2 == 0 || fast == nil {
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
