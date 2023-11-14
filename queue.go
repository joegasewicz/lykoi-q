package main

// H_|_|T
type Queue struct {
	Data *interface{}
}

// Init Initializes the queue
// The Init() method must be called before the queue can be used.
func (q *Queue) Init() {

}

// Destroy destroys the queue. No operations are permitted after calling this method.
func (q *Queue) Destroy() {

}

// Peak inspect without removing
func (q *Queue) Peak() {

}

// Enqueue enqueue - place element at tail of queue
func (q *Queue) Enqueue(data *interface{}) {

}

// Dequeue dequeue - remove element at head of queue
func (q *Queue) Dequeue(data *interface{}) {

}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return 0
}
