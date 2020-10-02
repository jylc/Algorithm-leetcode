package main

//771. 宝石与石头
//https://leetcode-cn.com/problems/jewels-and-stones/
func numJewelsInStones(J string, S string) int {
	m := make(map[rune]int)
	for _, v := range J {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}
	}
	cnt := 0
	for _, k := range S {
		if _, ok := m[k]; ok {
			cnt++
		}
	}
	return cnt
}
func main() {

}
