package Sort

func BubbleSort(params []int, order int) ([]int,int) {
	length := len(params)
	if length <= 1 {
		return params,0
	}

	swap := 0
	//总共遍历次数为数组的长度减一
	for circle := 0; circle < len(params); circle++ {
		//每次数组的遍历都会将最小的或者最大的置顶
		for itemIndex := 0; itemIndex < length-1; itemIndex++ {
			//降序
			if order == 0 && params[itemIndex] < params[itemIndex+1] {
				params[itemIndex], params[itemIndex+1] = params[itemIndex+1], params[itemIndex]
				swap += 1
			}
			//升序
			if order == 1 && params[itemIndex] > params[itemIndex+1] {
				params[itemIndex], params[itemIndex+1] = params[itemIndex+1], params[itemIndex]
				swap += 1
			}
		}
	}

	return params, swap
}
