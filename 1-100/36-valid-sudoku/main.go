package main

//36. 有效的数独
//https://leetcode-cn.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	a := make([][]int, 9)
	b := make([][]int, 9)
	c := make([][][]int, 3)
	for i := 0; i < 9; i++ {
		a[i] = make([]int, 9)
		b[i] = make([]int, 9)
	}
	for i := 0; i < 3; i++ {
		c[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			c[i][j] = make([]int, 9)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v != '.' {
				index := v - '0' - 1
				a[i][index]++
				b[j][index]++
				c[i/3][j/3][index]++
				if a[i][index] > 1 || b[j][index] > 1 || c[i/3][j/3][index] > 1 {
					return false
				}
			}
		}
	}
	return true
}
func main() {

}
