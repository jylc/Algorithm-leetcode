package main

//https://leetcode-cn.com/problems/smallest-range-i/
//908. 最小差值 I

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func smallestRangeI(nums []int, k int) int {
	max, min := nums[0], nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}
	return maxNum(0, max-min-2*k)
}

func main() {

}
