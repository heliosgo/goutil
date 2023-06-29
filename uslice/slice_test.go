package uslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, Exists([]int{1, 2}, 1))
	assert.Equal(false, Exists([]int{1, 2}, 0))
}

func TestExistsf(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, Existsf([]int{1, 2}, func(x int) bool { return x == 1 }))
	assert.Equal(false, Existsf([]int{1, 2}, func(x int) bool { return x == 0 }))
}
