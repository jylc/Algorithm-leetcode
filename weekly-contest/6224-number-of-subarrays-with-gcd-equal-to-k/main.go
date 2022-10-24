package main

import "fmt"

// https://leetcode.cn/contest/weekly-contest-316/problems/number-of-subarrays-with-gcd-equal-to-k/
// 6224. 最大公因数等于 K 的子数组数目

func subarrayGCD(nums []int, k int) int {
	ans := 0
	for i := range nums {
		g := 0
		for _, x := range nums[i:] {
			g = gcd(g, x)
			if g%k > 0 {
				break
			}
			if g == k {
				ans++
			}
		}
	}
	return ans
}

// 辗转相除法求公因数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func main() {
	//nums := []int{9, 3, 1, 2, 6, 3}
	//k := 3
	nums := []int{4}
	k := 7
	ans := subarrayGCD(nums, k)
	fmt.Println(ans)
	ans = gcd(3, 23)
	fmt.Println(ans)
	fmt.Println(2 % 1)
}
