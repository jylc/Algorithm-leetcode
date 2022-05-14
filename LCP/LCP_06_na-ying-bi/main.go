package main

import "fmt"

//https://leetcode.cn/problems/na-ying-bi/
//LCP 06. 拿硬币
func minCount(coins []int) int {
	cnt := 0
	for _, coin := range coins {
		if coin%2 == 0 {
			cnt += coin / 2
		} else {
			cnt += coin/2 + 1
		}
	}
	return cnt
}

func main() {
	coins := []int{2, 3, 10}
	ans := minCount(coins)
	fmt.Println(ans)
}
