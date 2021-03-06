package main

import "fmt"

//5689. 统计匹配检索规则的物品数量
//https://leetcode-cn.com/contest/weekly-contest-230/problems/count-items-matching-a-rule/

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var index int
	if ruleKey == "type" {
		index = 0
	}
	if ruleKey == "color" {
		index = 1
	}
	if ruleKey == "name" {
		index = 2
	}
	lens := len(items)
	results := make([][]string, 0)
	for i := 0; i < lens; i++ {
		if items[i][index] == ruleValue {
			results = append(results, items[i])
		}
	}
	return len(results)
}

func main() {
	items := [][]string{{"phone", "blue", "pixel"},
		{"computer", "silver", "lenovo"},
		{"phone", "gold", "iphone"}}
	ruleKey := "type"
	ruleValue := "phone"
	result := countMatches(items, ruleKey, ruleValue)
	fmt.Print("result = ", result)
}
