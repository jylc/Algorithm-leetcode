package main

//79. 单词搜索
//https://leetcode-cn.com/problems/word-search/
type pair struct {
	x, y int
}

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := 0; i < h; i++ {
		vis[i] = make([]bool, w)
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		defer func() {
			vis[i][j] = false
		}()

		for _, dir := range directions {
			if newX, newY := i+dir.x, j+dir.y; 0 <= newX && newX < h && 0 <= newY && newY < w && !vis[newX][newY] {
				if dfs(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
}
