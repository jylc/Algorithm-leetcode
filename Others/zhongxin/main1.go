package main

import (
	"fmt"
	"sort"
)

func algo(nums []int) int {
	fmt.Println(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	length := len(nums)
	if length <= 3 {
		return nums[0]
	}
	ans := 0
	fmt.Println(nums)
	for i := 0; i < length; {
		ans += nums[i]
		i += 2
		if (i+1)%length == 0 || i > length {
			break
		}
	}
	return ans
}
func main() {
	ans := algo([]int{33, 55, 88, 90, 101, 22})
	fmt.Println(ans)
}
