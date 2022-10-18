package main

import "fmt"

// https://leetcode.cn/problems/possible-bipartition/
// 886. 可能的二分法

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, dislike := range dislikes {
		x, y := dislike[0]-1, dislike[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	color := make([]int, n)
	var dfs func(int, int) bool
	dfs = func(x int, c int) bool {
		color[x] = c
		for _, y := range g[x] {
			if color[y] == c || color[y] == 0 && !dfs(y, 3^c) {
				return false
			}
		}
		return true
	}
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}
func main() {
	//n := 4
	//dislikes := [][]int{{1, 2}, {1, 3}, {2, 4}}
	//n := 3
	//dislikes := [][]int{{1, 2}, {1, 3}, {2, 3}}
	//n := 5
	//dislikes := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	n := 10
	dislikes := [][]int{{1, 2}, {3, 4}, {5, 6}, {6, 7}, {8, 9}, {7, 8}}
	ans := possibleBipartition(n, dislikes)
	fmt.Println(ans)
}
