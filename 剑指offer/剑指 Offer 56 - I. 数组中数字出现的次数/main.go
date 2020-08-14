package main

import "fmt"

//剑指 Offer 56 - I. 数组中数字出现的次数
//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

func singleNumbers(nums []int) []int {
	var a int
	for i := range nums{
		a ^= nums[i]
	}
	mask := a & (-a)
	res := make([]int,2)
	for _,v :=range nums{
		if (v & mask) == 0{
			res[0] ^= v
		}else{
			res[1] ^= v
		}
	}
	return res
}


func main() {
	nums:=[]int{1,2,10,4,1,4,3,3}
	result := singleNumbers(nums)
	fmt.Println("result = ",result)
}
