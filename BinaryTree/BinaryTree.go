package BinaryTree

type Node struct {
	data 	int
	parent 	*Node
	left	*Node
	right	*Node
}

type BinaryTree struct {
	root *Node
}

