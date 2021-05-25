/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-21 16:23:06
 * @LastEditors: Caoxiang
 */
package Sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	params := []int{2, 3, 1, 6, 5}
	MyInsertSort(params)
	t.Log(params)
}

func MyInsertSort(params []int) []int {
	n := len(params)
	var tmp int
	for i := 1; i < n; i++ {
		tmp = params[i]
		j := i - 1
		//将tmp插入[0,i-1]有序列表中
		for j >= 0 && tmp < params[j] {
			params[j+1] = params[j]
			j--
		}
		params[j+1] = tmp
	}
	return params
}
