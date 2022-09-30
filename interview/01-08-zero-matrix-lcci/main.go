package main

// https://leetcode.cn/problems/zero-matrix-lcci/
// 面试题 01.08. 零矩阵

type pos struct {
	row int
	col int
}

func setZeroes(matrix [][]int) {
	p := make([]pos, 0)
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				p = append(p, pos{
					row: i,
					col: j,
				})
			}
		}
	}
	for _, v := range p {
		for i := 0; i < col; i++ {
			matrix[v.row][i] = 0
		}
		for i := 0; i < row; i++ {
			matrix[i][v.col] = 0
		}
	}
}

func main() {

}
