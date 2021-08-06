package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
//1337. 矩阵中战斗力最弱的 K 行
func kWeakestRows(mat [][]int, k int) []int {
	nums := make(map[int]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(mat); i++ {
		cnt := 0
		for j := 0; j < len(mat[i]); j++ {
			cnt += mat[i][j]
		}
		nums[i] = cnt
		res = append(res, i)
	}
	sort.Slice(res, func(i, j int) bool {
		a, b := nums[res[i]], nums[res[j]]
		return a < b || (a == b && res[i] < res[j])
	})
	return res[:k]
}
func main() {

}
