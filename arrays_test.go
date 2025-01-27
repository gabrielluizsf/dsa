package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestContainsDuplicateII(t *testing.T) {
	arr := []int{1,2,3,1}
	ok := ContainsNearbyDuplicate(arr, 3)
	assert.True(t, ok)
	
	arr = []int{1, 0, 1, 1}
	ok = ContainsNearbyDuplicate(arr, 1)
	assert.True(t, ok)

	arr = []int{1,2,3,1,2,3}
	ok = ContainsNearbyDuplicate(arr, 2)
	assert.False(t, ok)
}

func TestFirstUniqueCharacterInAString(t *testing.T) {
	index := FirstUniqChar("leetcode")
	assert.Equal(t, index, 0)

	index = FirstUniqChar("loveleetcode")
	assert.Equal(t, index, 2)

	index = FirstUniqChar("aabb")
	assert.Equal(t, index, -1)
}

func TestMaximumLengthSubstringWithTwoOccurrences(t *testing.T) {
	result := MaximumLengthSubstring("bcbbbcba")
	assert.Equal(t, result, 4)
	result = MaximumLengthSubstring("aaaa")
	assert.Equal(t, result, 2)
}

func TestExponentialSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	index, steps := ExponentialSearch(arr, 32)
	assert.Equal(t, index, 31)
	assert.Equal(t, steps, 4)
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1}
	steps, index := BinarySearch(arr, 1)
	assert.Equal(t, steps, 1)
	assert.Equal(t, index, 0)

	arr = []int{1, 2, 3, 4, 5}
	steps, index = BinarySearch(arr, 3)
	assert.Equal(t, steps, 1)
	assert.Equal(t, index, 2)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	steps, index = BinarySearch(arr, 3)
	assert.Equal(t, steps, 2)
	assert.Equal(t, index, 2)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	steps, index = BinarySearch(arr, 3)
	assert.Equal(t, steps, 3)
	assert.Equal(t, index, 2)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	steps, index = BinarySearch(arr, 3)
	assert.Equal(t, steps, 4)
	assert.Equal(t, index, 2)
}

func TestReverseWords(t *testing.T) {
	words := "Hello World"
	result := ReverseWords(words)
	expected := "olleH dlroW"
	assert.Equal(t, result, expected)

	words = "Let's take LeetCode contest"
	result = ReverseWords(words)
	expected = "s'teL ekat edoCteeL tsetnoc"
	assert.Equal(t, result, expected)

	words = "Mr Ding"
	result = ReverseWords(words)
	expected = "rM gniD"
	assert.Equal(t, result, expected)
}
