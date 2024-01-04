func oddEvenList(head *ListNode) *ListNode {
	if head==nil{
        return nil
    }
    if head.Next==nil{
        return head
    }

    odd := &ListNode{}
	even := &ListNode{}
	evenT := &ListNode{}
	oddT := &ListNode{}
	count := 1

	for head != nil {
		if count%2 == 0 {
			if even.Val == 0 && even.Next == nil {
				even = head
				evenT = head
			} else {
				evenT.Next = head
				evenT = evenT.Next
			}
			head = head.Next
		} else {
			if odd.Val == 0 && odd.Next == nil {
				odd = head
				oddT = head
			} else {
				oddT.Next = head
				oddT = oddT.Next
			}
			head = head.Next
		}
		count++
	}

    

	if odd == nil {
		return even
	}
	if even == nil {
		return odd
	}

	evenT.Next = nil
	oddT.Next = even
	return odd
}
