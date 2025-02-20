package dsa

import "cmp"

type quickSorter[T cmp.Ordered] struct {
	arr []T
}

func QuickSort[T cmp.Ordered](arr []T, recursion ...bool) {
	sorter := quickSorter[T]{arr}
	if len(recursion) > 0 && recursion[len(recursion)-1] {
		sorter.recursive(0, len(arr)-1)
	} else {
		sorter.iterative(0, len(arr)-1)
	}
}

func (s quickSorter[T]) partition(low, high int) int {
	arr := s.arr
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {

		if min(arr[j], pivot) == arr[j] || arr[j] == pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (s quickSorter[T]) recursive(low, high int) {
	if low < high {
		pi := s.partition(low, high)
		s.recursive(low, pi-1)
		s.recursive(pi+1, high)
	}
}

func (s quickSorter[T]) iterative(low, high int) {
	stack := make([]int, 0)
	stack = append(stack, low, high)

	for len(stack) > 0 {
		high, stack = stack[len(stack)-1], stack[:len(stack)-1]
		low, stack = stack[len(stack)-1], stack[:len(stack)-1]

		pi := s.partition(low, high)

		if pi-1 > low {
			stack = append(stack, low, pi-1)
		}
		if pi+1 < high {
			stack = append(stack, pi+1, high)
		}
	}
}
