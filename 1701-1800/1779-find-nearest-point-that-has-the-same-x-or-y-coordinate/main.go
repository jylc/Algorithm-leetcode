package main

import "math"

// https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/description/
// 1779. 找到最近的有相同 X 或 Y 坐标的点

func nearestValidPoint(x int, y int, points [][]int) int {
	distance := math.MaxInt64
	index := -1
	for i, point := range points {
		if point[0] == x || point[1] == y {
			if distance > abs(x-point[0])+abs(y-point[1]) {
				distance = abs(x-point[0]) + abs(y-point[1])
				index = i
			}
		}
	}
	return index
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

}
