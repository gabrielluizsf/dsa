package dsa

import (
	"math/rand"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5}
	QuickSort(arr)
	assert.Equal(t, arr, []int{1, 5, 7, 8, 9, 10})

	strArr := []string{"banana", "apple", "orange", "grape"}
	QuickSort(strArr)
	assert.Equal(t, strArr, []string{"apple", "banana", "grape", "orange"})

	floatArr := []float64{64.4872894172478, 34.213, 25.41743281478327194732714723198477, 12.381318731, 22.5, 11.2, 90.173189378127371}
	QuickSort(floatArr)
	sorted := []float64{11.2, 12.381318731, 22.5, 25.41743281478327194732714723198477, 34.213, 64.4872894172478, 90.173189378127371}
	assert.Equal(t, floatArr, sorted)
}

func BenchmarkQuickSortRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateRandomArray(1000)
		QuickSort(arr, true)
	}
}

func BenchmarkQuickSortIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateRandomArray(1000)
		QuickSort(arr)
	}
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(10000)
	}
	return arr
}
