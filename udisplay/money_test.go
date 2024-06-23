package udisplay

import (
	"goutil/utype"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperRMB(t *testing.T) {
	assert := assert.New(t)
	type tcase[T utype.Number | string] struct {
		amount   T
		expected string
	}
	cases := []tcase[string]{
		{"¥1", "壹元整"},
		{"¥1.01", "壹元零角壹分"},
		{"¥1001100110011001", "壹仟零壹万亿壹仟零壹亿壹仟零壹万壹仟零壹元整"},
	}
	for _, c := range cases {
		actual, _ := UpperRMB(c.amount)
		assert.Equal(c.expected, actual)
	}
}
