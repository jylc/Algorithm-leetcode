package main

import "fmt"

// https://leetcode.cn/problems/check-if-number-is-a-sum-of-powers-of-three/description/
// 1780. 判断一个数字是否可以表示成三的幂的和

func checkPowersOfThree(n int) bool {
	for n != 0 {
		tmp := n % 3
		if tmp != 1 && tmp != 0 {
			return false
		}
		n = n / 3
	}
	return true
}
func main() {
	ans := checkPowersOfThree(21)
	fmt.Println(ans)
}
