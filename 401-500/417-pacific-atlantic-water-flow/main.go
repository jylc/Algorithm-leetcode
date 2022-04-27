package main

import "fmt"

//https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
//417. 太平洋大西洋水流问题
func pacificAtlantic(heights [][]int) [][]int {
	var dfs func(i, j, h int)
	moves := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	_, right, _, bottom := 0, len(heights[0]), 0, len(heights)
	first := make([][]bool, bottom)
	second := make([][]bool, bottom)
	ans := make([][]int, 0)
	for i := 0; i < bottom; i++ {
		first[i] = make([]bool, right)
		second[i] = make([]bool, right)
	}
	flag := true
	dfs = func(i, j, h int) {
		if heights[i][j] >= h {
			if flag {
				first[i][j] = true
			} else {
				second[i][j] = true
			}
		} else {
			return
		}
		for _, move := range moves {
			x, y := i+move[0], j+move[1]
			if x >= 0 && x < bottom && y >= 0 && y < right {
				if flag {
					if !first[x][y] {
						dfs(x, y, heights[i][j])
					}
				} else {
					if !second[x][y] {
						dfs(x, y, heights[i][j])
					}
				}
			}
		}
	}
	//左侧
	for i := 0; i < bottom; i++ {
		dfs(i, 0, heights[i][0])
	}

	//上方
	for i := 0; i < right; i++ {
		dfs(0, i, heights[0][i])
	}
	flag = false
	//右侧
	for i := 0; i < bottom; i++ {
		dfs(i, right-1, heights[i][right-1])
	}

	//下方
	for i := 0; i < right; i++ {
		dfs(bottom-1, i, heights[bottom-1][i])
	}
	for i := 0; i < bottom; i++ {
		for j := 0; j < right; j++ {
			if first[i][j] && second[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
func main() {
	//heights := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	heights := [][]int{{1, 1}, {1, 1}, {1, 1}}
	ans := pacificAtlantic(heights)
	fmt.Println(ans)
}
