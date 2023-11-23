package main

import (
	"errors"
	"log"
)

type Queue struct {
	ShouldDestroy bool
	LinkedList    *LinkedList
	tailNode      *Node
}

// QueueDestroy destroys the queue. No operations are permitted after calling this method.
func QueueDestroy(q *Queue) (err error) {
	if q.ShouldDestroy != false {
		err = LinkedListDestroy(q.LinkedList)
	}
	q = nil
	if q != nil {
		return errors.New("could not delete queue memory")
	}
	return err
}

// QueueInit method must be called before the queue can be used.
// If the caller wants to destroy queue data then shouldDestroy parameter
// should be set to true & false if the data should not be destroyed.
func QueueInit(shouldDestroy bool) *Queue {
	return &Queue{ShouldDestroy: shouldDestroy}
}

// Peak inspect the head of the queue without removing it.
// Returns nil if the queue is empty
func (q *Queue) Peak() *interface{} {
	head := q.LinkedList.ListHead()
	if head == nil {
		return nil
	}
	return head.Data
}

// Enqueue enqueue - place element at tail of queue
func (q *Queue) Enqueue(data *interface{}) error {
	if q.LinkedList == nil {
		ll := LinkedListInit(data)
		q.LinkedList = ll
		q.tailNode = ll.ListTail()
		return nil
	}
	if q.LinkedList == nil {
		log.Fatal("expected linked list to exist but got nil")
	}
	err, tailNode := q.LinkedList.InsertNext(q.tailNode, data)
	if err != nil {
		return err
	}
	q.tailNode = tailNode
	return nil
}

// Dequeue removes element at head of queue, then returns an error
// if one exists & the head's data.
func (q *Queue) Dequeue() (error, *interface{}) {
	return q.LinkedList.RemoveNext(nil)
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return q.LinkedList.ListSize()
}
