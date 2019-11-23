package Sort

func BubbleSort(params []int, order int) []int {
	length := len(params)
	if length <= 1 {
		return params
	}

	swap := 1
	for swap > 0 {
		//只要交换了就要重新做一次全局位置交换
		swap = 0
		for itemIndex := 0; itemIndex < length-1; itemIndex++ {
			//降序
			if order ==0 && params[itemIndex] < params[itemIndex+1] {
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
	return params
}
