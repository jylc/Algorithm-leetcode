package main

import "math"

//https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/
//2024. 考试的最大困扰度
//滑动窗口
func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	n, l, r := len(answerKey), 0, 0
	ans = 0
	m := make(map[rune]int)
	for r < n {
		m[rune(answerKey[r])] += 1
		if math.Min(float64(m['T']), float64(m['F'])) > float64(k) {
			m[rune(answerKey[l])] -= 1
			l += 1
		}
		r += 1
		ans = int(math.Max(float64(ans), float64(r-l)))
	}
	return
}
func main() {

}
