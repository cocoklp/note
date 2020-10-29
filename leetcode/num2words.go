package main

import (
	"flag"
	"fmt"
)

var (
	input = flag.Int("num", 0, "num")
)
var (
	BILLION  = 1000000000
	THOUSAND = 1000
	HUNDERD  = 100
	TEN      = 10
)
var thousandToUnit = map[int]string{
	1000000000: "Billion",
	1000000:    "Million",
	1000:       "Thousand",
}

var tenToUnit = map[int]string{
	10:  "Ten",
	11:  "Eleven",
	12:  "Twelve",
	13:  "Thirteen",
	14:  "Fourteen",
	15:  "Fifteen",
	16:  "Sixteen",
	17:  "Seventeen",
	18:  "Eighteen",
	19:  "Nineteen",
	20:  "Twenty",
	30:  "Thirty",
	40:  "Forty",
	50:  "Fifty",
	60:  "Sixty",
	70:  "Seventy",
	80:  "Eighty",
	90:  "Ninty",
	100: "Hundred",
}

var oneToUnit = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

func numberToWords(num int) string {
	divider := BILLION
	rst := ""
	for divider != 0 {
		fmt.Println(num)
		high := num / divider
		if high/HUNDERD != 0 {
			rst = rst + " " + oneToUnit[high/HUNDERD] + " Hunderd"
		}
		tmp := high % HUNDERD
		for tmp != 0 {
			if tmp/10 == 1 {
				rst = rst + " " + tenToUnit[tmp]
				break
			} else if tmp/10 == 0 {
				rst = rst + " " + oneToUnit[tmp]
				break
			} else {
				rst = rst + " " + tenToUnit[tmp/10*10]
			}
			tmp = tmp % 10
		}
		if high != 0 {
			rst = rst + " " + thousandToUnit[divider]
		}
		num = num % divider
		divider = divider / THOUSAND
	}
	fmt.Println(rst)
	return rst
}
func main() {
	flag.Parse()
	numberToWords(*input)
}
