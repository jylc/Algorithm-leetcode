package main

// ListNode /**
//https://leetcode.cn/problems/swap-nodes-in-pairs/
//24. 两两交换链表中的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head

	return head
}
func main() {

}
