package main

//486. 预测赢家
//https://leetcode-cn.com/problems/predict-the-winner/
func PredictTheWinner(nums []int) bool {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = nums[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[length-1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

}
