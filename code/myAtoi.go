package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INT_MAX = int32(^uint32(0) >> 1)
const INT64_MAX = int64(^uint64(0) >> 1)
const INT_MIN = ^INT_MAX
const INT64_MIN = ^INT64_MAX

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	strNew := ""
	if strings.HasPrefix(str, "-") {
		strNew += "-"
	}
	if strings.HasPrefix(str, "+") {
		strNew += "+"
	}
	isAdd := false
	for k, v := range str {
		if v < '0' || v > '9' {
			if k != 0 || string(v) != strNew {
				break
			} else {
				continue
			}
		}
		if v == '0' && !isAdd {
			continue
		}
		isAdd = true
		strNew = strNew + string(v)
	}
	fmt.Println(strNew)
	if strNew == "" || strNew == "-" || strNew == "+" {
		return 0
	}
	strmax := strconv.FormatInt(int64(INT64_MAX), 10)
	fmt.Println(INT64_MAX, strmax, len(strmax))
	if len(strNew) >= len(strmax)-1 {
		strNew = strNew[:len(strmax)-1]
	}

	res, err := strconv.ParseInt(strNew, 10, 64)
	fmt.Println(res, err)
	if err != nil {
		return 0
	}

	fmt.Println(res, "\n", INT_MIN)
	if res > int64(INT_MAX) {
		res = int64(INT_MAX)
	}
	if res < int64(INT_MIN) {
		res = int64(INT_MIN)
	}
	return int(res)
}

func main() {
	str := os.Args[1]
	fmt.Println(myAtoi(str))
}
