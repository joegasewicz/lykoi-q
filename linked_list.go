package main

import "errors"

type LinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

// Node is the first node in the list
type Node struct {
	Next *Node
	Data *interface{}
}

// Init should be called before any operation can be used on with
// the linked list
func Init(data *interface{}) *LinkedList {
	n := Node{
		Data: data,
		Next: nil,
	}
	return &LinkedList{
		Size: 1,
		Head: &n,
	}
}

// Destroy destroys the linked list. No operations are permitted to
// be called on the linked list after Destroy is called
func Destroy(l *LinkedList) error {
	l = nil
	if l != nil {
		return errors.New("expected list to be nil")
	}
	return nil
}

// InsertNext inserts a node after the current linked list element.
// If the element is nil then the new element gets placed at the head
// of the list & then return the tail.
func (l *LinkedList) InsertNext(currNode *Node, data *interface{}) (error, *Node) {
	var node *Node
	newNode := Node{
		Next: nil,
		Data: data,
	}
	if currNode.Data == l.Head.Data {
		if l.Head.Next != nil {
			err := errors.New("cannot add element to list as next element is not empty")
			return err, l.Head.Next
		}
		l.Head.Next = &newNode
		return nil, l.Head.Next
	}
	node = l.Head.Next
	for {
		if currNode.Data == node.Data || node.Next == nil {
			node.Next = &newNode
			break
		}
		node = node.Next
	}
	return nil, node
}

// RemoveNext removes the next element after element from the list.
func (l *LinkedList) RemoveNext(node *Node) error {
	return nil
}

// ListSize returns the total number of nodes in the list.
func (l *LinkedList) ListSize() int {
	return 0
}

// ListHead returns the head of the list
func (l *LinkedList) ListHead() *Node {
	return l.Head
}

// ListTail returns the tail node of the list.
func (l *LinkedList) ListTail() *Node {
	return nil
}

// IsHead determines if element is the head of the list.
func (l *LinkedList) IsHead(node *Node) bool {
	return true
}

// IsTail determines if element is the tail of the list.
func (l *LinkedList) IsTail(node *Node) bool {
	return true
}

// Evaluate returns the data of element's Node
func (l *LinkedList) Evaluate(node *Node) *interface{} {
	return node.Data
}

// NextNode returns the next Node in the list.
func (l *LinkedList) NextNode(node *Node) *Node {
	return node.Next
}
