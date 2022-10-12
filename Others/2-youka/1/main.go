package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param questions int整型一维数组
 * @return int整型
 */
func halfQuestions(questions []int) int {
	// write code here
	questionsNum := len(questions)
	studentsNum := questionsNum / 4
	typeMap := make(map[int]int)
	for _, question := range questions {
		typeMap[question]++
	}
	typesNum := len(typeMap)
	if typesNum <= 2*studentsNum {
		return typesNum
	} else {
		return 2 * studentsNum
	}
}

func main() {
	var l []int
	halfQuestions(l)
}
