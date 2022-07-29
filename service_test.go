package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessage(t *testing.T) {
	messageExpected := "Itss Alaiiiiiivvv!"

	result := GetMessage()

	assert.Equal(t, result, messageExpected)
}
