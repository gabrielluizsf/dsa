package dsa

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestDeque(t *testing.T) {
	dq := NewDeque()
	dq.Append(10)
	dq.Append(20)
	v, ok := dq.Popleft()
	assert.True(t, ok)
	assert.Equal(t, v, any(10))
	assert.Equal(t, dq.Len(), 1)
	v, ok = dq.Popleft()
	assert.True(t, ok)
	assert.Equal(t, v, any(20))
	assert.Equal(t, dq.Len(), 0)
	v, ok = dq.Popleft()
	assert.False(t, ok)
	assert.Equal(t, v, any(0))
}
