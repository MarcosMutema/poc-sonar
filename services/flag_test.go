package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlternateMessage(t *testing.T) {
	expected := true

	result := AlternateMessage()

	assert.Equal(t, expected, result)
}
