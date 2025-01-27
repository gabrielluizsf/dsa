package dsa

func MaximumLengthSubstring(s string) int {
	l, r := 0, 0
	maxLength := 1
	counter := make(map[byte]int)

	counter[s[0]] = 1

	for r < len(s)-1 {
		r++
		counter[s[r]]++
		for counter[s[r]] == 3 {
			counter[s[l]]--
			l++
		}
		if currentLength := r - l + 1; currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

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
