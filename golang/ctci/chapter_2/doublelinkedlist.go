package chapter_2

type DoubleListNode struct {
	Value interface{}
	Prev, Next *DoubleListNode
}

func MakeDoublyLinkedList(e []interface{}) *DoubleListNode {
	n := len(e)
	if n == 0 {
		return nil
	}
	head := &DoubleListNode{e[0], nil, nil}
	curr := head
	for i := 1; i < n; i++ {
		curr.Next = &DoubleListNode{e[i], curr, nil}
		curr = curr.Next
	}
	return head
}
