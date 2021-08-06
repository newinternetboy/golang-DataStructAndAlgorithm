/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-07-29 20:02:05
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	arr := []int{-3, -2, 0, 1, 5}
	res := sqrtedSortArr(arr)
	fmt.Println(res)
}

func sqrtedSortArr(arr []int) []int {
	//双指针
	n := len(arr)
	i, j, k := 0, n-1, n-1
	res := make([]int, n)
	for i <= j {
		ls, rs := arr[i]*arr[i], arr[j]*arr[j]
		if ls > rs {
			res[k] = ls
			i++
		} else {
			res[k] = rs
			j--
		}
		k--
	}
	return res
}
