/*
 * @Author: your name
 * @Date: 2020-10-20 21:53:00
 * @LastEditTime: 2020-10-23 17:53:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /DataStruct/golang-DataStructAndAlgorithm/Sort/HeapSort.go
 */
package Sort

type Node struct {
	data           int
	lchild, rchild *Node
}

var FullBinaryTree Node

func HeapSort(params []int) Node {
	n := initFullBinaryTree(params)
	// fmt.Println(n)
	return n
}

func initFullBinaryTree(params []int) Node {
	//大顶堆 params[i] >= params[2i+1] && parasm[i] >= params[2i+2]
	//小顶堆 params[i] <= params[2i+1] && parasm[i] <= params[2i+2]
	node := Node{}
	// n := len(params)
	return node
}
