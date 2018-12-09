package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	fmt.Println("======= Running tests on samples....")
	assert.Equal(t, 8317, solveProblem(10, 1618), "should be equal")
	assert.Equal(t, 146373, solveProblem(13, 7999), "should be equal")
	assert.Equal(t, 2764, solveProblem(17, 1104), "should be equal")
	assert.Equal(t, 54718, solveProblem(21, 6111), "should be equal")
	assert.Equal(t, 37305, solveProblem(30, 5807), "should be equal")
}
