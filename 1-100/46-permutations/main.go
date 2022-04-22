package main

import "fmt"

//https://leetcode-cn.com/problems/permutations/
//46. 全排列

func permute(nums []int) [][]int {
	n := len(nums)
	arrs := make([][]int, 0)
	visited := make([]bool, n)
	arr := make([]int, 0)

	var backtrack func(visited []bool, arr []int)
	backtrack = func(visited []bool, arr []int) {
		if len(arr) == len(nums) {
			tmp := make([]int, n)
			copy(tmp, arr)
			arrs = append(arrs, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				visited[i] = true
				arr = append(arr, nums[i])
				backtrack(visited, arr)
				visited[i] = false
				arr = arr[:len(arr)-1]
			}
		}
	}
	backtrack(visited, arr)
	return arrs
}
func main() {
	nums := []int{5, 4, 6, 2}
	ans := permute(nums)
	fmt.Println(ans)
}
