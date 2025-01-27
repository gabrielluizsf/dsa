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

func binarySearch(nums []int, n, lo, hi int) (index, steps int) {
	if hi == -1 {
		hi = len(nums) - 1
	}

	for lo <= hi {
		steps++
		mid := (lo + hi) / 2
		if nums[mid] == n {
			index = mid
			return
		} else if nums[mid] < n {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if lo < len(nums) && nums[lo] == n {
		index = lo
		return
	}

	index = -1
	return
}

func ExponentialSearch(arr []int, target int) (index, steps int) {
	if arr[0] == target {
		index = 0
		steps = 0
		return
	}

	n := len(arr)
	i := 1
	for i < n && arr[i] < target {
		i *= 2
	}

	if i < n && arr[i] == target {
		index = i
		steps = 0
		return
	}
	index, steps = binarySearch(arr, target, i/2, min(i, n-1))
	return
}

func BinarySearch(nums []int, n int) (steps, index int) {
	low, high := 0, len(nums)-1
	index, steps = binarySearch(nums, n, low, high)
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
