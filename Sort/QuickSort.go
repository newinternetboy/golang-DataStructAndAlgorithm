package Sort

func QuickSort(params []int) []int {
	quickSort(params, 0, len(params)-1)
	return params
}

func quickSort(params []int, low, high int) {
	if low < high {
		tmp_index := partionSort(params, low, high)
		partionSort(params, low, tmp_index-1)
		partionSort(params, tmp_index+1, high)
	}
}

func partionSort(params []int, low, high int) int {
	//基准存储是因为先左找小于基准值，并复制给低位，这样就会覆盖之前的基准值
	tmp := params[low]
	for low < high {
		for low < high && params[high] >= tmp {
			high--
		}
		//将高位小于基准值的移动到低位
		params[low] = params[high]
		for low < high && params[low] <= tmp {
			low++
		}
		//将低位大于基准值的移动到高位
		params[high] = params[low]
	}
	params[low] = tmp
	return low
}
