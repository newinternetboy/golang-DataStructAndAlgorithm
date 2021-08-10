/*
 * @Description:
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
 * @Author: Caoxiang
 * @Date: 2021-08-09 16:38:46
 * @LastEditors: Caoxiang
*/
package main

import "fmt"

func main() {
	s := "({)()}"
	b := isValid(s)
	fmt.Println(b)
}

func isValid(s string) bool {
	validBucketPair := make(map[rune]rune, 3)
	validBucketPair['('] = ')'
	validBucketPair['{'] = '}'
	validBucketPair['['] = ']'
	res := []rune{}
	for _, v := range s {
		if len(res) > 0 {
			//取最上面的一位,看和下一位是否匹配
			tv := res[len(res)-1]
			if rb, ok := validBucketPair[tv]; ok {
				if rb == v {
					res = res[:len(res)-1]
				} else {
					res = append(res, v)
				}
			} else {
				res = append(res, v)
			}
		} else {
			res = append(res, v)
		}
	}
	isValid := true
	if len(res) > 0 {
		isValid = false
	}
	return isValid
}
