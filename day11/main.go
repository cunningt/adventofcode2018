package main

import "fmt"

type Cell struct {
	x int
	y int
}

func calcPowerLevel(input int, c *Cell) int {
	rackID := c.x + 10
	p1 := rackID * c.y
	p2 := p1 + input
	p3 := p2 * rackID

	p4 := p3 / 100
	p5 := p4 % 10
	return (p5 - 5)
}

var grid map[int]map[int]int

func createGrid(input int) {
	grid = make(map[int]map[int]int)
	for x := 1; x <= 300; x++ {
		grid[x] = make(map[int]int)
		for y := 1; y <= 300; y++ {
			c := Cell{x, y}
			grid[x][y] = calcPowerLevel(input, &c)
		}
	}
}

func calculateGridLevel(size int, cell *Cell) int {
	var total int = 0
	for x := cell.x; x < cell.x+size; x++ {
		for y := cell.y; y < cell.y+size; y++ {
			total += grid[x][y]
		}
	}
	return total
}

func main() {
	createGrid(7857)

	// Solve for Part 1
	max := -99999991999
	maxx := -1
	maxy := -1

	// Use size of 3, make sure we don't compute any sub-grids that aren't 3x3
	for x := 1; x <= (300 - 3 + 1); x++ {
		for y := 1; y <= (300 - 3 + 1); y++ {
			c := Cell{x, y}
			level := calculateGridLevel(3, &c)
			if max < level {
				max = level
				maxx = x
				maxy = y
			}
		}
	}

	fmt.Printf("======Part One =======\n")
	fmt.Printf("Max level of %d found at (%d,%d)\n", max, maxx, maxy)

	// Solve for Part 2
	max = -99999991999
	maxx = -1
	maxy = -1
	maxsize := -1

	// This takes a super-long time to complete
	// There must be a better way to do this
	for size := 1; size <= 300; size++ {
		fmt.Printf("Computing sizes of %d\n", size)
		for x := 1; x <= (300 - size + 1); x++ {
			for y := 1; y <= (300 - size + 1); y++ {
				c := Cell{x, y}
				level := calculateGridLevel(size, &c)
				if max < level {
					max = level
					maxx = x
					maxy = y
					maxsize = size
				}
			}
		}
		fmt.Printf("Max level of %d found at (%d,%d,%d)\n", max, maxx, maxy, maxsize)
	}
	fmt.Printf("======Part Two =======\n")
	fmt.Printf("Max level of %d found at (%d,%d,%d)\n", max, maxx, maxy, maxsize)

}
