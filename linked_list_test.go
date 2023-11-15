package main

import (
	"testing"
)

func TestLinkedList_Init(t *testing.T) {
	var data interface{}
	data = 1
	list := Init(&data)
	if list.Size != 1 {
		t.Fatalf("Expected 1 but got %v", list.Size)
	}
	if list.Head.Next != nil {
		t.Fatalf("Expected nil but got %v", list.Head)
	}
	if *list.Head.Data != data {
		t.Fatalf("Expected 1 but got %v", list.Head.Data)
	}
	if list.Head.Next != nil {
		t.Fatalf("Expected nil but got %v", list.Head.Next)
	}
}

func TestLinkedList_Destroy(t *testing.T) {
	var data interface{}
	data = map[string]string{"name": "Joe"}
	list := Init(&data)
	err := Destroy(list)
	if err != nil {
		t.Fatalf("Expected nil but got error: %v", err.Error())
	}
}

func TestLinkedList_InsertNext(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	node := list.Head
	err, tail1 := list.InsertNext(node, &data2)
	err2, tail2 := list.InsertNext(tail1, &data3)
	err3, _ := list.InsertNext(tail2, &data4)
	if err != nil {
		t.Fatalf("expected nil but got %e", err)
	}
	if err2 != nil {
		t.Fatalf("expected nil but got %e", err2)
	}
	if err3 != nil {
		t.Fatalf("expected nil but got %e", err3)
	}
	if list.Head.Next.Data != &data2 {
		t.Fatalf("expected %v to equal data2 but got %v", &data2, list.Head.Next.Data)
	}
	if list.Head.Next.Next.Data != &data3 {
		t.Fatalf("expected %v to equal data3 but got %v", &data3, list.Head.Next.Next.Data)
	}
	if list.Head.Next.Next.Next.Data != &data4 {
		t.Fatalf("expected %v to equal data3 but got %v", &data4, list.Head.Next.Next.Next.Data)
	}
}

func TestLinkedList_RemoveNext(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	node := list.Head
	_, tail1 := list.InsertNext(node, &data2)
	_, tail2 := list.InsertNext(tail1, &data3)
	list.InsertNext(tail2, &data4)
	err := list.RemoveNext(tail2)

	if err != nil {
		t.Fatalf("expected nil but got %e", err)
	}
	if list.Head.Next.Data != &data2 {
		t.Fatalf("expected %v to equal data2 but got %v", &data2, list.Head.Next.Data)
	}
	if list.Head.Next.Next.Data != &data3 {
		t.Fatalf("expected %v to equal data3 but got %v", &data3, list.Head.Next.Next.Data)
	}
	if node.Next.Next.Next != nil {
		t.Fatalf("expected nil but got %v", node.Next.Next.Next)
	}
	list.RemoveNext(nil)
	if list.Head != nil {
		t.Fatalf("expect Head of list to be nil but got %v", list.Head)
	}

}

func TestLinkedList_ListSize(t *testing.T) {
	var count int
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	node := list.Head
	count = list.ListSize()
	if count != 1 {
		t.Fatalf("expected 1 but got %d", count)
	}
	_, tail1 := list.InsertNext(node, &data2)
	count = list.ListSize()
	if count != 2 {
		t.Fatalf("expected 2 but got %d", count)
	}
	_, tail2 := list.InsertNext(tail1, &data3)
	count = list.ListSize()
	if count != 3 {
		t.Fatalf("expected 3 but got %d", count)
	}
	list.InsertNext(tail2, &data4)
	count = list.ListSize()
	if count != 4 {
		t.Fatalf("expected 4 but got %d", count)
	}
	list.RemoveNext(tail2)
	count = list.ListSize()
	if count != 3 {
		t.Fatalf("expected 3 but got %d", count)
	}
	list.RemoveNext(tail1)
	count = list.ListSize()
	if count != 2 {
		t.Fatalf("expected 2 but got %d", count)
	}
}

