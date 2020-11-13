package main

//328. 奇偶链表
//https://leetcode-cn.com/problems/odd-even-linked-list/
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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p, q, tmp := head, head.Next, head.Next

	for p.Next != nil && q.Next != nil {
		p.Next = q.Next
		p = p.Next
		q.Next = p.Next
		q = q.Next
	}

	p.Next = tmp
	return head
}
func main() {

}
