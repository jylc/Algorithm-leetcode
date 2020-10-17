package main

//698. 划分为k个相等的子集
//https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/

func canPartitionKSubsets(nums []int, k int) bool {
	sum, max := 0, 0
	used := make([]bool, len(nums))
	for _, num := range nums {
		sum += num
		if max < num {
			max = num
		}
	}
	if sum%k != 0 || max > sum/k {
		return false
	}
	return backtracking(nums, k, sum/k, 0, 0, used)

}

func backtracking(nums []int, k, target, cur, start int, used []bool) bool {
	if k == 0 {
		return true
	}
	if cur == target {
		return backtracking(nums, k-1, target, 0, 0, used)
	}

	for i := start; i < len(nums); i++ {
		if used[i] == false && nums[i]+cur <= target {
			used[i] = true
			if backtracking(nums, k, target, cur+nums[i], i+1, used) {
				return true
			}
			used[i] = false
		}
	}
	return false
}
func main() {

}
