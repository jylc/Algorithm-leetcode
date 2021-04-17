package main

import "math"

//220. 存在重复元素 III
//https://leetcode-cn.com/problems/contains-duplicate-iii/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= i+k && j < n; j++ {
			if int(math.Abs(float64(nums[j]-nums[i]))) <= t {
				return true
			}
		}
	}
	return false
}

type node struct {
	ch       [2]*node
	priority int
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

//使用桶方法模拟窗口
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	var getId func(x int) int
	getId = func(x int) int {
		//向下取整
		if x >= 0 {
			return x / (t + 1)
		}
		return (x+1)/t - 1
	}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		id := getId(x)
		if _, ok := m[id]; ok {
			return true
		}
		if y, ok := m[id+1]; ok && abs(x-y) <= t {
			return true
		}
		if y, ok := m[id-1]; ok && abs(x-y) <= t {
			return true
		}
		m[id] = x
		if i >= k {
			delete(m, getId(nums[i-k]))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
