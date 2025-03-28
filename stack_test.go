package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestArrStack(t *testing.T) {
	stack := NewArrStack[int](3)
	testStack(t, stack)
	stack = NewStack[int]()
	testStack(t, stack)
}

func testStack(t assert.T, stack Stack[int]) {
	_, err := stack.Pop()
	assert.Equal(t, err, ErrStackIsEmpty)

	err = stack.Push(1)
	assert.NoError(t, err)

	err = stack.Push(2)
	assert.NoError(t, err)

	err = stack.Push(3)
	assert.NoError(t, err)

	top, _ := stack.Peek()
	assert.Equal(t, top, 3)

	val, _ := stack.Pop()
	assert.Equal(t, val, 3)

	val, _ = stack.Pop()
	assert.Equal(t, val, 2)

	val, _ = stack.Pop()
	assert.Equal(t, val, 1)

	_, err = stack.Pop()
	assert.Equal(t, err, ErrStackIsEmpty)
}
