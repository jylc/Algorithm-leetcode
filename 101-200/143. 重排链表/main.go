package main

//143. 重排链表
//https://leetcode-cn.com/problems/reorder-list/

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

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var list []*ListNode
	p := head
	for p != nil {
		list = append(list, p)
		p = p.Next
	}

	i, j := 0, len(list)-1
	for i < j {
		list[i].Next = list[j]
		list[j].Next = list[i+1]
		i++
		j--
	}

	list[i].Next = nil

	/*p := head
	newHead := &ListNode{0, nil}
	cnt := 0
	for p != nil {
		cnt++
		tmp := &ListNode{p.Val,nil}
		p = p.Next
		tmp.Next = newHead.Next
		newHead.Next = tmp
	}

	fmt.Print(cnt)

	p = head
	q := newHead.Next
	for p != nil {
		t1 := q
		q = q.Next
		t2 := p.Next
		t1.Next = t2
		p.Next = t1
		p = t2
	}

	p=head
	for i:=0;i<cnt;i++{
		p=p.Next
	}
	p.Next=nil*/
}

func main() {

}
