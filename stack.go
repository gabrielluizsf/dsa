package dsa

import "errors"

var (
	ErrStackIsEmpty  = errors.New("stack is empty")
	ErrStackOverflow = errors.New("stack overflow")
)

// Stack is a generic interface representing a stack data structure.
// It supports common stack operations such as push, pop, and peek.
type Stack[T any] interface {
    // Push adds a new element to the top of the stack.
    // Returns an error if the operation fails.
    Push(value T) error

    // Pop removes and returns the top element of the stack.
    // Returns an error if the stack is empty.
    Pop() (val T, err error)

    // Peek returns the top element of the stack without removing it.
    // Returns an error if the stack is empty.
    Peek() (val T, err error)

    // Size returns the number of elements currently in the stack.
    Size() int
}

type StackNode[T any] struct {
	Value T
	Next  *StackNode[T]
}

type LLStack[T any] struct {
	top  *StackNode[T]
	size int
}

func NewStack[T any]() Stack[T] {
	return &LLStack[T]{}
}

func (s *LLStack[T]) Push(value T) error {
	newNode := &StackNode[T]{Value: value}
	newNode.Next = s.top
	s.top = newNode
	s.size++
	return nil
}

func (s *LLStack[T]) Pop() (val T, err error) {
	if s.top == nil {
		err = ErrStackIsEmpty
		return
	}

	poppedNode := s.top
	s.top = poppedNode.Next
	s.size--
	val = poppedNode.Value
	return
}

func (s *LLStack[T]) Peek() (val T, err error) {
	if s.top == nil {
		err = ErrStackIsEmpty
		return
	}
	val, err = s.top.Value, nil
	return
}

func (s *LLStack[T]) Size() int {
	return s.size
}

type ArrStack[T any] struct {
	items     []T
	maxLength int
}

func NewArrStack[T any](maxLength int) Stack[T] {
	return &ArrStack[T]{
		items:     make([]T, 0, maxLength),
		maxLength: maxLength,
	}
}

func (s *ArrStack[T]) Push(value T) error {
	if len(s.items) >= s.maxLength {
		return ErrStackOverflow
	}
	s.items = append(s.items, value)
	return nil
}

func (s *ArrStack[T]) Pop() (val T, err error) {
	if len(s.items) == 0 {
		err = ErrStackIsEmpty
		return
	}

	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	val, err = value, nil
	return
}

func (s *ArrStack[T]) Peek() (val T, err error) {
	if len(s.items) == 0 {
		err = ErrStackIsEmpty
		return
	}
	val, err = s.items[len(s.items)-1], nil
	return
}

func (s *ArrStack[T]) Size() int {
	return len(s.items)
}
