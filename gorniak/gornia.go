package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var curr = new(ListNode)
	result := curr

	for l1 != nil || l2 != nil {

		if l1 != nil && l2 != nil && l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
			curr = curr.Next
			continue
		}

		if l1 != nil && l2 != nil && l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
			curr = curr.Next
			continue
		}
		if l1 != nil && l2 == nil {
			curr.Next = l1
			l1 = l1.Next
			curr = curr.Next
			continue
		}
		if l2 != nil && l1 == nil {
			curr.Next = l2
			l2 = l2.Next
			curr = curr.Next
			continue
		}

	}
	return result.Next
}
