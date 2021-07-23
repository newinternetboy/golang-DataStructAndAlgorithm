/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-07-06 15:02:34
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	for i := 33; i < 122; i++ {
		fmt.Printf("%#U\n", i)
	}
}
