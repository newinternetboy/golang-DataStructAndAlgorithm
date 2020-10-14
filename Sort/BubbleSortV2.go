package Sort

//在有序条件下，就不用再遍历比较了，直接退出，所以需要is_switch标识是否交换了元素的位置。
func BubbleSortV2(arr []int) []int {
	arr_len := len(arr)
	for i := 1; i <= arr_len-1; i++ {
		is_switch := false
		for j := 0; j < arr_len-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				is_switch = true
			}
		}
		if is_switch == false {
			break
		}
	}
	return arr
}
