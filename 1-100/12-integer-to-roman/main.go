package main

import "fmt"

//https://leetcode.cn/problems/integer-to-roman/
//12. 整数转罗马数字
var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, symbol := range valueSymbols {
		for num >= symbol.value {
			num -= symbol.value
			roman = append(roman, symbol.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}
func main() {
	num := 4
	ans := intToRoman(num)
	fmt.Println(ans)
}
