package main

//剑指 Offer 53 - I. 在排序数组中查找数字 I
//https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func search(nums []int, target int) int {
	cnt := 0
	for _, num := range nums {
		if num == target {
			cnt++
		}
	}
	return cnt
}
func main() {

}
