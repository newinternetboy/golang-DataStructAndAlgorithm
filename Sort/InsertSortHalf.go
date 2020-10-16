package Sort

//折半插入排序，关键字插入的序列已经有序，可以通过折半查找需要插入的位置，并将插入位置后边到关键字原位置的元素后移，然后插入关键字到找到的合适位置
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
		//最终tmp插入的位置在low
		j := i
		for ; j > low; j-- {
			params[j] = params[j-1]
		}

		params[low] = tmp
	}
	return params
}
