package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessage(t *testing.T) {
	messageExpected := "Aii Carammbaaa"

	result := GetMessage()

	assert.Equal(t, messageExpected, result)
}
