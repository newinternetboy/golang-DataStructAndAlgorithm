/*线性表*/
package LinearList

const MaxLength =  20

//define LinearList
type LinearList struct {
	MaxLength int
	Length	  int
	LinearListContent []interface{}
}

//InitLinearList
func InitLinearListContent() *LinearList {
	return &LinearList{MaxLength:MaxLength,LinearListContent:make([]interface{},MaxLength)}
}

//is Empty
func (list *LinearList) IsEmpty() bool {
	return list.Length == 0
}

//is Full
func (list *LinearList) IsFull() bool {
	return list.Length == list.MaxLength
}

//Index is Valid
func (list *LinearList) IsIndexValid(i int) bool  {
	if i < 0 || i >= list.MaxLength {
		return false
	} else {
		return true
	}
}

//Insert Value By Index
func (list *LinearList) InsertValueByPositionIndex(i int, v int) bool {
	if i < 1 || i > list.Length + 1 {
		return false
	}

	if list.IsFull() == true {
		return false
	}

	for itemIndex := list.Length; itemIndex > i-1; itemIndex-- {
		list.LinearListContent[itemIndex] = list.LinearListContent[itemIndex-1]
	}
	list.LinearListContent[i-1] = v
	list.Length++
	return true
}

func (list *LinearList) DelByIndex(i int) bool {
	if i < 0 || i > list.Length - 1 {
		return false
	}
	if list.IsEmpty() == true {
		return false
	}
	//下标为i后面的数据都要前移
	for itemIndex := i; itemIndex < list.Length; itemIndex++ {
		list.LinearListContent[itemIndex] = list.LinearListContent[itemIndex+1]
	}
	list.Length--
	return true
}

//Insert value in Last Position
func (list *LinearList) Append(v interface{}) bool  {
	if list.IsFull() == true {
		//allocate
		list.LinearListContent = append(list.LinearListContent,nil)
		list.MaxLength++
	}
	list.LinearListContent[list.Length] = v
	list.Length++
	return true
}