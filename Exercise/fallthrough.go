/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-07-06 15:06:56
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	switch {
	case 1 == 1:
		fmt.Println("1=1 is true")
		fallthrough
	case 2 == 2:
		fmt.Println("2=2 is true")
	}
}
