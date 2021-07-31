package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/excel-sheet-column-number/
//171. Excel 表列序号

func titleToNumber(columnTitle string) int {
	columns := []rune(columnTitle)
	ans := 0
	for i, column := range columns {
		n := len(columns) - i - 1
		m := column - 'A' + 1
		ans += int(math.Pow(26, float64(n))) * int(m)
	}
	return ans
}
func main() {
	ans := titleToNumber("ZY")
	fmt.Println(ans)
}
