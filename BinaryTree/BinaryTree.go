/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-21 16:03:09
 * @LastEditors: Caoxiang
 */
package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
)

type BinaryTree struct {
	val   interface{}
	left  *BinaryTree
	right *BinaryTree
}

type ListStack struct {
	list *list.List
	rw   *sync.RWMutex
}

func NewStack() *ListStack {
	list := list.New()
	l := &sync.RWMutex{}
	return &ListStack{list: list, rw: l}
}

func (ls *ListStack) Push(i interface{}) {
	ls.rw.Lock()
	defer ls.rw.Unlock()
	ls.list.PushBack(i)
}

func (ls *ListStack) Pop() interface{} {
	ls.rw.Lock()
	defer ls.rw.Unlock()
	e := ls.list.Back()
	if e != nil {
		ls.list.Remove(e)
		return e.Value
	}
	return nil
}

var level int

func CreateTree(bt *BinaryTree, list *ListStack) *BinaryTree {
	if list.list.Len() < 1 {
		return bt
	}
	firstEle := list.Pop()
	fmt.Println("pop:", firstEle)
	if bt == nil {
		bt = &BinaryTree{val: firstEle}
	} else {
		bt.val = firstEle
	}

	//判断下一个元素插入到左右子节点中的哪一个
	//左小右大
Loop:
	next := list.list.Back()
	if next != nil {
		nextstring := list.list.Back().Value.(string)
		nextdata, err := strconv.Atoi(nextstring)
		v, err := strconv.Atoi(fmt.Sprint(firstEle))
		if err == nil {
			if v == nextdata {
				list.list.Remove(next)
				goto Loop
			}

			if v < nextdata {
				bt.left = CreateTree(bt.left, list)
			} else {
				bt.right = CreateTree(bt.right, list)
			}
		}
	}
	return bt
}

func main() {
	var bt *BinaryTree
	path := []string{"2", "1", "1", "4", "4", "3"}
	stack := NewStack()
	for _, v := range path {
		stack.Push(v)
	}
	bt = CreateTree(bt, stack)
	PreOrderTraveral(bt, level)
}

func PreOrderTraveral(bt *BinaryTree, level int) {
	if bt == nil {
		return
	}
	level++
	fmt.Printf("level:%d,%s\n", level, bt.val)
	//输出左右子结点
	PreOrderTraveral(bt.left, level)
	PreOrderTraveral(bt.right, level)
}
