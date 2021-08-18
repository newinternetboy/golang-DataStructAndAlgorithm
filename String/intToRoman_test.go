/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-08-18 10:25:27
 * @LastEditors: Caoxiang
 */
package String

import "testing"

func TestIntToRoman(t *testing.T) {
	num := 1994
	res := IntToRoman(num)
	t.Log(res)
}
