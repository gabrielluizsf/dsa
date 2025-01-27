package dsa

func BinarySearch(nums []int, n int) (steps, index int) {
	low, high := 0, len(nums)-1
	for low <= high {
		steps++
		middle := (low + high) / 2
		if nums[middle] == n {
			index = middle
			return
		} else if nums[middle] < n {
			low = middle + 1
		} else {
			high = middle
		}
	}
	index = -1
	return
}

func ReverseWords(s string) string {
	chars := []rune(s)
	start := 0

	for i := 0; i <= len(chars); i++ {
		if i == len(chars) || chars[i] == ' ' {
			reverse(chars, start, i-1)
			start = i + 1
		}
	}

	return string(chars)
}

func reverse(chars []rune, start, end int) {
	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
}
