/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-05-25 18:01:23
 * @LastEditors: Caoxiang
 */
package String

import (
	"testing"
)

func TestLeftRotateOneString(t *testing.T) {
	s := "abcd"
	m := 2
	for i := 0; i < m; i++ {
		s = LeftRotateOneString(s, len(s))
	}
	if s != "cdab" {
		t.Fatal()
	}
}
