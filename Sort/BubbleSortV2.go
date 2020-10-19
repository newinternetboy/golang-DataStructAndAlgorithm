package Sort

//在有序条件下，就不用再遍历比较了，直接退出，所以需要is_switch标识是否交换了元素的位置。
func BubbleSortV2(arr []int) []int {
	arr_len := len(arr)
	for i := arr_len - 1; i >= 1; i-- {
		flag := false
		for j := 1; j <= i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return arr
}
