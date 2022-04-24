package main

import "fmt"

//https://leetcode-cn.com/problems/flower-planting-with-no-adjacent/
//1042. 不邻接植花
func gardenNoAdj(n int, paths [][]int) []int {
	var ans = make([]int, n)
	var c = make([][]int, n)
	for _, line := range paths {
		c[line[0]-1] = append(c[line[0]-1], line[1]-1)
		c[line[1]-1] = append(c[line[1]-1], line[0]-1)
	}

	fmt.Println(c)
	for i := 0; i < n; i++ {
		used := make(map[int]struct{})
		for _, point := range c[i] {
			if ans[point] != 0 {
				used[ans[point]] = struct{}{}
			}
		}

		for j := 1; j <= 4; j++ {
			if _, ok := used[j]; !ok {
				ans[i] = j
				break
			}
		}
	}
	return ans
}
func main() {
	n := 5
	//path := [][]int{{1, 2}, {2, 3}, {3, 1}}
	//path := [][]int{{1, 2}, {3, 4}}
	//path := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}
	path := [][]int{{4, 1}, {4, 2}, {4, 3}, {2, 5}, {1, 2}, {1, 5}}
	ans := gardenNoAdj(n, path)
	fmt.Println(ans)
}
