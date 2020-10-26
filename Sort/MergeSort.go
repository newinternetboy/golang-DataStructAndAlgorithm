/*
 * @Author: your name
 * @Date: 2020-10-14 15:05:19
 * @LastEditTime: 2020-10-26 14:43:40
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /DataStruct/golang-DataStructAndAlgorithm/Sort/MergeSort.go
 */
package Sort

//分治思想
//将数组进行分组,对左右两组进行递归分治，并将分治排好序的数组合并
//时间复杂度:O(nlogn)
func MergeSort(params []int, low, high int) []int {
	if low < high {
		mid := low + (high-low)/2
		MergeSort(params, low, mid)
		MergeSort(params, mid+1, high)
		merge(params, low, mid, high)
	}
	return params
}

func merge(params []int, low, mid, high int) {
	//需要将low -> high 排成有序序列
	i := low
	j := mid + 1
	tmp := []int{}
	for i <= mid && j <= high {
		if params[i] <= params[j] {
			tmp = append(tmp, params[i])
			i++
		} else {
			tmp = append(tmp, params[j])
			j++
		}
	}

	for i <= mid {
		tmp = append(tmp, params[i])
		i++
	}

	for j <= high {
		tmp = append(tmp, params[j])
		j++
	}

	//将分治排序后的临时数组赋值给原数组
	t := 0
	for low <= high {
		params[low] = tmp[t]
		low++
		t++
	}
}
