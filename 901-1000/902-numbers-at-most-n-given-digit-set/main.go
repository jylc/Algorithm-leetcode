package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/
// 902. 最大为 N 的数字组合
func atMostNGivenDigitSet(digits []string, n int) int {
	arr := []byte(strconv.Itoa(n))
	level := len(arr) - 1
	size := len(digits)
	ans := 0
	// n-1层的个数
	if size != 1 {
		ans = (pow(size, level) - 1) / (size - 1) * size
	} else {
		ans = level
	}
	count := 0
	flag := false
	// 第n层的个数
	for i := 0; i < len(arr); i++ {
		count = 0
		for j := 0; j < len(digits); j++ {
			num, _ := strconv.Atoi(digits[j])
			if num < int(arr[i]-'0') {
				count++
			}
			if num == int(arr[i]-'0') {
				flag = true
			}
		}

		for j := i + 1; j < len(arr); j++ {
			count = count * size
		}
		ans += count
		if !flag {
			break
		}

		if i+1 >= len(arr) {
			break
		}
		flag = false
	}
	if flag {
		ans++
	}

	return ans
}

func pow(a, b int) int {
	ans := 1
	for i := 0; i < b; i++ {
		ans = ans * a
	}
	return ans
}

func main() {
	/*digits := []string{"3", "4", "8"}
	n := 4*/

	digits := []string{"1", "3", "5", "7"}
	n := 100
	ans := atMostNGivenDigitSet(digits, n)
	fmt.Println(ans)
}
