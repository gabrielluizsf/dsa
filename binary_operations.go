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
