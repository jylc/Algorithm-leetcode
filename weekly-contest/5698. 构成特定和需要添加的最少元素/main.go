package main

import (
	"fmt"
	"math"
)

//5698. 构成特定和需要添加的最少元素
//https://leetcode-cn.com/contest/weekly-contest-231/problems/minimum-elements-to-add-to-form-a-given-sum/

/*
思路：添加值/limit=次数，写得复杂了-_-||
*/
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

/*
思路：目标值-已有值=添加值，一般添加值超过limit时添加abs(limit)的值，所以求要添加abs(limit)的次数；
次数=（添加值/limit）向上取整 等同于（（添加值+limit-1）/limit）向下取整
*/
func minElements1(nums []int, limit int, goal int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return (int(math.Abs(float64(sum-goal))) + limit - 1) / limit
}

func main() {
	nums := []int{-1, 0, 1, 1, 1}
	limit := 1
	goal := 771843707
	res := minElements(nums, limit, goal)
	fmt.Print("res = ", res)
}
