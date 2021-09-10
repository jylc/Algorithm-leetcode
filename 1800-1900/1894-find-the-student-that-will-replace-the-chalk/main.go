package main

//https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk/
//1894. 找到需要补充粉笔的学生编号
func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	k %= sum
	for i, v := range chalk {
		if k < v {
			return i
		}
		k -= v
	}
	return 0
}
func main() {

}
