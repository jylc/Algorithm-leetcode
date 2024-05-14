package main

import "fmt"

// https://leetcode.cn/problems/minimum-rounds-to-complete-all-tasks/description/?envType=daily-question&envId=2024-05-14
// 2244. 完成所有任务需要的最少轮数
func minimumRounds(tasks []int) int {
	cnt := make(map[int]int)
	for _, task := range tasks {
		cnt[task]++
	}
	res := 0
	for _, v := range cnt {
		if v == 1 {
			return -1
		} else if v%3 == 0 {
			res += v / 3
		} else {
			res += v/3 + 1
		}
	}
	return res
}

func main() {
	/*tasks1 := []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}
	fmt.Println(minimumRounds(tasks1))
	tasks2 := []int{2, 3, 3}
	fmt.Println(minimumRounds(tasks2))*/
	tasks3 := []int{7, 7, 7, 7, 7, 7}
	fmt.Println(minimumRounds(tasks3))
}
