/*
 * @Description: 二分查找
 * @Author: Caoxiang
 * @Date: 2021-07-27 11:44:38
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	index := binarySearch(arr, 5)
	fmt.Println(index)
}

/**
 * @description: 二分查找，区间为闭区间
 * @param {[5]int} arr
 * @param {int} target
 * @return {*}
 */
func binarySearch(arr [5]int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := left + (right-left)/2
		if target > arr[middle] {
			left = middle + 1
		} else if target < arr[middle] {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}
