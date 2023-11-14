package main

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
func Init() *Node {
	return &Node{}
}

// Destroy destroys the linked list. No operations are permitted to
// be called on the linked list after Destroy is called
func Destroy(l *LinkedList) error {
	l = nil
	return nil
}

// InsertNext inserts a node after the current linked list element.
// If the element is nil then the new element gets placed at the head
// of the list.
func (l *LinkedList) InsertNext() error {
	return nil
}

// RemoveNext removes the next element after element from the list.
func (l *LinkedList) RemoveNext(element *Node) error {
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
func (l *LinkedList) IsHead(element *Node) bool {
	return true
}

// IsTail determines if element is the tail of the list.
func (l *LinkedList) IsTail(element *Node) bool {
	return true
}

// Evaluate returns the data of element's Node
func (l *LinkedList) Evaluate(element *Node) *interface{} {
	return element.Data
}

// NextNode returns the next Node in the list.
func (l *LinkedList) NextNode(element *Node) *Node {
	return element.Next
}
