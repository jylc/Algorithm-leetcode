package main

import "fmt"

//https://leetcode.cn/problems/count-integers-with-even-digit-sum/
//2180. 统计各位数字之和为偶数的整数个数
func countEven(num int) (ans int) {
	ans = 0
	for i := 1; i <= num; i++ {
		if i <= num && isValid(i) {
			ans++
		}
	}
	return
}

func isValid(num int) bool {
	cnt := 0
	for num != 0 {
		cnt += num % 10
		num /= 10
	}
	if cnt%2 == 0 {
		return true
	}
	return false
}
func main() {
	ans := countEven(30)
	fmt.Println(ans)
}
