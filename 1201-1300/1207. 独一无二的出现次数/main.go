package main

//1207. 独一无二的出现次数
//https://leetcode-cn.com/problems/unique-number-of-occurrences/

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	n := make(map[int]bool)
	for _, v := range m {
		if _, ok := n[v]; ok {
			return false
		} else {
			n[v] = true
		}
	}
	return true
}

func main() {

}
