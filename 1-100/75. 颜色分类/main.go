package main

//75. 颜色分类
//https://leetcode-cn.com/problems/sort-colors/
func sortColors(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	i, p0, p1 := 0, 0, 0
	for i < length {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
		i++
	}
}
func main() {

}
