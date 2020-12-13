package main

//217. 存在重复元素
//https://leetcode-cn.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			m[num]++
		}
		v := m[num]
		if v >= 2 {
			return true
		}
	}
	return false
}
func main() {

}
