package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/count-number-of-rectangles-containing-each-point/
//6043. 统计包含每个点的矩形数目
func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][1] > rectangles[j][1]
	})
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] > points[j][1] })
	ans := make([]int, len(points))
	var xs []int
	i, n := 0, len(rectangles)
	for _, point := range points {
		start := i
		for ; i < n && rectangles[i][1] >= point[1]; i++ {
			xs = append(xs, rectangles[i][0])
		}
		if start < i {
			sort.Ints(xs)
		}
		ans[point[2]] = i - sort.SearchInts(xs, point[0])
	}
	return ans
}
func main() {
	rectangles := [][]int{{1, 2}, {2, 3}, {2, 5}}
	points := [][]int{{2, 1}, {1, 4}}
	ans := countRectangles(rectangles, points)
	fmt.Println(ans)
}
