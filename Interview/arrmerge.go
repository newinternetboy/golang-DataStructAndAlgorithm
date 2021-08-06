/*
 * @Description:算法题，二维数组，元素为一个区间，两个区间有重叠返回合并的区间，最终返回无重叠元素的二维数组？(前提无序，那么可以先排序)
 * @Author: Caoxiang
 * @Date: 2021-08-06 14:27:27
 * @LastEditors: Caoxiang
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][2]int{}
	arr = append(arr, [2]int{1, 2}, [2]int{4, 5}, [2]int{2, 3})
	res := MergeArr(arr)
	fmt.Println(res)
}

func MergeArr(arr [][2]int) [][2]int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	arr_map := make(map[int][2]int)
	arr_keys := []int{}
	for _, v := range arr {
		arr_keys = append(arr_keys, v[0])
		arr_map[v[0]] = v
	}
	sort.Ints(arr_keys)
	for i, v := range arr_keys {
		arr[i] = arr_map[v]
	}

	res := [][2]int{}
	//哨兵
	res = append(res, arr[0])
	i := 0
	for _, v := range arr {
		if v[0] <= res[i][1] {
			//合并并覆盖
			res[i] = [2]int{res[i][0], v[1]}
		} else {
			//append
			res = append(res, v)
			i++
		}
	}
	return res
}
