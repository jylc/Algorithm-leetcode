package main

//剑指 Offer 52. 两个链表的第一个公共节点
//https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
func main() {
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 5}
	headA := node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = nil

	node4 := &ListNode{Val: 2}
	node5 := &ListNode{Val: 1}
	node6 := &ListNode{Val: 5}
	headB := node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil

	getIntersectionNode(headA, headB)
}
