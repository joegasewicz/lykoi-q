package main

import "testing"

func TestInit(t *testing.T) {
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
