package dsa

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
