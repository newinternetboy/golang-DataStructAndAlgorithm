package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	ln := &ListNode{val: 1, next: &ListNode{val: 2}}
	nln := rmsamevalnode(ln, 1)
	fmt.Println(nln)
}

func rmsamevalnode(ln *ListNode, v int) *ListNode {
	//设置虚拟头结点,统一删除逻辑
	dumphead := &ListNode{}
	dumphead.next = ln
	cur := dumphead
	for cur.next != nil {
		if cur.next.val == v {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
	ln = dumphead.next
	return ln
}
