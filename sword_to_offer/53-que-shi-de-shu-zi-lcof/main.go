package main

//剑指 Offer 53 - II. 0～n-1中缺失的数字
//https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
func missingNumber(nums []int) int {
	for i, num := range nums {
		if i != num {
			return i
		}
	}
	if nums[len(nums)-1] == len(nums)-1 {
		return len(nums)
	}
	return -1
}
func main() {

}
