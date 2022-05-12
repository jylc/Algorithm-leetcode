package main

import "fmt"

//https://leetcode.cn/problems/permutations-ii/
//47. 全排列 II
func permuteUnique(nums []int) (ans [][]int) {
	n := len(nums)
	visited := make([]bool, n)

	var bfs func(arr []int, pos int)
	bfs = func(arr []int, pos int) {
		if pos == n {
			tmp := make([]int, n)
			copy(tmp, arr)
			ans = append(ans, tmp)
			return
		}
		m := make(map[int]bool)
		for i := 0; i < n; i++ {
			m[nums[i]] = false
		}
		for i := 0; i < n; i++ {
			if !visited[i] && !m[nums[i]] {
				arr[pos] = nums[i]
				visited[i] = true
				bfs(arr, pos+1)
				m[nums[i]] = true
				visited[i] = false
			}
		}
	}
	arr := make([]int, n)
	bfs(arr, 0)

	return
}
func main() {
	num := []int{1, 1, 2}
	ans := permuteUnique(num)
	fmt.Println(ans)
}
