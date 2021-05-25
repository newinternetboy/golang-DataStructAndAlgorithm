package LinkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList_InitList(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3, 4}
	list.InitList(params)
	for _, v := range params {
		if list.SearchValue(v) == false {
			t.Fail()
		}
	}
}
func TestLinkedList_InsertLast(t *testing.T) {
	var list LinkedList
	params := []int{1, 2}
	for _, v := range params {
		list.InsertLast(v)
	}
	for pi, pv := range params {
		if list.GetItems()[pi] != pv {
			t.Fail()
		}
	}
}

func TestLinkedList_InsertFirst(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3}
	for _, v := range params {
		list.InsertFirst(v)
	}
	length := len(params)
	//reverse params
	var params_res []int
	for i := length - 1; i >= 0; i-- {
		params_res = append(params_res, params[i])
	}

	for ri, rv := range params_res {
		if list.GetItems()[ri] != rv {
			t.Fail()
		}
	}
}

func TestLinkedList_RemoveByValue(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3, 4}
	list.InitList(params)
	list.RemoveByValue(1)
	for _, v := range list.GetItems() {
		if v == 1 {
			t.Fail()
		}
	}
}

func TestLinkedList_RemoveByIndex(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3, 4}
	list.InitList(params)
	list.RemoveByIndex(5)
}

func TestLinkedList_SearchValue(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3, 4}
	list.InitList(params)
	if list.SearchValue(1) != true {
		t.Fail()
	}
}

func TestLinkedList_GetFirst(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3, 4}
	list.InitList(params)
	if v, _ := list.GetFirst(); v != 1 {
		t.Fail()
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3, 4}
	list.InitList(params)
	if v, _ := list.GetLast(); v != 4 {
		t.Fail()
	}
}

func TestLinkedList_GetSize(t *testing.T) {
	var list LinkedList
	params := []int{1, 2, 3, 4}
	list.InitList(params)
	if list.GetSize() != 4 {
		t.Fail()
	}
}

func TestLinkedRevert(t *testing.T) {
	var list *Node
	list = &Node{data: 1}
	list.next = &Node{data: 2}
	list = list.LinkedRevert()
	head := list
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
