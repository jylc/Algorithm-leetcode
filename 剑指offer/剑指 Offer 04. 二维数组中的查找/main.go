package main

import "fmt"

//剑指 Offer 04. 二维数组中的func findNumberIn2DArray(matrix [][]int, target int) bool {
//	var binary func(arr []int, left, right, target int) bool
//	binary = func(arr []int, left, right, target int) bool {
//		if left > right {
//			return false
//		}
//		mid := (right+left)/2
//		if arr[mid] == target {
//			return true
//		} else if arr[mid] < target {
//			return binary(arr, mid+1, right, target)
//		} else {
//			return binary(arr, left, mid-1, target)
//		}
//	}
//	for i := 0; i < len(matrix); i++ {
//		if binary(matrix[i], 0, len(matrix[0])-1, target) {
//			return true
//		}
//	}
//	return false
//}查找
//https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/



func main() {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	result := findNumberIn2DArray(arr, 20)
	fmt.Println("result:", result)
}
