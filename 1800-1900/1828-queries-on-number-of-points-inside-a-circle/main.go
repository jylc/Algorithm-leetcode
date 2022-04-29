package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/queries-on-number-of-points-inside-a-circle/
//1828. 统计一个圆中点的数目
func isInside(point []int, circle []int) bool {
	x, y := point[0], point[1]
	x1, y1, r1 := circle[0], circle[1], circle[2]
	if (x-x1)*(x-x1)+(y-y1)*(y-y1) <= r1*r1 {
		return true
	}
	return false
}
func countPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, query := range queries {
		for _, point := range points {
			if isInside(point, query) {
				ans[i]++
			}
		}
	}
	return ans
}
func main() {
	points := [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}
	queries := [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}
	ans := countPoints(points, queries)
	fmt.Println(ans)
}
