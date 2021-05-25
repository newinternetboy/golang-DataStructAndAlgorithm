/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-21 16:03:09
 * @LastEditors: Caoxiang
 */
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

func SelectSortV2(params []int) []int {
	n := len(params)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if params[j] < params[k] {
				k = j
			}
		}
		params[i], params[k] = params[k], params[i]
	}
	return params
}

//后端服务开发
//数据量很大
//高并发分布式有一定经验，有认知
//kafka,redis,zookeeper.为什么快，怎么做到的，批量读写？
//单机的特点。
//服务的压测
