package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.AddToFront(3)
	dll.AddToFront(2)
	dll.AddToFront(1)
	dll.AddToEnd(4)
	dll.AddToEnd(5)

	result := dll.DisplayForward()
	expected := "[1, 2, 3, 4, 5]"
	assert.Equal(t, result, expected)

	result = dll.DisplayBackward()
	expected = "[5, 4, 3, 2, 1]"
	assert.Equal(t, result, expected)

	removed := dll.RemoveFromFront()
	assert.Equal(t, removed, any(1))

	removed = dll.RemoveFromEnd()
	assert.Equal(t, removed, any(5))

	result = dll.DisplayForward()
	expected = "[2, 3, 4]"
	assert.Equal(t, result, expected)

	result = dll.DisplayBackward()
	expected = "[4, 3, 2]"
	assert.Equal(t, result, expected)
}

func TestReverseLinkedList(t *testing.T) {
	ll := &ListNode{Val: 1}
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll = ReverseList(ll)
	assert.Equal(t, ll.String(), "5 -> 4 -> 3 -> 2 -> 1 -> nil")
}

func TestMiddleOfTheLinkedList(t *testing.T) {
	ll := &ListNode{Val: 1}
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll = MiddleNode(ll)
	assert.Equal(t, ll.Val, 3)
}

func TestLinkedListCycle(t *testing.T) {
    var head *ListNode
    assert.False(t, HasCycle(head))

    head = &ListNode{Val: 1}
    assert.False(t, HasCycle(head))

    node1 := &ListNode{Val: 1}
    node2 := &ListNode{Val: 2}
    node3 := &ListNode{Val: 3}
    node1.Next = node2
    node2.Next = node3
    node3.Next = node1 
    assert.True(t, HasCycle(node1))

    node1 = &ListNode{Val: 1}
    node2 = &ListNode{Val: 2}
    node3 = &ListNode{Val: 3}
    node1.Next = node2
    node2.Next = node3
    assert.False(t, HasCycle(node1))
}
