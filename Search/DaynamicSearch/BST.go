/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2020-10-27 16:11:27
 * @LastEditors: Caoxiang
 * @LastEditTime: 2020-10-27 16:12:02
 */
package DaynamicSearch

//二叉排序树的定义
type BTNode struct {
	key            int
	lchild, rchild *BTNode
}

//插入
func (bt *BTNode) Insert(value int) *BTNode {
	if bt == nil {
		return &BTNode{key: value}
	}
	if value < bt.key {
		bt.lchild = bt.lchild.Insert(value)
	} else if value > bt.key {
		bt.rchild = bt.rchild.Insert(value)
	}
	return bt
}

//删除操作
//删除的结点就是叶子结点，没有左右子树，操作:直接删除
//删除的结点只有一个子节点(单左，单右)操作:将子节点挂到删除结点的位置
//删除的结点有左右子树，操作:使用右子树最小的结点，删除结点的直接后继(或者左子树的最大结点,删除结点的直接前驱)替换删除结点，并删除最小结点(或者最大结点)
//注:这里的直接前驱和后继是中序遍历

func (bt *BTNode) Delete(value int) *BTNode {
	if bt == nil {
		return bt
	}
	if value < bt.key {
		bt.lchild = bt.lchild.Delete(value)
	} else if value > bt.key {
		bt.rchild = bt.rchild.Delete(value)
	} else {
		if bt.lchild != nil && bt.rchild != nil {
			bt.key = bt.rchild.GetMin()
			bt.rchild = bt.rchild.Delete(bt.key)
		} else if bt.lchild != nil {
			bt = bt.lchild
		} else {
			bt = bt.rchild
		}
	}
	return bt
}

//获取右子树最小值
func (bt *BTNode) GetMin() int {
	if bt == nil {
		return -1
	}
	for bt.rchild != nil {
		bt = bt.rchild
	}
	return bt.key
}

//获取左子树最大值
func (bt *BTNode) GetMax() int {
	if bt == nil || bt.lchild == nil {
		return -1
	}
	return bt.lchild.key
}

//按序获取所有的数据
func (bt *BTNode) GetAll() []int {
	list := []int{}
	list = MidOrder(list, bt)
	return list
}

//中序遍历
func MidOrder(values []int, bt *BTNode) []int {
	if bt != nil {
		values = MidOrder(values, bt.lchild)
		values = append(values, bt.key)
		values = MidOrder(values, bt.rchild)
	}
	return values
}

func (bt *BTNode) Search(value int) bool {
	if bt == nil {
		return false
	}
	if value < bt.key {
		return bt.rchild.Search(value)
	} else if value > bt.key {
		return bt.lchild.Search(value)
	}
	return true
}

func (bt *BTNode) CreateBST(params []int) *BTNode {
	n := len(params)
	for i := 0; i < n; i++ {
		bt = bt.Insert(params[i])
	}
	return bt
}
