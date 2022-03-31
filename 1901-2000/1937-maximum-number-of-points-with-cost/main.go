package main

import "math"

//https://leetcode-cn.com/problems/maximum-number-of-points-with-cost/
//1937. 扣分后的最大得分

func maxPoints(points [][]int) int64 {
	ans := 0
	m := len(points[0])
	f := make([][2]int, m)
	sufMax := make([]int, m)
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i, row := range points {
		if i == 0 {
			for j, v := range row {
				ans = max(ans, v)
				f[j][0] = v + j
				f[j][1] = v - j
			}
		} else {
			preMax := math.MinInt32
			for j, v := range row {
				preMax = max(preMax, f[j][0])
				res := max(v-j+preMax, v+j+sufMax[j])
				ans = max(ans, res)
				f[j][0] = res + j
				f[j][1] = res - j
			}
		}
		sufMax[m-1] = f[m-1][1]
		for j := m - 2; j >= 0; j-- {
			sufMax[j] = max(sufMax[j+1], f[j][1])
		}
	}
	return int64(ans)
}

func main() {

}
