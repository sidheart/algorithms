package chapter_2

func sumLists(a, b *ListNode) *ListNode {
	sum := &ListNode{}
	cur := sum
	var carry int
	for a != nil || b != nil {
		digit := carry
		if a != nil {
			digit += a.Value.(int)
			a = a .Next
		}
		if b != nil {
			digit += b.Value.(int)
			b = b.Next
		}
		carry = digit / 10
		digit %= 10
		cur.Next = &ListNode{digit, nil}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return sum.Next
}

func padZeros(a *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		a = &ListNode{0, a}
	}
	return a
}

func sumListsRec(a, b *ListNode) (xs *ListNode, carry int) {
	if a.Next == nil && b.Next == nil {
		digit := a.Value.(int) + b.Value.(int)
		carry = digit / 10
		digit %= 10
		return &ListNode{digit, nil}, carry
	}
	xs, carry = sumListsRec(a.Next, b.Next)
	digit := a.Value.(int) + b.Value.(int) + carry
	carry = digit / 10
	digit %= 10
	return &ListNode{digit, xs}, carry
}

func sumListsForward(a, b *ListNode) *ListNode {
	if a == nil && b == nil {
		return nil
	} else if a == nil {
		return b
	} else if b == nil {
		return a
	}
	m, n := a.Len(), b.Len()
	if m > n {
		b = padZeros(b, m - n)
	} else if n > m {
		a = padZeros(a, n - m)
	}
	sum, carry := sumListsRec(a, b)
	if carry > 0 {
		sum = &ListNode{carry, sum}
	}
	return sum
}
