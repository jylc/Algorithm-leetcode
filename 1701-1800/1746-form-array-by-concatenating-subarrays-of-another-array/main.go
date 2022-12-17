package main

// https://leetcode.cn/problems/form-array-by-concatenating-subarrays-of-another-array
// 1764. 通过连接另一个数组的子数组得到一个数组

func canChoose(groups [][]int, nums []int) bool {
next:
	for _, g := range groups {
		for len(nums) >= len(g) {
			if equal(nums[:len(g)], g) {
				nums = nums[len(g):]
				continue next
			}
			nums = nums[1:]
		}
		return false
	}
	return true
}

func equal(a, b []int) bool {
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}
	return true
}
func main() {

}
