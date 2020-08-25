package main

import "fmt"

//491. 递增子序列
//https://leetcode-cn.com/problems/increasing-subsequences/
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)
	dfs = func(index int, lis []int) {
		if len(lis) >= 2 {
			dest := make([]int, len(lis))
			copy(dest, lis)
			res = append(res, dest)
		}
		//保证一层循环中不会选择重复的值就ok了，画一下可能会比较清晰
		var visit [201]bool
		for i := index; i < len(nums); i++ {
			if visit[nums[i]+100] {
				continue
			}
			if len(lis) == 0 || nums[i] >= lis[len(lis)-1] {
				visit[nums[i]+100] = true
				dfs(i+1, append(lis, nums[i]))
			}
		}
	}
	dfs(0, []int{})
	return res
}

func main() {
	arr := []int{4, 6, 7, 7}
	result := findSubsequences(arr)
	for _, v := range result {
		for _, res := range v {
			fmt.Print(res, " ")
		}
		fmt.Println()
	}
}
