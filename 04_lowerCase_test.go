package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIdToLowerCase(t *testing.T) {
	newId := LowerCase(xiaomiId)
	assert.Equal(t, newId, expected)
}
