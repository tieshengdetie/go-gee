package main

import (
	"fmt"
	"strings"
)

/*判断两个给定的字符串排列后是否一致
问题描述

给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。
保证两串的长度都小于等于5000。

解题思路

首先要保证字符串长度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存在即可。

源码参考*/

func main() {
	s1 := "abades"
	s2 := "esdbaa"
	res := isStringSameAfterSort(s1, s2)
	fmt.Print(res)
}

func isStringSameAfterSort(s1, s2 string) bool {

	rs1, rs2 := []rune(s1), []rune(s2)
	ls1, ls2 := len(rs1), len(rs2)
	if ls1 > 5000 || ls2 > 5000 || ls1 != ls2 {
		return false
	}
	// 只要判断字符串中每一个字符出现的次数相等，就可以重新排列好相同
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}

	}
	return true
}
