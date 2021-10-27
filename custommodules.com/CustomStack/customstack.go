package CustomStack

import (
	"fmt"
)

type IStack interface {
	Push(interface{})
	Pop()
	String()
}

type stackNode struct {
	value    interface{}
	previous *stackNode
}

type stack struct {
	top *stackNode
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) Push(newValue interface{}) {
	if s.top == nil {
		newNode := stackNode{value: newValue}
		s.top = &newNode
		return
	}

	newNode := stackNode{value: newValue, previous: s.top}
	s.top = &newNode
}

func (s *stack) Pop() stackNode {
	poppedNode := s.top
	s.top = s.top.previous
	return *poppedNode
}

func (s *stack) String() {
	fmt.Println(s.top.StringifyNode(0))
}

func (s *stackNode) StringifyNode(index int) string {
	if s == nil {
		return ""
	}
	nodeInfo := fmt.Sprintf("[I: %v, Val: %v],", index, s.value)
	return nodeInfo + s.previous.StringifyNode(index+1)
}
