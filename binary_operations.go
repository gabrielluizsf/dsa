package dsa

func LeftShift(n int, k uint) int {
	return n << k
}

func RightShift(n int, k uint) int {
	return n >> k
}

func AND(n, m int) int {
	return n & m
}

func OR(n, m int) int {
	return n | m
}

func NOT(n int) int {
	return ^n
}

func XOR(n, m int) int {
	return n ^ m
}

//https://leetcode.com/problems/missing-number/
func MissingNumber(nums []int) int {
	x := 0
	for _, num := range nums {
		x ^= num
	}
	for i := range len(nums)+1 {
		x ^= i
	}
	return x
}

//https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func NumberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		steps++
	}
	return steps
}