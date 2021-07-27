/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-07-27 16:15:19
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 3, 4}
	size := removeSameValInArr(arr, 4)
	fmt.Println(size)
}

//移除数组中指定元素，并返回最后新数组的长度
//快慢指针方法
//slowindex会在等于目标值时停止，并将快指针指定的值用来覆盖，慢指针指向的位置
func removeSameValInArr(arr [5]int, val int) int {
	slowIndex := 0
	for _, v := range arr {
		if val != v {
			arr[slowIndex] = v
			slowIndex++
		}
	}
	return slowIndex
}
