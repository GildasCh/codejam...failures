package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBest(t *testing.T) {
	b, u := best(5)
	assert.Equal(t, 2, b)
	assert.Equal(t, 2, u)

	b, u = best(12)
	assert.Equal(t, 4, b)
	assert.Equal(t, 2, u)
}
