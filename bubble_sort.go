package dsa

import "cmp"

func BubbleSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	for i := range n {
		for j := 0; j < n - i - 1; j++ {
			if max(arr[j], arr[j + 1])  == arr[j] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}