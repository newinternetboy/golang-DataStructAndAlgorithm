/*
 * @Description:三数之和
 * @Author: Caoxiang
 * @Date: 2021-07-22 14:23:50
 * @LastEditors: Caoxiang
 */
package main

import (
	"fmt"
	"sort"
)

type Sli struct {
	sli []int
}

func (s *Sli) Swap(i, j int) {
	s.sli[i], s.sli[j] = s.sli[j], s.sli[i]
}

func (s *Sli) Less(i, j int) bool {
	return s.sli[i] < s.sli[j]
}

func (s *Sli) Len() int {
	return len(s.sli)
}

func main() {
	// res := threeNumSum()
	num := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(num)
	fmt.Println(res)
}

func threeNumSum() [][3]int {
	s := &Sli{sli: []int{-1, 0, 1, 2, -1, -4}}
	sort.Sort(s)
	sortedSli := s.sli
	res := [][3]int{}
	fmt.Println(res)
	//for i
	//right left
	//right + left > i, left左移
	//right + left < i, right右移
	//right + left = i, 判断right，left
	length := len(sortedSli)
	if length < 3 {
		return res
	}
	for i, v := range sortedSli {
		left := i + 1
		right := length - 1
		for left < right {
			if sortedSli[right]+sortedSli[left]+v == 0 {
				res = append(res, [3]int{v, sortedSli[right], sortedSli[left]})
				for right < left && sortedSli[right] == sortedSli[right-1] {
					right--
				}

				for right < left && sortedSli[left] == sortedSli[left+1] {
					left++
				}
				right--
				left++
			} else if sortedSli[right]+sortedSli[left]+v > 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	length := len(nums)
	if length < 3 {
		return res
	}
	for i, v := range nums {
		left := i + 1
		right := length - 1
		//之后和前面的重复了，就跳过
		if i > 0 && nums[i-1] == v {
			continue
		}
		for left < right {
			if nums[right]+nums[left]+v == 0 {
				res = append(res, []int{v, nums[left], nums[right]})
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			} else if nums[right]+nums[left]+v > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
