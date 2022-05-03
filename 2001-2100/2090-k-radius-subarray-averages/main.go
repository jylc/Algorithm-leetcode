package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/k-radius-subarray-averages/
//2090. 半径为 k 的子数组平均值
//1、滑动窗口
func getAverages1(nums []int, k int) []int {
	left, mid, right := 0, 0, 0
	n := len(nums)
	ans := make([]int, n)
	sum := 0
	for right < n || left < n {
		mid = (right + left) / 2
		if right < 2*k && right < n {
			ans[mid] = -1
			sum += nums[right]
			right++
		} else if left > (n-1)-2*k && left < n {
			ans[mid] = -1
			left++
		} else if right-left == 2*k {
			mid = (right + left) / 2
			sum += nums[right]
			ans[mid] = int(math.Floor(float64(sum / (2*k + 1))))
			sum -= nums[left]
			right++
			left++
		}
	}
	return ans
}

func getAverages2(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	sum := 0
	if 2*k+1 <= n {
		for i := 0; i < 2*k+1; i++ {
			sum += nums[i]
		}
		for i := k; i < n-k; i++ {
			if i != k {
				sum += nums[i+k] - nums[i-k-1]
			}
			ans[i] = sum / (2*k + 1)
		}
	}
	return ans
}
func getAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		if i < k || i+k >= len(nums) {
			ans[i] = -1
		}
		sum += num
		if i >= 2*k {
			ans[i-k] = sum / (2*k + 1)
			sum -= nums[i-2*k]
		}
	}
	return ans
}
func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	/*nums := []int{8}
	k := 100000*/
	ans := getAverages(nums, k)
	fmt.Println(ans)
}
