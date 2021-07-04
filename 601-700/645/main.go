package main

//645. 错误的集合
//https://leetcode-cn.com/problems/set-mismatch/
func findErrorNums(nums []int) []int {
	arr := make([]int, len(nums))
	for _, num := range nums {
		arr[num-1]++
	}
	ans := make([]int, 2)
	for i, v := range arr {
		if v == 2 {
			ans[0] = i + 1
		} else if v == 0 {
			ans[1] = i + 1
		}
	}
	return ans
}
func main() {

}
