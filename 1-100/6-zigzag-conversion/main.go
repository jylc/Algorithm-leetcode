package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/zigzag-conversion/
//6. Z 字形变换

func convert(s string, numRows int) (ans string) {
	zMap := make(map[int][]rune, numRows)
	if numRows == 1 {
		return s
	}

	var step int = (numRows << 1) - 2
	for i, v := range s {
		if t := i % step; t < numRows {
			zMap[t] = append(zMap[t], rune(v))
		} else {
			zMap[step-t] = append(zMap[step-t], rune(v))
		}
	}

	for i := 0; i < numRows; i++ {
		ans += string(zMap[i])
	}
	return
}
func main() {
	ans := convert("A", 1)
	fmt.Println("\n\nans=\n" + ans)
}
