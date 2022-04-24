package main

import "sort"

//https://leetcode-cn.com/problems/intersection-of-multiple-arrays/
//6041. 多个数组求交集
func intersection(nums [][]int) []int {
	m := make(map[int]int)
	length := len(nums)
	for _, num := range nums {
		for _, n := range num {
			m[n]++
		}
	}
	arr := make([]int, 0)
	for k, v := range m {
		if v == length {
			arr = append(arr, k)
		}
	}
	sort.Ints(arr)
	return arr
}
func main() {

}
