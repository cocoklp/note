package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	/*
		结束： n=0||n=1
		处理： 得到后面的string，找连续数字，拼接
		返回： 本级字符串
	*/
	if n == 0 {
		return ""
	} else if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	}
	str := countAndSay(n - 1)
	i, k := 1, 0
	count := 1
	result := ""
	for i = 1; i < len(str); i++ {
		if str[k] == str[i] {
			count++
		} else {
			result = fmt.Sprintf("%s%d%s", result, count, string(str[k]))
			count = 1
		}
		k++
	}
	result = fmt.Sprintf("%s%d%s", result, count, string(str[k]))

	return result
}
