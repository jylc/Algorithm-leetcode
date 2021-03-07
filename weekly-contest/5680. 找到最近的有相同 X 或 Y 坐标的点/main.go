package main

import (
	"fmt"
	"math"
)

//5680. 找到最近的有相同 X 或 Y 坐标的点
//https://leetcode-cn.com/contest/biweekly-contest-47/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/

func nearestValidPoint(x int, y int, points [][]int) int {
	m := len(points)
	min := math.MaxFloat64
	ans := -1
	for i := 0; i < m; i++ {
		for j := 0; j < 2; j++ {
			if x == points[i][0] || y == points[i][1] {
				dis := math.Abs(float64(points[i][0]-x)) + math.Abs(float64(points[i][1]-y))
				if dis < min {
					ans = i
					min = dis
				}
			}
		}
	}

	return ans
}

func main() {
	x := 3
	y := 4
	points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	res := nearestValidPoint(x, y, points)
	fmt.Print("res = ", res)
}
