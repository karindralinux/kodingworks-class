package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsoluteWithNegativeNumber(t *testing.T) {
	number := Absolute(-7.5)
	assert.Equal(t, 7.5, number)
}

func TestAbsoluteWithPositiveNumber(t *testing.T) {
	number := Absolute(5)

	assert.Equal(t, float64(5), number)
}
