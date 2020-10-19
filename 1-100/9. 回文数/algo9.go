package ___回文数

//9. 回文数
//https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertNum := 0
	for x > revertNum {
		revertNum = revertNum*10 + x%10
		x = x / 10
	}
	return x == revertNum || x == revertNum/10
}
