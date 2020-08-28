package main

import (
	"fmt"
	"math"
)

//剑指 Offer 57 - II. 和为s的连续正数序列
//https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
//暴力求解
func findContinuousSequence1(target int) [][]int {
	total := 0
	res := make([][]int, 0)
	for i := 1; i < target/2+1; i++ {
		var tmp []int
		total = 0
		for j := i; j <= target/2+1 && total <= target; j++ {
			tmp = append(tmp, j)
			total += j
			if total == target {
				res = append(res, tmp)
			}
		}
	}
	return res
}

//数学公式法
func findContinuousSequence2(target int) [][]int {
	res := make([][]int, 0)
	for i := int(math.Sqrt(float64(2 * target))); i >= 2; i-- {
		judge := target - i*(i-1)/2
		if judge%i == 0 {
			begin := judge / i
			temp := make([]int, 0)
			for j := 0; j < i; j++ {
				temp = append(temp, begin+j)
			}
			res = append(res, temp)
		}
	}
	return res
}

func main() {
	result := findContinuousSequence1(15)
	fmt.Println("result = ", result)
}
