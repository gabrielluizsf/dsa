package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	BubbleSort(arr)
	assert.Equal(t, arr, []int{11, 12, 22, 25, 34, 64, 90})

	strArr := []string{"banana", "apple", "orange", "grape"}
	BubbleSort(strArr)
	assert.Equal(t, strArr, []string{"apple", "banana", "grape", "orange"})

	floatArr := []float64{64.4872894172478, 34.213, 25.41743281478327194732714723198477, 12.381318731, 22.5, 11.2, 90.173189378127371}
	BubbleSort(floatArr)
	sorted := []float64{11.2, 12.381318731, 22.5, 25.41743281478327194732714723198477, 34.213, 64.4872894172478, 90.173189378127371}
	assert.Equal(t, floatArr, sorted)
}
