package main

//5701. 仅执行一次字符串交换能否使两个字符串相等
//https://leetcode-cn.com/contest/weekly-contest-232/problems/check-if-one-string-swap-can-make-strings-equal/
func areAlmostEqual(s1 string, s2 string) bool {
	n := len(s1)
	stack := make([]uint8, 0)
	cnt := 0
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			cnt++
			if len(stack) == 0 {
				stack = append(stack, s1[i])
				stack = append(stack, s2[i])
			} else {
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				if a == s1[i] && b == s2[i] {
					stack = stack[:0]
				}
			}
		}
	}

	if len(stack) == 0 && cnt == 2 || cnt == 0 {
		return true
	}
	return false
}

func main() {

}
