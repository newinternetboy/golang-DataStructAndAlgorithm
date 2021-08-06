/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-05-25 17:19:17
 * @LastEditors: Caoxiang
 */
package String

//字符串前m位移动到字符串尾部,要求时间复杂度O(n),空间复杂度O(1)
//方法一:暴力法
func LeftRotateOneString(s string, n int) string {
	sr := []rune(s)
	srf := sr[0]
	for i := 0; i < len(sr)-1; i++ {
		sr[i] = sr[i+1]
	}
	sr[len(sr)-1] = srf
	return string(sr)
}
