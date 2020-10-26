package main

//1365. 有多少小于当前数字的数字
//https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
func smallerNumbersThanCurrent(nums []int) []int {
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				if nums[j] < nums[i] {
					arr[i]++
				}
			}
		}
	}
	return arr
}

func smallerNumbersThanCurrent_1(nums []int) []int {
	cnt := [101]int{}

	for _, num := range nums {
		cnt[num]++
	}
	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}

	ans := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			ans[i] = cnt[v-1]
		}
	}
	return ans
}
func main() {

}
