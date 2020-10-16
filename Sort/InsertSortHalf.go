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

func InsertSortHalfV2(params []int) []int {
	n := len(params)
	if n <= 1 {
		return params
	}
	for i := 1; i < n; i++ {
		low := 0
		tmp := params[i]
		high := i - 1
		for low <= high {
			mid := low + (high-low)/2
			if tmp < params[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		//low就是插入的位置
		//插入位置之后的元素后移
		j := i
		for ; j > low; j-- {
			params[j] = params[j-1]
		}

		params[low] = tmp
	}
	return params
}

func InserSortHalfV3(params []int) []int {
	n := len(params)
	if n <= 1 {
		return params
	}
	for i := 1; i < n; i++ {
		//临时存储
		tmp := params[i]
		//寻找插入位置
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

		//插入位置后面到插入关键字原位置前的元素后移
		j := i
		for ; j > low; j-- {
			params[j] = params[j-1]
		}

		//插入
		params[low] = tmp
	}
	return params
}
