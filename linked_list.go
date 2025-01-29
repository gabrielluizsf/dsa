package dsa

import (
	"fmt"
	"strings"
)

type LinkedListNode struct {
	Value    any
	Next     *LinkedListNode
	Previous *LinkedListNode
}

type DoublyLinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
}

func (ll *DoublyLinkedList) AddToFront(value any) {
	node := &LinkedListNode{Value: value}
	node.Next = ll.Head
	if ll.Head != nil {
		ll.Head.Previous = node
	} else {
		ll.Tail = node
	}
	ll.Head = node
}

func (ll *DoublyLinkedList) AddToEnd(value any) {
	node := &LinkedListNode{Value: value}
	node.Previous = ll.Tail
	if ll.Head != nil {
		ll.Tail.Next = node
	} else {
		ll.Head = node
	}
	ll.Tail = node
}

func (ll *DoublyLinkedList) RemoveFromFront() any {
	if ll.Head == nil {
		return nil
	}
	removedValue := ll.Head.Value
	ll.Head = ll.Head.Next
	if ll.Head != nil {
		ll.Head.Previous = nil
	} else {
		ll.Tail = nil
	}
	return removedValue
}

func (ll *DoublyLinkedList) RemoveFromEnd() any {
	if ll.Tail == nil {
		return nil
	}
	removedValue := ll.Tail.Value
	ll.Tail = ll.Tail.Previous
	if ll.Tail != nil {
		ll.Tail.Next = nil
	} else {
		ll.Head = nil
	}
	return removedValue
}

func (ll *DoublyLinkedList) DisplayForward() string {
	var values []string
	for current := ll.Head; current != nil; current = current.Next {
		values = append(values, fmt.Sprintf("%v", current.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func (ll *DoublyLinkedList) DisplayBackward() string {
	var values []string
	for current := ll.Tail; current != nil; current = current.Previous {
		values = append(values, fmt.Sprintf("%v", current.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}
