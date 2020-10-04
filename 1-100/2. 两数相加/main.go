package main

//2. 两数相加
//https://leetcode-cn.com/problems/add-two-numbers/
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

//go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{
		Val:  0,
		Next: nil,
	}
	head := list.Next
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		tmp = tmp / 10
		list = list.Next
	}
	return head
}
func main() {

}
