package umath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, Max(1, 2, 3, 4, 9))
	assert.Equal(9.9, Max(9.8, 9.9, 2.5, 9.89))
}

func TestMin(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, Min(1, 2, 3, 4, 9))
	assert.Equal(9.79, Min(9.8, 9.9, 9.79, 9.89))
}
