package main

//141. 环形链表
//https://leetcode-cn.com/problems/linked-list-cycle/
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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p, q := head, head
	for {
		p = p.Next
		if q == nil || q.Next == nil || p == nil {
			return false
		}
		q = q.Next.Next
		if p == q {
			return true
		}
	}
}
func main() {

}
