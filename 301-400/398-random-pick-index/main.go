package main

import (
	"fmt"
	"math/rand"
	"time"
)

//https://leetcode-cn.com/problems/random-pick-index/
//398. 随机数索引

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for idx, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = make([]int, 0)
		}
		m[num] = append(m[num], idx)
	}
	rand.Seed(time.Now().Unix())
	return Solution{m: m}
}

func (this *Solution) Pick(target int) int {
	n := len(this.m[target])
	data := rand.Int() % n
	return this.m[target][data]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
func main() {
	nums := []int{1, 2, 3, 3, 3}
	solution := Constructor(nums)
	for i := 0; i < 3; i++ {
		ans := solution.Pick(3)
		fmt.Println(ans)
	}

}
