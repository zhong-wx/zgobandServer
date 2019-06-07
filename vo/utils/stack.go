package utils

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list: list.New()}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e == nil {
		return nil
	}
	stack.list.Remove(e)
	return e.Value
}

func (stack *Stack) Back() interface{} {
	if(stack.list.Back() == nil) {
		return nil
	}
	return stack.list.Back().Value
}

func (stack *Stack) Count() int {
	return stack.list.Len()
}

func (stack *Stack) ToSlice() []interface{} {
	p := stack.list.Front()
	s := []interface{}{}
	for p != nil {
		s = append(s, p.Value)
		p = p.Next()
	}
	return s
}