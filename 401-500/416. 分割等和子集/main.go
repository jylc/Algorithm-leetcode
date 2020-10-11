package main

//416. 分割等和子集
//https://leetcode-cn.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	maxNum := 0

	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	target := sum / 2
	if n < 2 || sum%2 == 1 || maxNum > target {
		return false
	}
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	dp[0][nums[0]] = true

	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j < target+1; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

func main() {

}
