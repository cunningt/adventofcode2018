package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	fmt.Println("======= Running tests on Part One samples....")
	circle := solveProblem(2028, "37")

	assert.Equal(t, "5158916779", findFirstAnswer(9, 10, circle), "should be equal")
	assert.Equal(t, "0124515891", findFirstAnswer(5, 10, circle), "should be equal")
	assert.Equal(t, "9251071085", findFirstAnswer(18, 10, circle), "should be equal")
	assert.Equal(t, "5941429882", findFirstAnswer(2018, 10, circle), "should be equal")
}

func TestPartTwo(t *testing.T) {
	fmt.Println("======= Running tests on Part Two samples....")
	assert.Equal(t, 9, solveSecondProblem("37", "515891"), "should be equal")
	assert.Equal(t, 5, solveSecondProblem("37", "012451"), "should be equal")
	assert.Equal(t, 18, solveSecondProblem("37", "925107"), "should be equal")
	assert.Equal(t, 2018, solveSecondProblem("37", "594142"), "should be equal")
}
