package chapter_3

type Stack struct {
	stack []interface{}
}

func NewStack(initialCapacity int) *Stack {
	stack := make([]interface{}, 0, initialCapacity)
	return &Stack{stack}
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Push(v interface{}) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("Pop() called on empty stack")
	}
	n := len(s.stack)
	ret := s.stack[n-1]
	s.stack = s.stack[:n-1]
	return ret
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic("Peek() called on empty stack")
	}
	return s.stack[len(s.stack)-1]
}

func (s *Stack) len() int {
	return len(s.stack)
}
