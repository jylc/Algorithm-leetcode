package main

import "fmt"

//https://leetcode.cn/problems/count-special-quadruplets/
//1995. 统计特殊四元组
func countQuadruplets1(nums []int) int {
	n := len(nums)
	cnt := 0
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				for l := k + 1; l < n; l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}
func countQuadruplets(nums []int) (ans int) {
	cnt := map[int]int{}
	for c := len(nums) - 2; c >= 2; c-- {
		cnt[nums[c+1]]++
		for a, x := range nums[:c] {
			for _, y := range nums[a+1 : c] {
				if sum := x + y + nums[c]; cnt[sum] > 0 {
					ans += cnt[sum]
				}
			}
		}
	}
	return
}
func main() {
	nums := []int{28, 8, 49, 85, 37, 90, 20, 8}
	ans := countQuadruplets(nums)
	fmt.Println(ans)
}
