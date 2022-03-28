package main

//https://leetcode-cn.com/problems/toeplitz-matrix/
//766. 托普利茨矩阵
func isToeplitzMatrix(matrix [][]int) bool {
	r, c := len(matrix), len(matrix[0])
	for i := 0; i < r-1; i++ {
		for j := 0; j < c-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
func main() {

}
