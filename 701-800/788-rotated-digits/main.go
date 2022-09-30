package main

import "fmt"

// https://leetcode.cn/problems/rotated-digits/
// 788. 旋转数字

func rotatedDigits(n int) int {
	m := make(map[int]int)
	m[0] = 0
	m[1] = 0
	m[8] = 0
	m[2] = 0
	m[5] = 0
	m[6] = 0
	m[9] = 0
	cnt := 0
	flag := true
	for i := 1; i <= n; i++ {
		arr := parse(i)
		for _, v := range arr {

			if _, ok := m[v]; !ok {
				flag = false
				break
			} else {
				m[v]++
			}
		}

		if flag {
			if m[2] != 0 || m[5] != 0 || m[6] != 0 || m[9] != 0 {
				// 在0,1,8,2,5,6,9以外的数都不是好数 且 需要包含2,5,6,9至少一个
				cnt++
			}
		} else {
			flag = true
		}
	}
	return cnt
}

func parse(num int) []int {
	var arr []int
	for num != 0 {
		arr = append(arr, num%10)
		num = num / 10
	}
	return arr
}

func t(n int) {
	a := make([]int, 6)
	dp := make([][2]int, 6)
	for i := range a {
		dp[i] = [2]int{-1, -1}
	}
	l := 0
	for n > 0 {
		l++
		a[l] = n % 10
		n /= 10
	}
	fmt.Println(a)
}
func main() {
	t(12)
}
