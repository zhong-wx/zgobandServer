package utils

import (
	"fmt"
	"testing"
)
type Pos struct {
	row int8
	column int8
}

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Count())
	v := stack.Pop()
	fmt.Println(v)
	v = stack.Pop()
	fmt.Println(v)
	fmt.Println(stack.Count())
}

func TestStack_ToSlice(t *testing.T) {
	stack := NewStack()
	stack.Push(Pos{1,2})
	stack.Push(Pos{1,123})
	s := stack.ToSlice()
	fmt.Println(s)
}
