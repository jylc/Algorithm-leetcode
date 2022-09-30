// 本题为考试多行输入输出规范示例，无需提交，不计分。
package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
测试用例
2
3
911
97625999
91125426
5
113
12340
123440
12345
98765

*/
func main1() {
	n := 0
	m := 0

	fmt.Scan(&n)
	flags := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		phoneNumbers := make([]string, m)
		for j := 0; j < m; j++ {
			var x string
			fmt.Scan(&x)
			phoneNumbers[j] = x
		}
		sort.Strings(phoneNumbers)
		// flag 判断列表中所有的号码是否可用
		flag := true
		for index := len(phoneNumbers) - 1; index >= 1; index-- {
			for index1 := index - 1; index1 >= 0; index1-- {
				if strings.HasPrefix(phoneNumbers[index], phoneNumbers[index1]) {
					flag = false
				}
			}
		}
		if flag {
			flags[i] = "YES"
		} else {
			flags[i] = "NO"
		}
	}
	for _, flag := range flags {
		fmt.Println(flag)
	}
}
