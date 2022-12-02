package main

import "fmt"

// https://leetcode.cn/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
// 1769. 移动所有球到每个盒子所需的最小操作数

func minOperations(boxes string) []int {
	arr := make([]int, 0)
	for index, box := range boxes {
		if box-'1' == 0 {
			arr = append(arr, index)
		}
	}
	result := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		for _, i3 := range arr {
			result[i] += abs(i3 - i)
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	boxes := "001011"
	result := minOperations(boxes)
	fmt.Println(result)
}
