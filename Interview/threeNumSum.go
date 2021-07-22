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
	res := threeNumSum()
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
