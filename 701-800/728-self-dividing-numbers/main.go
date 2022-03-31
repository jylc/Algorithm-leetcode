package main

import "fmt"

//https://leetcode-cn.com/problems/self-dividing-numbers/
//728. 自除数
func selfDividingNumbers(left int, right int) []int {
	l, r := left, right
	arr := make([]int, 0)
	for l <= r {
		tmp := l
		flag := true
		for tmp != 0 {
			if tmp%10 == 0 {
				flag = false
				break
			}
			if l%(tmp%10) != 0 {
				flag = false
				break
			}
			tmp /= 10
		}
		if flag {
			arr = append(arr, l)
		}
		l++
	}
	return arr
}
func main() {
	res := selfDividingNumbers(47, 85)
	fmt.Println(res)
}
