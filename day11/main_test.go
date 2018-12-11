package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerLeve(t *testing.T) {
	fmt.Println("======= Running tests on power levels....")
	c := Cell{3, 5}
	assert.Equal(t, 4, calcPowerLevel(8, &c), "should be equal")
	c = Cell{122, 79}
	assert.Equal(t, -5, calcPowerLevel(57, &c), "should be equal")
	c = Cell{217, 196}
	assert.Equal(t, 0, calcPowerLevel(39, &c), "should be equal")
	c = Cell{101, 153}
	assert.Equal(t, 4, calcPowerLevel(71, &c), "should be equal")
}

func TestSquare(t *testing.T) {
	fmt.Println("======= Running tests on square totals....")
	createGrid(18)
	c := Cell{33, 45}
	assert.Equal(t, 29, calculateGridLevel(3, &c), "should be equal")

	createGrid(42)
	c = Cell{21, 61}
	assert.Equal(t, 30, calculateGridLevel(3, &c), "should be equal")

}

func TestSquareSizes(t *testing.T) {
	fmt.Println("======= Running tests on square sizes....")
	createGrid(18)
	c := Cell{90, 269}
	assert.Equal(t, 113, calculateGridLevel(16, &c), "should be equal")

	createGrid(42)
	c = Cell{232, 251}
	assert.Equal(t, 119, calculateGridLevel(12, &c), "should be equal")

}
