package main

import "fmt"

//27. 移除元素
//https://leetcode-cn.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	length := len(nums)
	i, j := 0, 0
	for ;j < length;j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	result := removeElement(nums, 2)
	fmt.Println("result = ", result)
}
