/*
 * @Description: 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
 * @Author: Caoxiang
 * @Date: 2021-07-30 11:02:39
 * @LastEditors: Caoxiang
 */
package main

import "fmt"

func main() {
	arr := []int{1, 3, -1, 2, 4, 3}
	sum := 8
	minLen := getsuccessivearrmixLen(arr, sum)
	minLen1 := minSubArrayLen(arr, sum)
	minLen2 := getsuccessivearrmixLenV2(arr, sum)
	fmt.Println(minLen, minLen1, minLen2)
}

//方法-:暴力循环
func getsuccessivearrmixLen(arr []int, sum int) int {
	//滑动窗口 1,2,3,4,5
	//i,j
	n := len(arr)
	i := 0
	reslen := n + 1
	for i < n {
		wd_sum := arr[i]
		//单元素相同,直接返回
		if wd_sum == sum {
			reslen = 1
			break
		}
		j := i + 1
		for j < n && wd_sum+arr[j] < sum {
			wd_sum += arr[j]
			j++
		}
		//找出前n项和小于sum的位置，下一个位置要么大于sum，要么等于sum
		if j < n && wd_sum+arr[j] >= sum {
			if j-i < reslen {
				reslen = j - i + 1
			}
		}
		i++
	}
	//没有符合的
	if reslen > n {
		reslen = 0
	}
	return reslen
}

/**
 * @description: 这个很好的包容了临界情况，比如单元素等于s
 * @param {[]int} arr
 * @param {int} s
 * @return {*}
 */
func getsuccessivearrmixLenV2(arr []int, s int) int {
	n := len(arr)
	subLen := n + 1
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum >= s {
				if j-i+1 < subLen {
					subLen = j - i + 1
					break
				}
			}
		}
	}
	if subLen == n+1 {
		subLen = 0
	}
	return subLen
}

//滑动窗口 O(n)
func minSubArrayLen(arr []int, total int) int {
	i := 0
	sum := 0
	n := len(arr)
	minSubLength := n + 1
	for j := 0; j < n; j++ {
		sum += arr[j]
		for sum >= total {
			if j-i+1 < minSubLength {
				minSubLength = j - i + 1
			}
			//调整起始位置
			sum -= arr[i]
			i++
		}
	}
	if minSubLength == n+1 {
		minSubLength = 0
	}
	return minSubLength
}
