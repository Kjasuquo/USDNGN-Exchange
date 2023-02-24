package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeHash(t *testing.T) {
	got := ComputeHash("joseph", "12345678")
	want := ComputeHash("joseph", "12345678")

	assert.Equal(t, want, got)
}
