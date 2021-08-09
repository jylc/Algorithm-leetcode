package main

import (
	"container/heap"
	"sort"
)

//https://leetcode-cn.com/problems/super-ugly-number/
//313. 超级丑数

//方法一：最小堆
/*
1、将1放入堆中
2、从队列中选出最小的数与primes中每一个元素相乘
3、重复第二步，知道到达第n个
*/
type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func nthSuperUglyNumber(n int, primes []int) (ugly int) {
	set := map[int]bool{1: true}
	h := &hp{[]int{1}}
	for j := 0; j < n; j++ {
		ugly = heap.Pop(h).(int)
		for i := 0; i < len(primes); i++ {
			if next := ugly * primes[i]; !set[next] {
				set[next] = true
				heap.Push(h, next)
			}
		}
	}
	return
}

func main() {

}
