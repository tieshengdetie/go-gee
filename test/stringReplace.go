package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

/*问题描述

请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，
并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。 给定一个string为原始的串，返回替换后的string。

解题思路

两个问题，第一个是只能是英文字母，第二个是替换空格。*/

func main() {

	s := "abs d d"
	var isLetter = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(s)
	fmt.Print(isLetter)
}
func stringReplace(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}

	return strings.Replace(s, " ", "%20", -1), true
}

func stringReplace2(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false

	}
	var isLetter = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(s)
	if !isLetter {
		return s, false
	}
	return strings.Replace(s, " ", "%20", -1), true
}
