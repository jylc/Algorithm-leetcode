package main

import "fmt"

//1446. 连续字符
//https://leetcode-cn.com/problems/consecutive-characters/
func maxPower(s string) int {
	data := []byte(s)
	l := len(data)
	count := 0
	result := 0
	start := data[0]
	for i := 1; i < l; i++ {
		if start == data[i] {
			count++
		} else {
			start = data[i]
			count = 0
		}
		if count > result {
			result = count
		}
	}
	return result + 1
}

func main() {
	result := maxPower("tourist")
	fmt.Println("result = ", result)
}
