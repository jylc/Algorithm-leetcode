package main

import "fmt"

//https://leetcode.cn/problems/delete-columns-to-make-sorted/
//944. 删列造序
func minDeletionSize(strs []string) (ans int) {
	row, col := len(strs), len(strs[0])
	ans = 0
	for i := 0; i < col; i++ {
		for j := 0; j < row-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				ans++
				break
			}
		}
	}
	return
}
func main() {
	strs := []string{"zyx", "wvu", "tsr"}
	ans := minDeletionSize(strs)
	fmt.Println(ans)
}
