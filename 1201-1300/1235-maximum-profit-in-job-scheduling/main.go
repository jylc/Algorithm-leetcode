package main

import "sort"

// https://leetcode.cn/problems/maximum-profit-in-job-scheduling
// 1235. 规划兼职工作

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	length := len(startTime)
	jobs := make([][3]int, length)
	for i := 0; i < length; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	dp := make([]int, length+1)
	for i := 1; i <= length; i++ {
		k := sort.Search(i, func(j int) bool {
			// jobs[j][1] > jobs[i-1][0]:前i个工作没有包含job[i]，实际上是job[i-1]以及之前的工作报酬
			// 因此job[j]是前j-1个的工作不包含j，dp[j]本身就是最大的值
			return jobs[j][1] > jobs[i-1][0]
		})
		dp[i] = max(dp[i-1], dp[k]+jobs[i-1][2])
	}
	return dp[length]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
