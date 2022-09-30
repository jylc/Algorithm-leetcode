package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.cn/problems/get-kth-magic-number-lcci/
// 面试题 17.09. 第 k 个数

type heapData struct {
	dataM map[int]bool
	data  []int
}

func (h *heapData) Push(x interface{}) {
	v := x.(int)
	if _, exists := h.dataM[v]; exists {
		return
	} else {
		h.dataM[v] = true
	}
	h.data = append(h.data, x.(int))
}

func (h *heapData) Pop() interface{} {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[0 : n-1]
	return x
}
func (h *heapData) Len() int {
	return len(h.data)
}

func (h *heapData) Less(i, j int) bool {
	if h.data[i] < h.data[j] {
		return true
	}
	return false
}

func (h *heapData) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// 方法一：小根堆；时间复杂度：O(nlogn)；空间复杂度：O(k)
func getKthMagicNumber(k int) int {
	data := &heapData{data: make([]int, 0), dataM: make(map[int]bool)}
	heap.Init(data)
	heap.Push(data, 1)
	var v int
	for i := 0; i < k; i++ {
		v = heap.Pop(data).(int)
		a, b, c := v*3, v*5, v*7
		heap.Push(data, a)
		heap.Push(data, b)
		heap.Push(data, c)
	}
	return v
}

// 方法一：多路归并；时间复杂度：O(k)；空间复杂度：O(k)
func getKthMagicNumber1(k int) int {
	arr := make([]int, k+1)
	arr[1] = 1
	i3, i5, i7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		a, b, c := arr[i3]*3, arr[i5]*5, arr[i7]*7
		m := min(a, min(b, c))
		if m == a {
			i3++
		}
		if m == b {
			i5++
		}
		if m == c {
			i7++
		}
		arr[i] = m
	}
	return arr[k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	for i := 0; i < 7; i++ {
		ans := getKthMagicNumber(i + 1)
		fmt.Println(ans)
	}

}
