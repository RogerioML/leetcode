package lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	fmt.Println(list1, list2)
	return &ListNode{}
}
