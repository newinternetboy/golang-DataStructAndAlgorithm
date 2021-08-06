/*
 * @Description:9*9乘法表
 * @Author: Caoxiang
 * @Date: 2021-07-06 14:56:26
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}
}
