package main

import (
	"testing"
)

func TestQueueInit(t *testing.T) {
	queue := QueueInit(true)
	if queue.ShouldDestroy == false {
		t.Fatal("expected true but got false")
	}
}

func TestQueue_Destroy(t *testing.T) {
	queue := QueueInit(true)
	err := QueueDestroy(queue)
	if err != nil {
		t.Fatalf("expected error to be nil but got %e", err)
	}
	queue = QueueInit(false)
	err = QueueDestroy(queue)
	if err != nil {
		t.Fatalf("expected error to be nil but got %e", err)
	}
	if queue.ShouldDestroy != false {
		t.Fatal("expected ShouldDestroy to be false but is true")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	queue := QueueInit(true)
	data1 = map[string]string{"name": "joe"}
	data2 = "my name is joe"
	data3 = 2
	err := queue.Enqueue(&data1)
	if err != nil {
		t.Fatalf("expected error to be nil but got %e", err)
	}
	if queue.LinkedList.Head.Data != &data1 {
		t.Fatalf("expected data to be in queue but got %v", queue.LinkedList.Head.Data)
	}

	err2 := queue.Enqueue(&data2)
	if err2 != nil {
		t.Fatalf("expected error to be nil but got %e", err2)
	}
	if queue.LinkedList.Head.Next.Data != &data2 {
		t.Fatalf("expected data to be in queue but got %v", queue.LinkedList.Head.Next.Data)
	}

	err3 := queue.Enqueue(&data3)
	if err3 != nil {
		t.Fatalf("expected error to be nil but got %e", err3)
	}
	if queue.LinkedList.Head.Next.Next.Data != &data3 {
		t.Fatalf("expected data to be in queue but got %v", queue.LinkedList.Head.Next.Next.Data)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	queue := QueueInit(true)
	data1 = map[string]string{"name": "joe"}
	data2 = "my name is joe"
	data3 = 2
	queue.Enqueue(&data1)
	queue.Enqueue(&data2)
	queue.Enqueue(&data3)
	err, headData := queue.Dequeue()
	if err != nil {
		t.Fatalf("expected error to be nil but got %e", err)
	}
	if headData != &data1 {
		t.Fatalf("expected headData to equal %v but got %v", *headData, data1)
	}
	err, headData = queue.Dequeue()
	if err != nil {
		t.Fatalf("expected error to be nil but got %e", err)
	}
	if headData != &data2 {
		t.Fatalf("expected headData to equal %v but got %v", *headData, data2)
	}
	err, headData = queue.Dequeue()
	if err != nil {
		t.Fatalf("expected error to be nil but got %e", err)
	}
	if headData != &data3 {
		t.Fatalf("expected headData to equal %v but got %v", *headData, data3)
	}
}

func TestQueue_Peak(t *testing.T) {
	var data1 interface{}
	queue := QueueInit(true)
	data1 = map[string]string{"name": "joe"}
	queue.Enqueue(&data1)
	peak := queue.Peak()
	if peak != &data1 {
		t.Fatalf("expected peak to be %v but got %v", data1, *peak)
	}
	queue.Dequeue()
	peak = queue.Peak()
	if peak != nil {
		t.Fatalf("expected peak to be nil but got %v", *peak)
	}
}

func TestQueue_Size(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	queue := QueueInit(true)
	data1 = map[string]string{"name": "joe"}
	data2 = "my name is joe"
	data3 = 2
	queue.Enqueue(&data1)
	if queue.Size() != 1 {
		t.Fatalf("expected queue size to be 1 but got %v", queue.Size())
	}
	queue.Enqueue(&data2)
	if queue.Size() != 2 {
		t.Fatalf("expected queue size to be 2 but got %v", queue.Size())
	}
	queue.Enqueue(&data3)
	if queue.Size() != 3 {
		t.Fatalf("expected queue size to be 3 but got %v", queue.Size())
	}
	queue.Dequeue()
	if queue.Size() != 2 {
		t.Fatalf("expected queue size to be 2 but got %v", queue.Size())
	}
	queue.Dequeue()
	if queue.Size() != 1 {
		t.Fatalf("expected queue size to be 1 but got %v", queue.Size())
	}
	queue.Dequeue()
	if queue.Size() != 0 {
		t.Fatalf("expected queue size to be 0 but got %v", queue.Size())
	}
}
