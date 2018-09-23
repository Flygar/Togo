package main

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//dummy
	result := &ListNode{0, nil}
	root := result

	//Not required but I just do it anyway
	tl1, tl2 := l1, l2
	for tl1 != nil && tl2 != nil {
		if tl1.Val < tl2.Val {
			result.Next = &ListNode{tl1.Val, nil}
			tl1 = tl1.Next
		} else {
			result.Next = &ListNode{tl2.Val, nil}
			tl2 = tl2.Next
		}
		result = result.Next
	}
	if tl1 == nil {
		result.Next = tl2
	} else {
		result.Next = tl1
	}
	return root.Next

}
