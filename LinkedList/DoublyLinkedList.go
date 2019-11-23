package LinkedList

//define DoubleLinkedNode
type DoubleLinkedNode struct {
	prev 	*DoubleLinkedNode
	element interface{}
	next 	*DoubleLinkedNode
}

//define DoubleLinkedList
type DoubleLinkedList struct {
	head *DoubleLinkedNode
	tail *DoubleLinkedNode
	length int
}

//init sliently

//append  Insert Last
func (list *DoubleLinkedList) Append(i interface{})  {
	eleNode := &DoubleLinkedNode{element:i}
	if list.GetSize() == 0 {
		list.head = eleNode
		list.tail = eleNode
	}else{
		//Insert Last in list
		//elenode operate
		eleNode.prev  = list.tail
		list.tail.next = eleNode
		//elenode be the last one
		list.tail = eleNode
	}
	list.length++
}

//get size of list
func (list *DoubleLinkedList) GetSize() int  {
	return list.length
}

//get Next Node
func (n *DoubleLinkedNode) GetNext() *DoubleLinkedNode  		{
	return n.next
}

//isTail
func (list *DoubleLinkedList) isTail(n *DoubleLinkedNode) bool  {
	if list.GetTail() == n {
		return true
	}else{
		return false
	}
}

//Get Head
func (list *DoubleLinkedList) GetHead() *DoubleLinkedNode  {
	return list.head
}

//is Head
func (list *DoubleLinkedList) isHead(n *DoubleLinkedNode) bool  {
	if list.GetHead() != n {
		return false
	}else{
		return true
	}
}

//GetTail Node
func (list *DoubleLinkedList) GetTail() *DoubleLinkedNode {
	return list.tail
}

//Insert Next Node
func (list *DoubleLinkedList) InsertNext(n *DoubleLinkedNode, i interface{}) bool {
	if n == nil {
		return false
	}

	if list.isTail(n) == true {
		//Insert Last
		list.Append(i)
	}else{
		//Insert after this Node
		//operate insert node
		iNode := &DoubleLinkedNode{element:i}
		iNode.prev = n
		iNode.next = n.next

		//operate n node
		n.next = iNode

		//operate n next node
		n.next.prev = iNode
	}
	list.length++
	return true
}

//Insert Prev Node
func (list *DoubleLinkedList) InsertPrev(n *DoubleLinkedNode, i interface{}) bool  {
	if n == nil {
		return false
	}

	iNode := &DoubleLinkedNode{element:i}
	if list.isHead(n) == true {
		//operate iNode
		iNode.next = list.head
		//operate head Node
		list.head.prev = iNode
		//iNode be the head Node
		list.head = iNode
	}else{
		//InsertNext after n prev node
		prevNode := n.prev
		list.InsertNext(prevNode, i)
	}
	list.length++
	return true
}

//RemoveNode
func (list *DoubleLinkedList) RemoveNode(n *DoubleLinkedNode) interface{}  {
	if n == nil || (n.prev == nil && n.next == nil) {
		//not node in doubleLinkedList
		return false
	}

	//find next
	if list.isHead(n) == true {
		list.head = n.next
	}else {
		n.prev.next = n.next
	}

	//find prev
	if list.isTail(n) == true{
		list.tail = n.prev
	}else{
		n.next.prev = n.prev
	}

	list.length--
	return n.element
}

//Search Node by element
func (list *DoubleLinkedList) Search(i interface{}) *DoubleLinkedNode {
	if list.GetSize() == 0 {
		return nil
	}
	node := list.GetHead()
	find := false
	for ; node != nil; node = node.next {
		if node.element == i {
			find = true
			break
		}
	}
	if find == false {
		return nil
	}
	return node
}