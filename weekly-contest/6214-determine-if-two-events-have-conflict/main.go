package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/contest/weekly-contest-316/problems/determine-if-two-events-have-conflict/
// 6214. 判断两个事件是否存在冲突

func haveConflict(event1 []string, event2 []string) bool {
	if strings.Compare(event1[0], event2[0]) > 0 {
		event1, event2 = event2, event1
	}
	if strings.Compare(event1[1], event2[0]) >= 0 {
		return true
	}
	return false
}

func main() {
	event1, event2 := []string{"01:15", "02:00"}, []string{"02:00", "03:00"}
	//event1, event2 := []string{"01:00", "02:00"}, []string{"01:20", "03:00"}
	//event1, event2 := []string{"10:00", "11:00"}, []string{"14:00", "15:00"}
	//event1, event2 := []string{"14:13", "22:08"}, []string{"02:40", "08:08"}
	ans := haveConflict(event1, event2)
	fmt.Println(ans)
}
