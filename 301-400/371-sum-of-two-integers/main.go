package main

//https://leetcode-cn.com/problems/sum-of-two-integers/
//371. 两整数之和
func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

func main() {
}
