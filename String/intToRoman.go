/*
 * @Description:罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
 * @Author: Caoxiang
 * @Date: 2021-08-17 21:51:14
 * @LastEditors: Caoxiang
*/
package String

import (
	"strings"
)

type romanReflectSlice struct {
	Base        int
	RomanString string
}

func IntToRoman(num int) string {
	res := []string{}
	romanReflectSlice := []romanReflectSlice{{Base: 1000, RomanString: "M"}, {Base: 900, RomanString: "CM"}, {Base: 500, RomanString: "D"}, {Base: 400, RomanString: "CD"}, {Base: 100, RomanString: "C"}, {Base: 90, RomanString: "XC"}, {Base: 50, RomanString: "L"}, {Base: 40, RomanString: "XL"}, {Base: 10, RomanString: "X"}, {Base: 9, RomanString: "IX"}, {Base: 5, RomanString: "V"}, {Base: 4, RomanString: "IV"}, {Base: 1, RomanString: "I"}}
	for _, v := range romanReflectSlice {
		baseNum := num / v.Base
		if baseNum != 0 {
			num = num - v.Base*baseNum
			for i := 1; i <= baseNum; i++ {
				res = append(res, v.RomanString)
			}
		}
	}
	return strings.Join(res, "")
}
