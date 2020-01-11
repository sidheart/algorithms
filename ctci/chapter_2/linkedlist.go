package chapter_2

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func MakeLinkedList(e []interface{}) (head *ListNode) {
	if len(e) == 0 {
		return head
	}
	head = &ListNode{e[0], nil}
	curr := head
	for i := 1; i < len(e); i++ {
		curr.Next = &ListNode{e[i], nil}
		curr = curr.Next
	}
	return head
}

func (l *ListNode) Equals(other *ListNode) bool {
	if l == nil && other == nil {
		return true
	}
	if l == nil || other == nil {
		return false
	}
	for l != nil || other != nil {
		if l == nil || other == nil || l.Value != other.Value {
			return false
		}
		l, other = l.Next, other.Next
	}
	return true
}

func (l *ListNode) Len() (len int) {
	for l != nil {
		len++
		l = l.Next
	}
	return len
}