package main

import (
	"fmt"
	"math"
)

//https://leetcode.cn/problems/largest-triangle-area/
//812. 最大三角形面积
func isTriangle(p1, p2, p3 []int) (ans bool, a, b, c float64) {
	distance := func(p1, p2 []int) float64 {
		return math.Sqrt(float64((p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])))
	}
	a, b, c = distance(p1, p2), distance(p1, p3), distance(p2, p3)
	if a+b > c && a+c > b && b+c > a {
		ans = true
		return
	}
	ans = false
	return
}

func largestTriangleArea(points [][]int) float64 {
	max := 0.0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				if ans, a, b, c := isTriangle(points[i], points[j], points[k]); ans {
					total := (a + b + c) / 2
					area := math.Sqrt(total * (total - a) * (total - b) * (total - c))
					if max < area {
						max = area
					}
				}
			}
		}
	}
	return max
}
func main() {
	points := [][]int{
		{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	area := largestTriangleArea(points)
	fmt.Println(area)
}
