package main

import "fmt"

// https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch/
// 1700. 无法吃午餐的学生数量

// 时间复杂度：O(n);空间复杂度：O(C)
func countStudents(students []int, sandwiches []int) int {
	stuMap := make(map[int]int)
	for _, student := range students {
		stuMap[student]++
	}
	for _, sandwich := range sandwiches {
		v := stuMap[sandwich]
		if v != 0 {
			stuMap[sandwich]--
		} else {
			break
		}
	}
	ans := 0
	for _, v := range stuMap {
		ans += v
	}
	return ans
}
func main() {
	/*stu := []int{1, 1, 0, 0}
	sand := []int{0, 1, 0, 1}*/
	stu := []int{1, 1, 1, 0, 0, 1}
	sand := []int{1, 0, 0, 0, 1, 1}
	ans := countStudents(stu, sand)
	fmt.Println(ans)
}
