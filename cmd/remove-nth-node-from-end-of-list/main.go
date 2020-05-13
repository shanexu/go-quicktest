package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(head *ListNode) {
	for h := head; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	node = removeNthFromEnd(node, 3)
	printListNode(node)

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	a := head
	m := 0
	for b := head; b.Next != nil; b = b.Next {
		if m == n {
			a = a.Next
		} else {
			m++
		}
	}
	if m == n-1 {
		return a.Next
	}
	a.Next = a.Next.Next
	return head
}
