package _4__最小路径和

//64. 最小路径和
//https://leetcode-cn.com/problems/minimum-path-sum/
func minPathSum(grid [][]int) int {
	row := len(grid)    //行
	col := len(grid[0]) //列

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i != 0 && j != 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			} else if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			}
		}
	}
	return grid[row-1][col-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
