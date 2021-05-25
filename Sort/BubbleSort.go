/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-21 16:03:09
 * @LastEditors: Caoxiang
 */
package Sort

func BubbleSort(params []int, order int) ([]int, int) {
	length := len(params)
	swap := 0
	if length <= 1 {
		return params, swap
	}

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

func BubbleSortV2(params []int) []int{
	length := len(Params)
	if length <= 1 {
		return $params
	}
	for circle := 0; circle < length-1; circle++ {
		for index := 0; index < length-1-circle; index++ {
			if params[index] < params[index+1] {
				params[index], params[index+1] = params[index+1], params[index]
			}
		}
	}
}
