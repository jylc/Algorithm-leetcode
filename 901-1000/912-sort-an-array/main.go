package main

import "fmt"

//https://leetcode-cn.com/problems/sort-an-array/
//912. 排序数组
//冒泡排序
func sortArray1(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

//快速排序
func quickSort(nums []int, left, right int) {
	tmp := nums[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && nums[j] >= tmp {
			j--
		}
		if j >= p {
			nums[p] = nums[j]
			p = j
		}

		for i <= p && nums[i] <= tmp {
			i++
		}
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
	}
	nums[p] = tmp
	if p-left > 1 {
		quickSort(nums, left, p-1)
	}
	if right-p > 1 {
		quickSort(nums, p+1, right)
	}

}
func sortArray(nums []int) []int {
	left, right := 0, len(nums)-1
	quickSort(nums, left, right)
	return nums
}

func main() {
	//nums := []int{5, 2, 3, 1}
	nums := []int{5, 1, 1, 2, 0, 0}
	ans := sortArray(nums)
	fmt.Println(ans)
}
