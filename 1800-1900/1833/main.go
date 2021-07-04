package main

import (
	"fmt"
	"sort"
)

//1833. 雪糕的最大数量
//https://leetcode-cn.com/problems/maximum-ice-cream-bars/
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	sum := 0
	for i, cost := range costs {
		sum += cost
		if sum > coins {
			return i
		}
	}
	return len(costs)
}
func main() {
	costs := []int{1, 6, 3, 1, 2, 5}
	coins := 20
	ans := maxIceCream(costs, coins)
	fmt.Println(ans)
}
