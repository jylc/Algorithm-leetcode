package main

import "fmt"

//面试题 01.08. 零矩阵
//https://leetcode-cn.com/problems/zero-matrix-lcci/
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	row, col := len(matrix), len(matrix[0])
	var arrs [][]int
	arrs = make([][]int, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				arr := make([]int, 2)
				arr[0] = i
				arr[1] = j
				arrs = append(arrs, arr)
			}
		}
	}
	for _, arr := range arrs {
		x, y := arr[0], arr[1]
		for i := 0; i < col; i++ {
			matrix[x][i] = 0
		}

		for i := 0; i < row; i++ {
			matrix[i][y] = 0
		}
	}
}
func main() {
	arrs := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(arrs)
	for i := 0; i < len(arrs); i++ {
		for j := 0; j < len(arrs[0]); j++ {
			fmt.Print(arrs[i][j], " ")
		}
		fmt.Println()
	}
}
