package chapter_3

type MinStackElement struct {
	value    int
	localMin int
}

type MinStack struct {
	stack []MinStackElement
}

func NewMinStack(initialCapacity int) *MinStack {
	stack := make([]MinStackElement, 0, initialCapacity)
	return &MinStack{stack}
}

func (s *MinStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *MinStack) Peek() int {
	if s.IsEmpty() {
		panic("Peeked on empty stack")
	}
	return s.stack[len(s.stack) - 1].value
}

func (s *MinStack) Push(v int) {
	min := v
	n := len(s.stack)
	if !s.IsEmpty() {
		localMin := s.stack[n-1].localMin
		if localMin < min {
			min = localMin
		}
	}
	s.stack = append(s.stack, MinStackElement{v, min})
}

func (s *MinStack) Pop() int {
	if s.IsEmpty() {
		panic("Popped from empty stack")
	}
	n := len(s.stack)
	ret := s.stack[n-1].value
	s.stack = s.stack[:n-1]
	return ret
}

func (s *MinStack) Min() int {
	if s.IsEmpty() {
		panic("Min() called on empty stack")
	}
	return s.stack[len(s.stack)-1].localMin
}

func (s *MinStack) Size() int {
	return len(s.stack)
}