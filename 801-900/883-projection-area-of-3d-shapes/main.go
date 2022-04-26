package main

import "fmt"

//https://leetcode-cn.com/problems/projection-area-of-3d-shapes/
//883. 三维形体投影面积
func max(a, b int) (ans int) {
	if a > b {
		ans = a
	} else {
		ans = b
	}
	return
}
func projectionArea(grid [][]int) int {
	var xy, xz, yz int
	for i, row := range grid {
		xzHeight, yzHeight := 0, 0
		for j, v := range row {
			if v > 0 {
				xy++
			}
			yzHeight = max(yzHeight, grid[j][i])
			xzHeight = max(xzHeight, v)
		}
		yz += yzHeight
		xz += xzHeight
	}
	return xz + xy + yz
}
func main() {
	//grid := [][]int{{1, 2}, {3, 4}}
	//grid := [][]int{{2}}
	grid := [][]int{{1, 0}, {0, 2}}
	ans := projectionArea(grid)
	fmt.Println(ans)
}
