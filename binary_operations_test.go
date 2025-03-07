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

func TestAND(t *testing.T) {
	result := AND(16, 2)
	expected := 0
	assert.Equal(t, result, expected)
}

func TestOR(t *testing.T) {
	result := OR(16, 2)
	expected := 18
	assert.Equal(t, result, expected)
}

func TestNOT(t *testing.T) {
	result := NOT(5)
	expected := -6
	assert.Equal(t, result, expected)
}

func TestXOR(t *testing.T) {
	result := XOR(1, 1)
	expected := 0
	assert.Equal(t, result, expected)
}