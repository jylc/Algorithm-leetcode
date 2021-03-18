package main

//92. 反转链表 II
//https://leetcode-cn.com/problems/reverse-linked-list-ii/
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

//切出范围内的链表，再拼接；时间O(N)，空间O(1)
func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummyNode.Next = head
	pre := dummyNode

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode := pre.Next
	curr := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverseLinkedList(leftNode)

	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
func main() {

}
