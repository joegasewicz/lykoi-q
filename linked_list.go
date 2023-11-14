package main

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
func Destroy(n *Node) error {
	n = nil
	return nil
}

// InsertNext inserts a node after the current linked list element.
// If the element is nil then the new element gets placed at the head
// of the list.
func (n *Node) InsertNext() error {
	return nil
}

// RemoveNext removes the next element after element from the list.
func (n *Node) RemoveNext(element *Node) error {
	return nil
}

// Size returns the total number of nodes in the list.
func (n *Node) Size() int {
	return 0
}

// Head returns the head of the list
func (n *Node) Head() *Node {
	return n
}

// Tail returns the tail node of the list.
func (n *Node) Tail() *Node {
	return nil
}

// IsHead determines if element is the head of the list.
func (n *Node) IsHead(element *Node) bool {
	return true
}

// IsTail determines if element is the tail of the list.
func (n *Node) IsTail(element *Node) bool {
	return true
}

// Evaluate returns the data of element's Node
func (n *Node) Evaluate(element *Node) *interface{} {
	return element.Data
}

// NextNode returns the next Node in the list.
func (n *Node) NextNode(element *Node) *Node {
	return element.Next
}
