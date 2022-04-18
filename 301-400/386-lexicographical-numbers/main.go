package main

import "fmt"

//https://leetcode-cn.com/problems/lexicographical-numbers/
//386. 字典序排数
func lexicalOrder(n int) []int {
	arr := make([]int, n)
	num := 1
	for i := range arr {
		arr[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	fmt.Println(arr)
	return arr
}
func main() {
	lexicalOrder(13)
}
