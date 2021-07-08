package main

//930. 和相同的二元子数组
//https://leetcode-cn.com/problems/binary-subarrays-with-sum/
func numSubarraysWithSum1(nums []int, goal int) int {
	length := len(nums)
	ans := 0
	for i := 0; i < length; i++ {
		cnt := 0
		for j := i; j < length; j++ {
			cnt += nums[j]
			if cnt == goal {
				ans++
			} else if cnt > goal {
				break
			}
		}
	}
	return ans
}

func numSubarraysWithSum2(nums []int, goal int) int {
	cnt := map[int]int{}
	ans := 0
	sum := 0
	for _, num := range nums {
		cnt[sum]++
		sum += num
		ans += cnt[sum-goal]
	}
	return ans
}
func main() {

}
