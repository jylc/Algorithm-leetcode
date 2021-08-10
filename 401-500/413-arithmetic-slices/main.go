package main

//https://leetcode-cn.com/problems/arithmetic-slices/
//413. 等差数列划分
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	d, t := nums[1]-nums[0], 0
	cnt := 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == d {
			t++
		} else {
			d, t = nums[i]-nums[i-1], 0
		}
		cnt += t
	}
	return cnt
}
func main() {

}
