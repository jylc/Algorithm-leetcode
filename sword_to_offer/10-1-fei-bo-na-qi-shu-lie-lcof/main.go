package main

//https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
//剑指 Offer 10- I. 斐波那契数列

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		p, q, r := 0, 0, 1
		for i := 2; i <= n; i++ {
			p = q
			q = r
			r = (p + q) % (1e9 + 7)
		}
		return r
	}
}

func main() {

}
