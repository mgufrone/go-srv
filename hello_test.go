package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello_Greet(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in  string
		out string
	}{
		{
			"World",
			"Hello, World",
		},
		{
			"dan",
			"Hello, Dan",
		},
		{
			"dan1",
			"Hello, Dan1. A weird name, but okay.",
		},
	}
	h := &Hello{}
	for _, c := range testCases {
		result := h.Greet(c.in)
		assert.Equal(t, c.out, result)
	}
}
