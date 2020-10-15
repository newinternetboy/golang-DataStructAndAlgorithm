package LinkedList

import (
	"testing"
)

func TestDoubleLinkedList_GetSize(t *testing.T) {
	var list DoubleLinkedList
	list.Append(1)
	list.Append(2)
	list.Append(3)
	if list.GetSize() != 3 || list.tail.element != 3 {
		t.Fail()
	}
}

func TestDoubleLinkedList_GetTail(t *testing.T) {
	var list DoubleLinkedList
	list.Append(1)
	if list.tail.element != 1 {
		t.Fail()
	}
}

func TestDoubleLinkedList_InsertNext(t *testing.T) {
	var list DoubleLinkedList
	list.Append(1)
	head := list.head
	inr := list.InsertNext(head, 2)
	if inr != true && head.next.element != 2 {
		t.Fail()
	}
}

func TestDoubleLinkedList_InsertPrev(t *testing.T) {
	var list DoubleLinkedList
	list.Append(1)
	list.Append(2)
	nNode := list.GetHead().next
	r := list.InsertPrev(nNode, 2)
	if r != true || list.head.next.element != 2 {
		t.Fail()
	}

	head := list.GetHead()
	hr := list.InsertPrev(head, 3)
	if hr == false || list.GetHead().element != 3 {
		t.Fail()
	}
}

func TestDoubleLinkedList_RemoveNode(t *testing.T) {
	var list DoubleLinkedList
	list.Append(1)
	list.Append(2)

	head := list.GetHead()
	rhead := list.RemoveNode(head)
	if rhead != 1 {
		t.Fail()
	}

	list.Append(3)
	tail := list.GetTail()
	rtail := list.RemoveNode(tail)
	if rtail != 3 {
		t.Fail()
	}

	list.Append(3)
	list.Append(4)
	mNode := list.GetHead().next
	rm := list.RemoveNode(mNode)
	if rm != 3 {
		t.Fail()
	}
}

func TestDoubleLinkedList_Search(t *testing.T) {
	var list DoubleLinkedList
	list.Append(1)
	if list.Search(1).element != 1 {
		t.Fail()
	}
	if list.Search(2) != nil {
		t.Fail()
	}
}
