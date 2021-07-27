/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-05-27 18:08:47
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

func main() {
	n := &Node{Val: 1, Next: &Node{Val: 2}}
	revertedNode := n.revertLinkedList()
	fmt.Println(revertedNode)
}

func (n *Node) revertLinkedList() *Node {
	if n == nil {
		return n
	}
	var pre *Node
	cur := n
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
