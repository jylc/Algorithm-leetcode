package main

//566. 重塑矩阵
//https://leetcode-cn.com/problems/reshape-the-matrix/
func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	ans, i, j := make([][]int, r), 0, 0
	for i := range ans {
		ans[i] = make([]int, c)
	}
	for _, row := range nums {
		for _, ele := range row {
			ans[i][j] = ele
			j++
			if j == c {
				i++
				j = 0
			}
		}
	}
	return ans
}

func main() {

}
