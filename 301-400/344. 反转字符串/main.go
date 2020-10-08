package main

//344. 反转字符串
//https://leetcode-cn.com/problems/reverse-string/
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	p, q := 0, len(s)-1
	for p < q {
		s[p], s[q] = s[q], s[p]
		p++
		q--
	}
}
func main() {

}
