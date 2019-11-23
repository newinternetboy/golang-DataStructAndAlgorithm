package LinearList

import (
	"testing"
)

func TestInitLinearListContent(t *testing.T) {
	linearList := InitLinearListContent()
	if len(linearList.LinearListContent) != 20 {
		t.Fail()
	}
}

func TestLinearList_IsEmpty(t *testing.T) {
	linearList := InitLinearListContent()
	if linearList.IsEmpty() == false{
		t.Fail()
	}
}

func TestLinearList_IsFull(t *testing.T) {
	linearList := InitLinearListContent()
	if linearList.IsFull() == true {
		t.Fail()
	}
}

func TestLinearList_IndexValid(t *testing.T) {
	linearList := InitLinearListContent()
	if linearList.IsIndexValid(0) == false {
		t.Fail()
	}
	if linearList.IsIndexValid(20) == true {
		t.Fail()
	}
}

func TestLinearList_InsertValueByIndex(t *testing.T) {
	linearList := InitLinearListContent()
	linearList.InsertValueByPositionIndex(1,1)
	linearList.InsertValueByPositionIndex(2,2)
	linearList.InsertValueByPositionIndex(3,3)
	if linearList.Length != 3 {
		t.Fail()
	}

	for i, v := range linearList.LinearListContent {
		if i == 0 && v != 1 {
			t.Fail()
		}else if i == 1 && v != 2 {
			t.Fail()
		}else if i == 2 && v != 3 {
			t.Fail()
		}
	}
}

func TestLinearList_DelByIndex(t *testing.T) {
	linearList := InitLinearListContent()
	linearList.InsertValueByPositionIndex(1,1)
	linearList.InsertValueByPositionIndex(2,2)
	if linearList.Length != 2 {
		t.Fail()
	}
	if linearList.DelByIndex(1) != true && linearList.Length != 1 {
		t.Fail()
	}
}

func TestLinearList_Append(t *testing.T) {
	linearList := InitLinearListContent()
	params := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	for _, v := range params {
		linearList.Append(v)
	}
	linearList.Append(21)
	for li, _ := range linearList.LinearListContent {
		if li == linearList.Length && linearList.LinearListContent[linearList.Length] != 21 {
			t.Fail()
		}
	}
}