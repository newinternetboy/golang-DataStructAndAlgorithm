/*
 * @Description: 分治思想的一些应用
 * @Author: Caoxiang
 * @Date: 2020-11-02 14:34:35
 * @LastEditors: Caoxiang
 * @LastEditTime: 2020-11-02 14:39:20
 */
package algorithmIdea

import "fmt"

func Print(params []int, L, R int) {
	n := len(params)
	if n == 0 || L > R || L > n-1 || R > n-1 {
		return
	}
	if L == R {
		fmt.Println(params[L])
	} else {
		fmt.Println(params[L])
		Print(params, L+1, R)
	}
}
