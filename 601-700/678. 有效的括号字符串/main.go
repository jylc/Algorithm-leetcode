package main

//678. 有效的括号字符串
//https://leetcode-cn.com/problems/valid-parenthesis-string/
func checkValidString(s string) bool {
	min, max := 0, 0

	for _, w := range s {
		if w == '(' {
			max++
			min++
		} else if w == '*' {
			max++
			if min > 0 {
				min--
			}
		} else if w == ')' {
			max--
			if max < 0 {
				return false
			}
			if min > 0 {
				min--
			}
		}
	}
	return min == 0
}
func main() {

}
