package main

import "fmt"

//剑指 Offer 03. 数组中重复的数字
//https: //leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

func findRepeatNumber(nums []int) int {
	m:=make(map[int]int)

	for _,value:=range nums{
		if _,ok:=m[value];ok{
			return value
		}else {
			m[value]=1
		}
	}
	return -1
}

func main() {
	arr:=[]int{2, 3, 1, 0, 2, 5, 3}
	result := findRepeatNumber(arr)
	fmt.Println("result:",result)
}
