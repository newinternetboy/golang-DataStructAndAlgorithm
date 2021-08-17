/*
 * @Description: 98
 假设一个二叉搜索树具有如下特征：
	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
 * @Author: Caoxiang
 * @Date: 2021-08-16 11:06:02
 * @LastEditors: Caoxiang
*/
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	// root := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	// root := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}
	isValid := isValidBST(root)
	fmt.Println(isValid)
}

func isValidBST(root *TreeNode) bool {
	//递归的方法
	//先校验根结点是否满足搜索二叉树
	//如果根结点不满足，直接放弃遍历
	//如果根结点满足，遍历左结点是否满足搜索二叉树
	//左结点满足
	//遍历右结点是否满足搜索二叉树
	min, max := math.MinInt64, math.MaxInt64
	valid := validBST(root, max, min)
	return valid
}

func validBST(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return validBST(root.Left, root.Val, min) && validBST(root.Right, max, root.Val)
}
