package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestReverseWords(t *testing.T) {
	words := "Hello World"
	result := ReverseWords(words)
	expected := "olleH dlroW"
	assert.Equal(t, result, expected)
	
	words = "Let's take LeetCode contest"
	result  = ReverseWords(words)
	expected = "s'teL ekat edoCteeL tsetnoc"
	assert.Equal(t, result, expected)

	words = "Mr Ding"
	result = ReverseWords(words)
	expected = "rM gniD"
	assert.Equal(t, result, expected)
}
