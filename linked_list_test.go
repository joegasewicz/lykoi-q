package main

import "testing"

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
	data4 = "data 4 value"
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
	// TODO continue testing this...
}
