package main

import (
	"fmt"
)

func converti(s string, numRows int) string {
	if len(s) <= 2 || numRows < 2 {
		return s
	}
	strSlice := make([][]byte, 0)
	cnt := 2*numRows - 2
	for k := 0; k < len(s); {
		tmp := make([]byte, numRows)
		if k%cnt == 0 {
			end := k + numRows
			if end >= len(s) {
				end = len(s)
			}
			tmp = []byte(s[k:end])
			k = k + numRows
		} else {
			idx := cnt - k%cnt
			tmp[idx] = s[k]
			k++
		}
		strSlice = append(strSlice, tmp)
	}
	rst := make([]byte, 0, len(s))
	for k := 0; k < numRows; k++ {
		for _, v := range strSlice {
			if len(v) <= k {
				continue
			}
			rst = append(rst, v[k])
		}
	}
	return string(rst)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	runes := []rune(s)
	n := len(runes)
	runesOut := make([]rune, n)
	j := 0
	for row := 0; row < numRows; row++ {
		k := row
		skip := 2 * (numRows - row - 1)
		for j < n && k < n {
			fmt.Println(row, k, skip)
			if skip > 0 {
				runesOut[j] = runes[k]
				k = k + skip
				j++
			}
			skip = 2*(numRows-1) - skip
		}
	}
	return string(runesOut)
}
func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

/*
首先将字符按照Z字型排列起来，之后按照第一行到最后一行在组成串:AHNBGIMOCFJLPDKQ。我们可以看到第一行和最后一行的字母的下标是呈现出一个等差数列,根据关系观察很容易就可以发现这等差数列的差值是: 2 * numRows -2

实际上我们发现除了第一行和最后一行，其它行之中依旧存在等差数列，如在第二行中(B,I,O)同样其标也是呈现这样的等差关系，那么实际其实要解决这道题就只剩下最后一个问题了就是解决两个等差数中间的那个数的坐标（如BI之间G的坐标)。我们很容易就可以发现它的左边和其前一个数的间隔为 2 * rowRows-2-2 * row
*/
func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	res := make([]byte, len(s))
	inter1 := 2*numRows - 2
	m := 0
	for i := 0; i < numRows; i++ {
		inter2 := 2 * i
		for j := i; j < len(s); j += inter2 {
			res[m] = s[j]
			m++
			if i == 0 || i == numRows-1 {
				inter2 = inter1
			} else {
				inter2 = inter1 - inter2
			}
		}
	}
	return string(res)
}
