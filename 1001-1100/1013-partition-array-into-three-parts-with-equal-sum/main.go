package main

import "fmt"

//https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
//1013. 将数组分成和相等的三个部分
func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	ans := false
	if sum%3 != 0 {
		return false
	} else {
		average := sum / 3
		i, j := 0, len(arr)-1
		left, right := arr[i], arr[j]
		for i < j {
			if left != average {
				i++
				left += arr[i]
			}

			if right != average {
				j--
				right += arr[j]
			}

			if right == average && left == average {
				ans = true
				break
			}
		}
		return ans && j-i > 1
	}
}
func main() {
	/*arr := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}*/
	arr := []int{1, -1, 1, -1}
	ans := canThreePartsEqualSum(arr)
	fmt.Println(ans)
}
