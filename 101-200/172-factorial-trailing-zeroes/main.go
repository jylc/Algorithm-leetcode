package main

//https://leetcode-cn.com/problems/factorial-trailing-zeroes/
//172. 阶乘后的零
func trailingZeroes(n int) int {
	cnt := 0
	for n >= 5 {
		cnt += n / 5
		n /= 5
	}
	return cnt
}

func main() {

}
