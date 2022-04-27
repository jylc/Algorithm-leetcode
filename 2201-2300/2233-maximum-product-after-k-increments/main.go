package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/maximum-product-after-k-increments/
//2233. K 次增加后的最大乘积
func maximumProduct(nums []int, k int) int {
	h := hp{nums}
	for heap.Init(&h); k > 0; k-- {
		h.IntSlice[0]++
		heap.Fix(&h, 0)
	}
	ans := 1
	for _, num := range nums {
		ans = (ans * num) % (1e9 + 7)
	}
	return ans
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{})   {}
func (h *hp) Pop() (_ interface{}) { return }

func main() {
	nums := []int{24, 5, 64, 53, 26, 38}
	k := 54
	ans := maximumProduct(nums, k)
	fmt.Println(ans)
}
