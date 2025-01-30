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

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(val int) {
	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Val: val}
}

func (l *ListNode) String() string {
	var result []string
	current := l
	for current != nil {
		result = append(result, fmt.Sprintf("%d", current.Val))
		current = current.Next
	}
	return strings.Join(result, " -> ") + " -> nil"
}

func ReverseList(head *ListNode) *ListNode {
	var newList *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = newList
		newList = head
		head = nextNode
	}
	return newList
}

func MiddleNode(head *ListNode) *ListNode {
	aHead := head
	for aHead != nil && aHead.Next != nil {
		aHead = aHead.Next.Next
		head = head.Next
	}
	return head
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}


func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	current := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return head
}
