package main

import (
	"fmt"
	"sort"
)

//692. 前K个高频单词
//https://leetcode-cn.com/problems/top-k-frequent-words/
func topKFrequent(words []string, k int) []string {
	strMap := make(map[string]int, 0)
	var results []string
	for _, word := range words {
		strMap[word]++
	}

	for word := range strMap {
		results = append(results, word)
	}

	sort.Slice(results, func(i, j int) bool {
		s, t := results[i], results[j]
		return strMap[s] > strMap[t] || strMap[s] == strMap[t] && s < t
	})
	return results[:k]
}
func main() {
	word := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	result := topKFrequent(word, k)
	fmt.Print(result)
}
