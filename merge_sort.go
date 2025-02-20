package dsa

func MergeSort(head *ListNode) *ListNode {
	sorter := &mergeSorter{head: head, lists: make([]*ListNode, 0)}
	if head == nil || head.Next == nil {
		return head
	}
	sorter.createSubLists()
	return sorter.mergeLists()
}

type mergeSorter struct {
	head  *ListNode
	lists []*ListNode
}

func (s *mergeSorter) createSubLists() {
	for curr := s.head; curr != nil; {
		next := curr.Next
		curr.Next = nil
		s.lists = append(s.lists, curr)
		curr = next
	}
}

func (s mergeSorter) mergeLists() *ListNode {
	if len(s.lists) == 0 {
		return nil
	}
	for len(s.lists) > 1 {
		var merged []*ListNode
		for i := 0; i < len(s.lists)-1; i += 2 {
			merged = append(merged, MergeTwoLists(s.lists[i], s.lists[i+1]))
		}
		if len(s.lists)%2 == 1 {
			merged = append(merged, s.lists[len(s.lists)-1])
		}
		s.lists = merged
	}
	return s.lists[0]
}
