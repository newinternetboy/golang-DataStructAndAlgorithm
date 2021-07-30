/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-05-27 18:08:47
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 6, 4, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(sortArray []int, left, right int) {
	if left < right {
		key := sortArray[(left+right)/2]
		i := left
		j := right

		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}

		quickSort(sortArray, left, i-1)
		quickSort(sortArray, j+1, right)
	}
}
