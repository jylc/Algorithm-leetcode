package main

import (
	"container/heap"
	"sort"
)

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:4]
}

func smallestK2(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h hp) Push(interface{}) {

}

func (h hp) Pop() (_ interface{}) {
	return
}

func main() {

}
