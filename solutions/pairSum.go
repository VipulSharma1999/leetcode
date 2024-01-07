package main

func pairSum(head *ListNode) int {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	node1 := reverse(slow)
	node2 := head
	res := 0

	for node1 != nil {
		res = maxx(res, node1.Val+node2.Val)
		node1 = node1.Next
		node2 = node2.Next
	}
	return res

}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
