package ds

import (
	"errors"
	"log"
)

type Element[T any] struct {
	val  T
	prev *Element[T]
}

type Stack[T any] struct {
	top    *Element[T]
	length int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(v T) {
	newEl := Element[T]{val: v, prev: stack.top}
	stack.top = &newEl
	stack.length += 1
}

func (stack *Stack[T]) Pop() (T, bool) {
	if stack.length == 0 {
		var zeroValue T
		log.Fatal("Stack is empty")
		return zeroValue, false
	}

	val := stack.top.val
	stack.top = stack.top.prev
	stack.length -= 1

	return val, true
}

func (stack *Stack[T]) Peek() (T, error) {
	if stack.length == 0 {
		var zeroValue T
		return zeroValue, errors.New("Peek:Stack is empty")
	} else {
		return stack.top.val, nil
	}
}

func (stack *Stack[T]) Length() int {
	return stack.length
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.length == 0
}
