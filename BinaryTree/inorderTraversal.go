/*
 * @Description:
 给定一个二叉树的根节点 root ，返回它的 中序 遍历。

 * @Author: Caoxiang
 * @Date: 2021-08-16 10:48:04
 * @LastEditors: Caoxiang
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// root := &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2}}
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	res := inorderTraversal(root)
	fmt.Println(res)
}

func inorderTraversal(root *TreeNode) []int {
	path := []int{}
	if root == nil {
		return path
	}
	path = append(path, inorderTraversal(root.Left)...)
	path = append(path, root.Val)
	path = append(path, inorderTraversal(root.Right)...)
	return path
}
