package main

import "fmt"

//26. 删除排序数组中的重复项
//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length==1{
		return 1
	}
	i, j := 0, 1
	for j < length {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i+1
}
func main() {
	nums:=[]int{0,0,1,1,1,2,2,3,3,4}
	length := removeDuplicates(nums)
	fmt.Println("result = ",length)
	fmt.Println(nums)
}
