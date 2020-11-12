package main

//922. 按奇偶排序数组 II
//https://leetcode-cn.com/problems/sort-array-by-parity-ii/
func sortArrayByParityII(a []int) []int {
	ans := make([]int, len(a))
	i := 0
	for _, v := range a {
		if v%2 == 0 {
			ans[i] = v
			i += 2
		}
	}
	i = 1
	for _, v := range a {
		if v%2 == 1 {
			ans[i] = v
			i += 2
		}
	}
	return ans
}
func main() {

}
