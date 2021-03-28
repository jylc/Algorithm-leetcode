package main

//5714. 替换字符串中的括号内容
//https://leetcode-cn.com/contest/weekly-contest-234/problems/evaluate-the-bracket-pairs-of-a-string/
func evaluate(s string, knowledge [][]string) string {
	n := len(s)
	flag := false
	index := 0
	last := 0
	var arr string
	m := make(map[string]string)
	for i := 0; i < len(knowledge); i++ {
		m[knowledge[i][0]] = knowledge[i][1]
	}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			flag = true
			index = i
		} else if s[i] == ')' {
			flag = false
			last = i
			str := s[index+1 : last+1]
			if v, ok := m[str]; ok {
				arr += v
			} else {
				arr += "?"
			}
		}
		if !flag && s[i] != ')' {
			arr += string(s[i])
		}
	}
	return arr
}
func main() {

}
