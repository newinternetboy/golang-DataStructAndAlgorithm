package Sort

//希尔排序
//原排序序列按照增量进行分组，再对每组执行插入排序，其中增量逐级减少，直到为1
//希尔增量一般取序列长度的一半
func ShellSort(params []int) []int {
	shellSort(params)
	return params
}

//A1-A2-A3
//B1-B2-B3
//C1-C2-C3
//shellSort代码是在一个步长的限制内，进行插入排序
//每组比较A1-A2-A3
func shellSort(params []int) {
	n := len(params)
	for dk := n / 2; dk > 0; dk = dk / 2 {
		//第一组，遍历次数为步长数
		for i := 0; i < dk; i++ {
			for j := i + dk; j < n; j = j + dk {
				//要插入的值
				tmp := params[j]
				//待插入值得前一步长位
				k := j - dk
				//已经有序
				for k >= 0 && params[k] > tmp {
					params[k+dk] = params[k]
					k = k - dk
				}
				params[k+dk] = tmp
			}
		}
	}
}

//还可在步长分组间进行插入排序
//比较方式A1-B1-C1
