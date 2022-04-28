package main

import "fmt"

//https://leetcode-cn.com/problems/sort-array-by-parity/
//905. 按奇偶排序数组
func sortArrayByParity(nums []int) []int {
	even, odd := make([]int, 0), make([]int, 0)
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	even = append(even, odd...)
	return even
}
func main() {
	nums := []int{3, 1, 2, 4}
	ans := sortArrayByParity(nums)
	fmt.Println(ans)
}
