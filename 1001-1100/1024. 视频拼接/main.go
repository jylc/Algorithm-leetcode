package main

import "math"

//1024. 视频拼接
//https://leetcode-cn.com/problems/video-stitching/
func videoStitching(clips [][]int, T int) int {
	const inf = math.MaxInt64 - 1
	dp := make([]int, T+1)
	for i, _ := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}

	if dp[T] == inf {
		return -1
	}
	return dp[T]
}
func main() {
	arrs := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	videoStitching(arrs, 10)
}
