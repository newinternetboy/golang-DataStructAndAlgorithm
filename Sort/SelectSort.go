package Sort

//简单选择排序
//选择方式:顺序遍历找到最小值
//每一次排序将最小值和遍历序列起始值进行交换
func SelectSort(params []int) []int {
	n := len(params)
	for i := 0; i < n; i++ {
		//从i+1遍历到结束
		k := i
		for j := i + 1; j < n; j++ {
			if params[j] < params[k] {
				k = j
			}
		}
		//交换
		params[i], params[k] = params[k], params[i]
	}
	return params
}
