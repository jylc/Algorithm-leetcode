package main

//529. 扫雷游戏
//https://leetcode-cn.com/problems/minesweeper/

var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(board, x, y)
	}
	return board
}

func dfs(board [][]byte, x, y int) {
	cnt := 0
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) {
			continue
		}
		if board[tx][ty] == 'M' {
			cnt++
		}
	}

	if cnt == 0 {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}
			dfs(board, tx, ty)
		}
	} else {
		board[x][y] = byte(cnt + '0')
	}
}
func main() {

}
