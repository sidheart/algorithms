package chapter_3

type StackQueue struct {
	new, old *Stack
}

func NewStackQueue(capacity int) StackQueue {
	return StackQueue{NewStack(capacity), NewStack(capacity)}
}

func (q StackQueue) IsEmpty() bool {
	return q.new.IsEmpty() && q.old.IsEmpty()
}

func (q StackQueue) Add(v interface{}) {
	q.new.Push(v)
}

func (q StackQueue) shiftStacks() {
	if q.old.IsEmpty() {
		for !q.new.IsEmpty() {
			q.old.Push(q.new.Pop())
		}
	}
}

func (q StackQueue) Remove() interface{} {
	if q.IsEmpty() {
		panic("Remove() called on empty queue")
	}
	q.shiftStacks()
	return q.old.Pop()
}

func (q StackQueue) Peek() interface{} {
	if q.IsEmpty() {
		panic("Peek() called on empty queue")
	}
	q.shiftStacks()
	return q.old.Peek()
}