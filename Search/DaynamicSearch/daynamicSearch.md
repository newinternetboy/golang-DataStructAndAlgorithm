# 动态查找表
## 性质
表结构本身是在查找过程中动态生成的，对于查找的key值，如果在表中查找到就返回，没有查找到就插入关键字为key值的节点。

## ADT
ADT dynamicSearchTable {
    数据对象D:  D是具有相同数据元素的集合。各个数据元素均含有类型相同，可唯一标识数据元素的关键字。
    数据关系R:  数据元素同属于一个集合
    基本操作P:
        InitDSTable(&DT);
            操作结果:构造一个空的动态查找表DT
        DestroyDSTable(&DT);
            初始条件:动态查找表存在
            操作结果:销毁动态查找表DT
        SearchDSTable(DT, key);
            初始条件:动态查找表存在，且key和数据元素关键字的类型一致
            操作结果:如果DT中存在其关键字等于key的数据元素，则函数值为该元素的值或在表中的位置，否则为空。
        InsertDSTable(&DT, e);
            初始条件:动态查找表存在，e为待插入的数据元素
            操作结果:若DT中不存在数据元素e，那么将e插入到DT
        DeleteDSTable(&DT, key);        
            初始条件:动态查找表存在，key为和关键字类型相同的给定值。
            操作结果:如果DT中存在数据元素关键字和key值相同的数据元素，删除该数据元素
        TraverseDSTable(DT, visit());
            初始条件:动态查找表DT存在，Visit是对节点操作的函数
            操作结果:按照某种次序对DT的每个结点调用函数visit()一次至多次，一旦visit()失败，则操作失败。    
}

## 二叉排序树
### 定义
一颗空数或者具有下列性质的树
1 如果存在左子树，那么左子树上所有结点的值均小于它的根结点的值
2 如果存在右子树，那么右子树上所有结点的值均大于它的根结点的值
3 左子树和右子树都是二叉排序树

## 存储结构
可采用二叉链表作为二叉排序树的存储结构

## 伪代码
### 只有查询
BiTree SearchBST(BiTree T, KeyType key) {
    if((!T) || EQ(key, T->data.key)) {
        return(T);
    } else if LT(key, T->data.key) {
        return(SearchBST(T->lchild, key))
    } else {
        return(SearchBST(T->rchild, key))
    }
} //SearchBST

### 查询不成功插入
在查找的过程中需要记录查找的结点位置，以便在查找不成功的时候将数据元素的关键字为key的结点插入
Status SearchBST(BiTree T, KeyType key, BiTree f, BiTree &p) {
    //f指向当前查找结点的父节点，初始值为null，主要查找路径上最后一个结点
    //p指针指向查找成功数据元素的结点或者查找路径上最后一个结点
    if(!T) {
        p = f;
        return false;
    } else EQ(key, T->data.key) {
        p = T;
        return true;
    } else LT(key, T->data.key) {
        SearchBST(T->lchild, key, T, p);
    } else {
        SearchBST(T->rchild, key, T, p);
    }
}

Status InsertBST(BiTree &T, ElemType e) {
    //当二叉排序树T中不存在关键字等于e.key的数据元素时，插入e，并返回True
    if(!SearchBST(T, e.key, null, p)) {
        s = (BITree) malloc(sizeof(BiNode))
        s->data = e;
        s->lchild = s->rchild = null;
        if(!p) {
            //空二叉排序树
            T = s;
        } else if LT(e.key, p->data.key) {
            p->lchild = s;
        } else {
            p->rchild = s;
        }
        return true;
    } else {
        return false;
    }
}