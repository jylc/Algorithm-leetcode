package main

import (
	"fmt"
	"sort"
)

//https://leetcode.cn/contest/weekly-contest-293/problems/maximum-consecutive-floors-without-special-floors/
//6064. 不含特殊楼层的最大连续楼层数
func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	max := 0
	for i := 0; i < len(special); i++ {
		dis := 0
		if i == 0 {
			dis = special[i] - bottom
		} else {
			dis = special[i] - special[i-1] - 1
		}
		if max < dis {
			max = dis
		}
	}
	if top-special[len(special)-1] > max {
		max = top - special[len(special)-1]
	}

	return max
}
func main() {
	bottom := 2
	top := 9
	special := []int{4, 6}
	ans := maxConsecutive(bottom, top, special)
	fmt.Println(ans)
}
