package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/reduction-operations-to-make-the-array-elements-equal
// 1887. 使数组元素相等的减少操作次数

func reductionOperations(nums []int) int {
	numMap := make(map[int]int)
	numArr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; !ok {
			numArr = append(numArr, nums[i])
		}
		numMap[nums[i]]++
	}
	sort.Ints(numArr)
	index := len(numArr) - 1
	count := 0
	for index > 0 {
		count += index * numMap[numArr[index]]
		index--
	}
	return count
}

func main() {
	nums := []int{7, 9, 10, 8, 6, 4, 1, 5, 2, 3}
	ans := reductionOperations(nums)
	fmt.Println(ans)
}
