package main

//142. 环形链表 II
//https://leetcode-cn.com/problems/linked-list-cycle-ii/
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

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head
	for {
		p = p.Next
		if q == nil || q.Next == nil || p == nil {
			return nil
		}
		q = q.Next.Next
		if p == q {
			p = head
			for p != q {
				p = p.Next
				q = q.Next
			}
			return p
		}
	}
}
func main() {

}
