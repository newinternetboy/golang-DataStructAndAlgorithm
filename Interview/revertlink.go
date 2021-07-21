/*
 * @Description:反转链表
 * @Author: Caoxiang
 * @Date: 2021-07-21 16:42:00
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

type Node struct {
	Val  int
	next *Node
}

func main() {
	node := &Node{Val: 1, next: &Node{Val: 2, next: &Node{Val: 3}}}
	revertNode := node.DoublePointRevert()
	for revertNode != nil {
		fmt.Println(revertNode.Val)
		revertNode = revertNode.next
	}
}

//双指针
func (l *Node) DoublePointRevert() *Node {
	if l == nil || l.next == nil {
		return l
	}

	var pre *Node
	cur := l
	for cur != nil {
		//反转指针+双指针右移
		cur.next, pre, cur = pre, cur, cur.next
	}
	return pre
}
