package Sort

//插入排序
//算法思想:插入关键字和前面已经有序的列表比较，将列表中大于关键字的元素后移，然后将关键字插入到最后移动元素的前面
//时间复杂度:
func InsertSort(params []int) []int {
	n := len(params)
	var tmp int
	for i := 1; i < n; i++ {
		tmp = params[i]
		j := i - 1
		for j >= 0 && tmp < params[j] {
			params[j+1] = params[j]
			j--
		}
		params[j+1] = tmp
	}
	return params
}
