package main

import (
	"sort"
)

//406. 根据身高重建队列
//https://leetcode-cn.com/problems/queue-reconstruction-by-height/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	ans := make([][]int, len(people))
	for _, person := range people {
		spaces := person[1] + 1
		for i := range ans {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}
func main() {
	arr := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue(arr)
}
