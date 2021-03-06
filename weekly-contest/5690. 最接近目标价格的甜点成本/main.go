package main

import (
	"fmt"
	"math"
)

//5690. 最接近目标价格的甜点成本
//https://leetcode-cn.com/contest/weekly-contest-230/problems/closest-dessert-cost/

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	res := math.MaxInt64
	m, n := len(baseCosts), len(toppingCosts)
	for i := 0; i < m; i++ {
		s := baseCosts[i]
		for j := 0; j < 1<<(n*2); j++ {
			r := s
			flag := false
			for k := 0; k < n; k++ {
				t := (j >> (k * 2)) & 3
				if t == 3 {
					flag = true
					break
				}
				r += toppingCosts[k] * t
			}
			if flag {
				continue
			}
			if math.Abs(float64(res-target)) > math.Abs(float64(r-target)) ||
				math.Abs(float64(res-target)) == math.Abs(float64(r-target)) && r < res {
				res = r
			}
		}
	}
	return res
}
func main() {
	baseCosts := []int{1, 7}
	toppingCosts := []int{3, 4}
	target := 10
	cost := closestCost(baseCosts, toppingCosts, target)
	fmt.Print("cost = ", cost)
}
