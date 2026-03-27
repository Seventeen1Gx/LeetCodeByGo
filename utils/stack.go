package utils

import "errors"

type MyStack struct {
	elments []int
}

func NewMyStack() *MyStack {
	return &MyStack{
		elments: make([]int, 0),
	}
}

func (s *MyStack) IsEmpty() bool {
	return len(s.elments) == 0
}

func (s *MyStack) Push(x int) {
	s.elments = append(s.elments, x)
}

func (s *MyStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	lastIdx := len(s.elments) - 1
	x := s.elments[lastIdx]
	s.elments = s.elments[0:lastIdx]
	return x, nil
}
