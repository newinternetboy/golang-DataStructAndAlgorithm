/*
 * @Description:返回二叉树任意最长路径
 * @Author: Caoxiang
 * @Date: 2021-07-22 16:41:43
 * @LastEditors: Caoxiang
 */
package main

import (
	"fmt"
	"strings"
)

type BiTree struct {
	val string
	lch *BiTree
	rch *BiTree
}

var path []string

func main() {
	bt := &BiTree{val: "a",
		rch: &BiTree{val: "d", rch: &BiTree{val: "e", rch: &BiTree{val: "f"}}},
		lch: &BiTree{val: "b", lch: &BiTree{val: "c", lch: &BiTree{val: "h", lch: &BiTree{val: "i"}}}}}
	bt.VisitedPath()
	fmt.Printf("%s\n", strings.Join(path, "->"))
}

func (bt *BiTree) Depth() int {
	if bt == nil {
		return 0
	}
	maxDepthPath := 0
	if bt.lch.Depth() > bt.rch.Depth() {
		maxDepthPath = bt.lch.Depth()
	} else {
		maxDepthPath = bt.rch.Depth()
	}
	return maxDepthPath + 1
}

func (bt *BiTree) VisitedPath() {
	if bt != nil {
		path = append(path, bt.val)
		if bt.lch.Depth() > bt.rch.Depth() {
			bt.lch.VisitedPath()
		} else {
			bt.rch.VisitedPath()
		}
	}
}
