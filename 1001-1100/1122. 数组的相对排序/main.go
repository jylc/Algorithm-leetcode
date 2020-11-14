package main

import (
	"fmt"
	"sort"
)

//1122. 数组的相对排序
//https://leetcode-cn.com/problems/relative-sort-array/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for _, v := range arr1 {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	tmp := make([]int, 0)

	for _, v := range arr2 {
		value := m[v]
		for i := 0; i < value; i++ {
			tmp = append(tmp, v)
		}
		m[v] = 0
	}

	arr1 = tmp
	tmp = nil
	for k, v := range m {
		if v != 0 {
			tmp = append(tmp, k)
		}
	}

	sort.Ints(tmp)
	for _, v := range tmp {
		value := m[v]
		for i := 0; i < value; i++ {
			arr1 = append(arr1, v)
		}
	}

	return arr1
}
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for i, k := range arr2 {
		m[k] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := m[x]
		rankY, hasY := m[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})
	return arr1
}
func main() {

	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	result := relativeSortArray1(arr1, arr2)
	fmt.Print(result)
}
