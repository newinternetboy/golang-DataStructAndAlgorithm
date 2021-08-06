package LinkedList

//单链表
//当尾结点指向头结点就是单向循环链表
//define node
type Node struct {
	data int
	next *Node
}

//define linklist
type LinkedList struct {
	head *Node
}

//Init LinkList
func (list *LinkedList) InitList(params []int) {
	for _, v := range params {
		list.InsertLast(v)
	}
}

//Insert First(Linklist implement InsertFirst)
func (list *LinkedList) InsertFirst(i int) {
	//Init Insert Node
	data := &Node{data: i}
	if list.head != nil {
		data.next = list.head
	}
	list.head = data
}

//Insert Last
func (list *LinkedList) InsertLast(i int) {
	dataNode := &Node{data: i}
	if list.head == nil {
		list.head = dataNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = dataNode
}

//Remove By Value
func (list *LinkedList) RemoveByValue(i int) bool {
	//First Node
	if list.head == nil {
		return false
	}
	if list.head != nil && list.head.data == i {
		//删除头结点
		list.head = list.head.next
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			current.next = current.next.next
			return true
		}
	}
	return false
}

//Remove By Index
func (list *LinkedList) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	current.next = current.next.next
	return true
}

//Search Value
func (list *LinkedList) SearchValue(i int) bool {
	current := list.head
	for current != nil {
		if current.data == i {
			return true
		}
		current = current.next
	}
	return false
}

//Get First
func (list *LinkedList) GetFirst() (int, bool) {
	head := list.head
	if head == nil {
		return 0, false
	}
	return head.data, true
}

//Get Last
func (list *LinkedList) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

//Get Size
func (list *LinkedList) GetSize() int {
	size := 0
	current := list.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

//Get Items
func (list *LinkedList) GetItems() []int {
	var items []int
	if list.head == nil {
		return items
	}
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

/**
 * @description: 单链表反转
 * @param {*}
 * @return {*}
 */
func (list *Node) LinkedRevert() *Node {
	//双指针
	cur := list
	var pre *Node
	for cur != nil {
		tmp := cur.next
		//指针反转
		cur.next = pre
		//双指针更新
		pre = cur
		cur = tmp
	}
	return pre
}

func (list *LinkedList) RemoveElement(val int) *LinkedList {
	//设置虚拟头结点
	dummyHead := &Node{}
	dummyHead.next = list.head
	cur := dummyHead
	for cur != nil && cur.next != nil {
		if cur.next.data == val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
	list.head = dummyHead.next
	return list
}
