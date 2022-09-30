package main

// https://leetcode.cn/problems/missing-two-lcci/
// 面试题 17.19. 消失的两个数字

// 数学问题；时间复杂度:O(n);空间复杂度:O(1)
func missingTwo(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	n := len(nums) + 2
	// 两个消失的数之和
	sumTwo := (1+n)*n/2 - sum
	limits := int(sumTwo / 2)
	sum = 0
	for _, num := range nums {
		if num <= limits {
			sum += num
		}
	}
	one := (1+limits)*limits/2 - sum
	return []int{one, sumTwo - one}
}

func main() {

}
