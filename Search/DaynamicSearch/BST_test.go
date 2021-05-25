/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-05-25 17:03:53
 * @LastEditors: Caoxiang
 */
package DaynamicSearch

import (
	"testing"
)

func TestInsert(t *testing.T) {
	var bst *BTNode
	bst = bst.Insert(2)
	bst = bst.Insert(1)
	bst = bst.Insert(3)
	if bst.key != 2 {
		t.Fail()
	}
	if bst.lchild.key != 1 {
		t.Fail()
	}
	if bst.rchild.key != 3 {
		t.Fail()
	}
}
