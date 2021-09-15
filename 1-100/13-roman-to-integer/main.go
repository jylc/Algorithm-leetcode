package main

//https://leetcode-cn.com/problems/roman-to-integer/
//13. 罗马数字转整数
func romanToInt(s string) int {
	roman := make(map[byte]int)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000
	n := len(s)
	ans := 0
	for i, _ := range s {
		value := roman[s[i]]
		if i < n-1 && value < roman[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}
func main() {

}
