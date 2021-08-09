package main

//https://leetcode-cn.com/problems/n-th-tribonacci-number/
//1137. 第 N 个泰波那契数
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	d := 0
	for i := 3; i <= n; i++ {
		d = a + b + c
		a = b
		b = c
		c = d
	}
	return d
}
func main() {

}
