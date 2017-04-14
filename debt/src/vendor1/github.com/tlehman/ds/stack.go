package ds

import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{l: list.New()}
}

func (s *Stack) Push(v interface{}) {
	s.l.PushFront(v)
}

func (s *Stack) Pop() interface{} {
	e := s.l.Front()
	v := e.Value
	s.l.Remove(e)
	return v
}

func (s *Stack) Len() int {
	return s.l.Len()
}
