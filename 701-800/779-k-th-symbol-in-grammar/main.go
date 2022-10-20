package main

// https://leetcode.cn/problems/k-th-symbol-in-grammar
// 779. 第K个语法符号

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k <= (1 << (n - 2)) {
		return kthGrammar(n-1, k)
	}
	return kthGrammar(n-1, k-(1<<(n-2))) ^ 1
}
func main() {

}
