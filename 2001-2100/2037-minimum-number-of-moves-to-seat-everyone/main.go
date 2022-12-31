package main

import "sort"

// https://leetcode.cn/problems/minimum-number-of-moves-to-seat-everyone/
// 2037. 使每位学生都有座位的最少移动次数

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ans := 0
	for i, seat := range seats {
		ans += abs(seat - students[i])
	}
	return ans
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

}
