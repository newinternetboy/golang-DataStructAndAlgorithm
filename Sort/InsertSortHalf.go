package Sort

import "fmt"

//
func InsertSortHalf(params []int) []int {
	n := len(params)
	if n <= 1 {
		return params
	}
	for i := 1; i < n; i++ {
		tmp := params[i]
		low := 0
		high := i - 1
		for low <= high {
			mid := low + (high-low)/2
			if tmp < params[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Println(low, high)
		//最终tmp插入的位置在low
		j := i
		for ; j > low; j-- {
			params[j] = params[j-1]
		}

		params[low] = tmp
	}
	return params
}
