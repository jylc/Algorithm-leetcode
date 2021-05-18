package main

import "fmt"

//1442. 形成两个异或相等数组的三元组数目
//https://leetcode-cn.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
func countTriplets(arr []int) int {
	length := len(arr)
	a, b := 0, 0
	cnt := 0
	for k := length - 1; k > 0; k-- {
		for j := k; j > 0; j-- {
			b ^= arr[j]
			for i := j - 1; i >= 0; i-- {
				a ^= arr[i]
				if a == b {
					cnt++
				}
			}
			a = 0
		}
		b = 0
	}
	return cnt
}
func main() {
	arr := []int{1, 3, 5, 7, 9}
	result := countTriplets(arr)
	fmt.Print("result = ", result)
}
