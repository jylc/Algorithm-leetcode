package main

//234. 回文链表
//https://leetcode-cn.com/problems/palindrome-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	p := head
	count := 0
	var arr []int
	for p != nil {
		count++
		arr = append(arr, p.Val)
		p = p.Next
	}
	if arr == nil {
		return true
	}
	i, j := 0, count-1
	for i <= j {
		if arr[i] == arr[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
func main() {

}
