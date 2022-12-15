package main

import (
	"fmt"
)

// https://leetcode.cn/problems/sum-of-digits-of-string-after-convert/
// 1945. 字符串转化后的各位数字之和

func getLucky(s string, k int) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += add(int(s[i] - 'a' + 1))
	}
	for j := 1; j < k; j++ {
		ans = add(ans)
	}
	return ans
}

func add(num int) int {
	ans := 0
	for num != 0 {
		ans += num % 10
		num /= 10
	}
	return ans
}

func main() {
	ans := getLucky("zbax", 2)
	fmt.Println(ans)
}
