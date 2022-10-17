package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-cost-of-buying-candies-with-discount/
// 2144. 打折购买糖果的最小开销

func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})
	ans := 0
	for i := 0; i < len(cost); i++ {
		if i%3 != 2 {
			ans += cost[i]
		}
	}
	return ans
}

func main() {
	cost := []int{6, 5, 7, 9, 2, 2}
	ans := minimumCost(cost)
	fmt.Println(ans)
}
