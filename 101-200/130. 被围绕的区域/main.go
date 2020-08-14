package main

import "fmt"

//130. 被围绕的区域
//https://leetcode-cn.com/problems/surrounded-regions/

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	//深度优先搜索边界的‘O’
	var dfs func(board [][]byte, i, j int)
	dfs = func(board [][]byte, i, j int) {
		if i > len(board)-1 || i < 0 || j > len(board[0])-1 || j < 0 {
			return
		}
		if board[i][j] == 'O' {
			board[i][j] = 'Y'
			dfs(board, i, j+1)
			dfs(board, i, j-1)
			dfs(board, i+1, j)
			dfs(board, i-1, j)
		}
	}

	for i := 0; i < len(board); i++ {
		dfs(board, i, 0)
		dfs(board, i, len(board[0])-1)
	}

	for i := 0; i < len(board[0]); i++ {
		dfs(board, 0, i)
		dfs(board, len(board)-1, i)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
		fmt.Println()
	}
}
func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}
