package main

import "fmt"

//5698. 构成特定和需要添加的最少元素
//https://leetcode-cn.com/contest/weekly-contest-231/problems/minimum-elements-to-add-to-form-a-given-sum/

func minElements(nums []int, limit int, goal int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	l := goal - sum
	cnt := 0
	a := 0
	for {
		if l == 0 {
			return cnt
		}
		if l >= -limit && l <= limit {
			cnt++
			return cnt
		}
		if l >= 0 {
			a = l
		} else {
			a = -l
		}
		m := 0
		if l < -limit || l > limit {
			m = a / limit
			l = l % limit
		}
		cnt += m
	}
}

func main() {
	nums := []int{-1, 0, 1, 1, 1}
	limit := 1
	goal := 771843707
	res := minElements(nums, limit, goal)
	fmt.Print("res = ", res)
}
