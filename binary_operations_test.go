package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestRightShift(t *testing.T) {
	result := RightShift(16, 2)
	expected := 4
	assert.Equal(t, result, expected)
}

func TestLeftShift(t *testing.T) {
	result := LeftShift(16, 2)
	expected := 64
	assert.Equal(t, result, expected)
}