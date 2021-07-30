/*
 * @Description:给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 * @Author: Caoxiang
 * @Date: 2021-07-30 14:33:59
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	res := generateMatrix(6)
	fmt.Println(res)
}

func generateMatrix(n int) [][]int {
	//画四条边，左闭右开，每条边的元素递增
	//i,j
	//(0,0)->(0,1)
	//(0,2)->(1,2)
	//(2,2)->(2,1)
	//(2,0)->(1,0)
	res := make([][]int, n)
	for is := 0; is < n; is++ {
		res[is] = make([]int, n)
	}
	fmt.Println(res)
	startx, starty := 0, 0
	mid := n / 2
	//循环次数
	loop := n / 2
	count := 1
	offset := 1
	for loop > 0 {
		i, j := startx, starty
		//左闭右开
		for j < starty+n-offset {
			res[startx][j] = count
			count++
			j++
		}
		//左竖
		for i < startx+n-offset {
			res[i][j] = count
			count++
			i++
		}
		//低横
		for j > starty {
			res[i][j] = count
			count++
			j--
		}
		//右竖
		for i > startx {
			res[i][j] = count
			count++
			i--
		}
		startx++
		starty++
		offset += 2
		loop--
	}

	if n%2 == 1 {
		res[mid][mid] = count
	}

	return res
}
