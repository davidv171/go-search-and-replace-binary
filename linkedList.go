package main

import "fmt"

//Useless for now
type Node struct {
	bit  bool
	next *Node
}
type LinkedList struct {
	start *Node
	end   *Node
	//The node that denotes the first one, that we manually reattached when replacing
	//So if we replaced 0101 with 00001010|1| the |1| is the node that we cache so we can access it for futher iterations faster
	cachedNode *Node
}

//Add an element to the end of a linked list
func (linkedList *LinkedList) appendElement(bit bool) {
	//Add an element to the edge of the linked list
	node := &Node{bit, nil}
	linkedList.end.next = node
	//push the end of the linkedList for another element to the right
	linkedList.end = node
}

//For debugging sometimes its easier to jsut use this
func (linkedList *LinkedList) printEverything() {
	node := linkedList.start
	count := 0
	for {
		if node != nil {
			fmt.Print(node.bit)
			fmt.Printf(" %d ", count)
			node = node.next
		}
		count++
	}
}

//Replace the LinkedList at an index with the replacedLinkedList then reattach the end of replacedLinkedList with the reattached node
//Reattached node is kept track of in the
func (linkedList *LinkedList) replaceSublist(replaceLinkedList *LinkedList, startIndex int, replaceSize int, searchSize int) {
	//The cached node is our starting point, initialized to linkedList.start
	startNode := linkedList.cachedNode
	currNode := startNode
	replacableNode := linkedList.cachedNode
	for i := 0; i < startIndex; i++ {
		currNode = currNode.next
		if i == startIndex {
			//find the node that doesn't get replaced anymore
			for j := 0; j <= searchSize; j++ {
				replacableNode = currNode.next
				if j == searchSize {
					//cache the node that doesnt get replaced anymore
					linkedList.cachedNode = replacableNode.next
				}
			}
			//Retie the end of replaced node to the cached node
			//Next iteration, we start from the cached node to avoid looping from the start
			currNode.next = replaceLinkedList.start
			replaceLinkedListEnd := replaceLinkedList.end
			replaceLinkedList.end = linkedList.end
			replaceLinkedListEnd.next = linkedList.cachedNode
		}
	}
}
