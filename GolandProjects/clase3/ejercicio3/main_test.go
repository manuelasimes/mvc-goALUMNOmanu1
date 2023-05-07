package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivison(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input1   int
		input2   int
		expected int
		err      error
	}{
		{2, 1, 2, nil},
		{9, 3, 3, nil},
		{10, 0, 0, errors.New("no puedo dividir por cero")},
	}

	for _, test := range tests {
		result, err := Division(test.input1, test.input2)
		assert.Equal(result, test.expected)
		assert.Equal(err, test.err)
	}
}
