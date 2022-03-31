package main

import (
	"fmt"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/reformat-date/
//1507. 转变日期格式
func reformatDate(date string) string {
	parts := strings.Split(date, " ")
	Day, _ := strconv.Atoi(parts[0][:len(parts[0])-2])
	Months := map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4,
		"May": 5,
		"Jun": 6,
		"Jul": 7,
		"Aug": 8,
		"Sep": 9,
		"Oct": 10,
		"Nov": 11,
		"Dec": 12,
	}
	Month := Months[parts[1]]
	Year, _ := strconv.Atoi(parts[2])
	return fmt.Sprintf("%d-%02d-%02d", Year, Month, Day)
}
func main() {

}
