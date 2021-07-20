/*
 * @Description: list.List双向链表+map实现lru
 * @Author: Caoxiang
 * @Date: 2021-07-20 15:05:53
 * @LastEditors: Caoxiang
 */
package main

import (
	"container/list"
	"fmt"
)

type LRUCacheFunc interface {
	Size() int
	Set(k, v interface{})
	Get(k interface{}) (v interface{}, has bool)
	Remove(k interface{}) bool
}

/**
List结点类型
*/
type CacheNode struct {
	key, value interface{}
}

//LRUCache
type LRUCache struct {
	Capacity int
	list     *list.List
	cacheMap map[interface{}]*list.Element
}

func (lru *LRUCache) NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		list:     list.New(),
		cacheMap: make(map[interface{}]*list.Element),
	}
}

func (lru *LRUCache) Size() int {
	return lru.list.Len()
}

func (lru *LRUCache) Set(k, v interface{}) {
	if lru.list == nil {
		panic("LRUCache未初始化")
	}
	if ele, ok := lru.cacheMap[k]; ok {
		lru.list.MoveToFront(ele)
		return
	}

	node := &CacheNode{key: k, value: v}
	newElement := lru.list.PushFront(node)
	lru.cacheMap[k] = newElement

	if lru.list.Len() > lru.Capacity {
		//获取最后链表最后一个元素
		lastElement := lru.list.Back()
		if lastElement != nil {
			cacheNode := lastElement.Value.(*CacheNode)
			delete(lru.cacheMap, cacheNode.key)
			lru.list.Remove(lastElement)
		}
	}
}

func (lru *LRUCache) Get(k interface{}) (v interface{}, exist bool) {
	if lru.cacheMap == nil {
		return v, false
	}
	if ele, ok := lru.cacheMap[k]; ok {
		return ele.Value.(*CacheNode).value, true
	}
	return v, false
}

func (lru *LRUCache) Remove(k interface{}) bool {
	if lru.cacheMap == nil {
		return false
	}
	if ele, ok := lru.cacheMap[k]; ok {
		cacheNode := ele.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.key)
		lru.list.Remove(ele)
		return true
	}
	return false
}

func Output(lru *LRUCache) {
	n := lru.list.Front()
	for n != nil {
		fmt.Println(n.Value)
		n = n.Next()
	}
}

func main() {
	lru := new(LRUCache).NewLRUCache(3)
	lru.Set("name", "caoxiang")
	lru.Set("age", 29)
	lru.Set("city", "xinyang")
	if v, ok := lru.Get("city"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("city don`t exist")
	}
	Output(lru)
	fmt.Println("####full#####")
	lru.Set("hobby", "music")
	Output(lru)
}