func TestLinkedList_ListHead(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	list := Init(&data1)
	head := list.Head
	result := list.ListHead()
	list.InsertNext(head, &data2)
	if result != head {
		t.Fatalf("expected %v to equal %v", head, result)
	}
}

func TestLinkedList_ListTail(t *testing.T) {
	var tail *Node
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	node := list.Head
	tail = list.ListTail()
	if tail != node {
		t.Fatalf("expected %v to equal %v", node, tail)
	}
	_, tail1 := list.InsertNext(node, &data2)
	tail = list.ListTail()
	if tail != tail1 {
		t.Fatalf("expected %v to equal %v", node, tail1)
	}
	_, tail2 := list.InsertNext(tail1, &data3)
	tail = list.ListTail()
	if tail != tail2 {
		t.Fatalf("expected %v to equal %v", node, tail2)
	}
	_, tail3 := list.InsertNext(tail2, &data4)
	tail = list.ListTail()
	if tail != tail3 {
		t.Fatalf("expected %v to equal %v", node, tail3)
	}
}

func TestLinkedList_IsHead(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	head := list.Head
	if list.IsHead(head) != true {
		t.Fatalf("expected true but got false")
	}
	_, tail1 := list.InsertNext(head, &data2)
	if list.IsHead(tail1) != false {
		t.Fatalf("expected true but got true")
	}
	_, tail2 := list.InsertNext(tail1, &data3)
	list.InsertNext(tail2, &data4)
	if list.IsHead(tail2) != false {
		t.Fatalf("expected true but got true")
	}
}

func TestLinkedList_IsTail(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	head := list.Head
	if list.IsTail(head) != true {
		t.Fatalf("expected true but got false")
	}
	_, tail1 := list.InsertNext(head, &data2)
	if list.IsTail(tail1) != true {
		t.Fatalf("expected true but got false")
	}
	if list.IsTail(head) != false {
		t.Fatalf("expected true but got true")
	}
	_, tail2 := list.InsertNext(tail1, &data3)
	if list.IsTail(tail2) != true {
		t.Fatalf("expected true but got false")
	}
	if list.IsTail(tail1) != false {
		t.Fatalf("expected true but got true")
	}
	_, tail3 := list.InsertNext(tail2, &data4)
	if list.IsTail(tail3) != true {
		t.Fatalf("expected true but got false")
	}
	if list.IsTail(tail2) != false {
		t.Fatalf("expected true but got true")
	}
}

func TestLinkedList_Evaluate(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	node := list.Head
	_, result := list.Evaluate(node)
	if result != node.Data {
		t.Fatalf("expected %v but got %v", node, result)
	}
	_, tail1 := list.InsertNext(node, &data2)
	_, result = list.Evaluate(tail1)
	if result != tail1.Data {
		t.Fatalf("expected %v but got %v", tail1.Data, result)
	}
	_, tail2 := list.InsertNext(tail1, &data3)
	list.InsertNext(tail2, &data4)
	_, result = list.Evaluate(tail2)
	if result != tail2.Data {
		t.Fatalf("expected %v but got %v", tail2.Data, result)
	}
}

func TestLinkedList_NextNode(t *testing.T) {
	var data1 interface{}
	var data2 interface{}
	var data3 interface{}
	var data4 interface{}
	data1 = map[string]string{"name": "Joe"}
	data2 = map[string]string{"age": "99"}
	data3 = map[string]string{"cat": "cosmo"}
	data4 = "data 4 test"
	list := Init(&data1)
	node := list.Head
	_, tail1 := list.InsertNext(node, &data2)
	_, result := list.NextNode(node)
	if result != tail1 {
		t.Fatalf("expected %v but got %v", tail1, result)
	}
	_, tail2 := list.InsertNext(tail1, &data3)
	_, result = list.NextNode(tail1)
	if result != tail2 {
		t.Fatalf("expected %v but got %v", tail2, result)
	}
	_, tail3 := list.InsertNext(tail2, &data4)
	_, result = list.NextNode(tail2)
	if result != tail3 {
		t.Fatalf("expected %v but got %v", tail3, result)
	}
}
