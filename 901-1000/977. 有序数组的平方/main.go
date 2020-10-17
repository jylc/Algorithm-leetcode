package main

import "sort"

//977. 有序数组的平方
//https://leetcode-cn.com/problems/squares-of-a-sorted-array/

func sortedSquares(A []int) []int {
	arr := make([]int, 0)
	for _, i := range A {
		arr = append(arr, i*i)
	}

	sort.Ints(arr)
	return arr
}
func main() {

}
