package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/corporate-flight-bookings/
//1109. 航班预订统计
func corpFlightBookings(bookings [][]int, n int) []int {
	m := make([]int, n+1)
	for _, booking := range bookings {
		for x, y := booking[0], booking[1]; x <= y; x++ {
			m[x] += booking[2]
		}
	}
	fmt.Println(m)
	return m[1:]
}
func main() {

}
