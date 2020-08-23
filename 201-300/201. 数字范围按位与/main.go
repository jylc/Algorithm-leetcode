package main

//201. 数字范围按位与
//https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/
func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	//&运算寻找公共前缀
	for m < n {
		m = m >> 1
		n = n >> 1
		shift++
	}
	return m << shift
}
func main() {

}
