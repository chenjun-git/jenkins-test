package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	ret := testFunc()
	assert := assert.New(t)
	assert.Equal(1, ret)
}