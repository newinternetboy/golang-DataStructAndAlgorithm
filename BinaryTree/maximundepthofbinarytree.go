/*
 * @Description:104
 给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，
 * @Author: Caoxiang
 * @Date: 2021-08-16 10:34:27
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
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	maxDep := maxDepth(tree)
	fmt.Println(maxDep)
}

func maxDepth(root *TreeNode) int {
	maxDep := 0
	if root == nil {
		return maxDep
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)
	if leftMaxDepth > rightMaxDepth {
		maxDep = leftMaxDepth + 1
	} else {
		maxDep = rightMaxDepth + 1
	}
	return maxDep
}
