package main

import "fmt"

func itoa(nums []int32) string {
	var ans []rune
	for _, num := range nums {
		if num >= 0 && num <= 9 {
			ans = append(ans, num+'0')
		} else {
			ans = append(ans, num)
		}
	}
	return string(ans)
}
func ambiguousCoordinates(s string) string {
	// write code here
	var nums []int32
	for _, num := range s {
		nums = append(nums, num-'0')
	}
	ans := make([]string, 0)
	for i := 1; i < len(nums)-1; i++ {
		leftNums := nums[0:i] //左边
		rightNums := nums[i:] //右边
		for j := 1; j < len(leftNums); j++ {
			for k := 1; k < len(rightNums); k++ {
				var tmp1 []int32
				tmp1 = append(tmp1, leftNums[0:j]...)
				tmp1 = append(tmp1, leftNums[j:]...)

				var tmp2 []int32
				tmp2 = append(tmp2, rightNums[0:k]...)
				tmp2 = append(tmp2, rightNums[k:]...)
				ans = append(ans, itoa(tmp1)+","+itoa(tmp2))
			}
		}
	}
	fmt.Println(nums)
	return ""
}
func main() {
	ambiguousCoordinates("123")
}
