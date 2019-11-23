// 二分查找算法，已经排序的数组适用,正序
package BinarySearch

import "fmt"

func BinarySearch(array []int, number int) int {
	sortArray(array)
	fmt.Println(array)
	//二分查找
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex := int((minIndex + maxIndex) / 2)
		midItem  := array[midIndex]
		if number == midItem {
			return midIndex
		}

		if midItem < number {
			minIndex = midIndex + 1
		}else if midItem > number {
			maxIndex = midIndex - 1
		}
	}
	return -1
}

func sortArray(array []int)  {
	for itemIndex, itemValue := range array {
		//冒泡排序
		for itemIndex != 0 && array[itemIndex-1] > itemValue {
			array[itemIndex] = array[itemIndex-1]
			itemIndex -= 1
		}
		array[itemIndex] = itemValue
	}
}
