package main

// https://leetcode.cn/problems/fruit-into-baskets/
// 904. 水果成篮

// 时间复杂度：O(n);空间复杂度：O(1)
func totalFruit(fruits []int) int {
	cnt := make(map[int]int)
	left := 0
	ans := 0
	for right, value := range fruits {
		cnt[value]++
		for len(cnt) > 2 {
			x := fruits[left]
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

}
