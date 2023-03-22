package main

import "leetcode/lists"

func main() {
	list1 := &lists.ListNode{Val: 1, Next: &lists.ListNode{Val: 2, Next: &lists.ListNode{Val: 4, Next: nil}}}
	list2 := &lists.ListNode{Val: 1, Next: &lists.ListNode{Val: 3, Next: &lists.ListNode{Val: 4, Next: nil}}}

	lists.MergeTwoLists(list1, list2)
	
}
