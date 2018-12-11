package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	x int
	y int
}

type Point struct {
	initialposition Position
	currentposition Position
	velocity        Position
}

func printPoint(p *Point) {
	fmt.Printf("Current [%d, %d] Velocity [%d, %d]\n",
		p.currentposition.x, p.currentposition.y, p.velocity.x, p.velocity.y)
}

func parsePoints(point string) *Point {
	p := new(Point)
	re := regexp.MustCompile("position=<\\s*([-0-9]+),\\s*([-0-9]+)> velocity=<\\s*([-0-9]+),\\s*([-0-9]+)>")

	if re.MatchString(point) {
		matches := re.FindAllStringSubmatch(point, -1)

		p.initialposition.x, _ = strconv.Atoi(matches[0][1])
		p.initialposition.y, _ = strconv.Atoi(matches[0][2])
		p.currentposition.x, _ = strconv.Atoi(matches[0][1])
		p.currentposition.y, _ = strconv.Atoi(matches[0][2])
		p.velocity.x, _ = strconv.Atoi(matches[0][3])
		p.velocity.y, _ = strconv.Atoi(matches[0][4])
	}
	return p
}

// Update the current position with the velocity
func updatePoint(p *Point) {
	p.currentposition.x += p.velocity.x
	p.currentposition.y += p.velocity.y
}

func simulate(seconds int, l *list.List) {

	// Defaults
	maxx := -999999
	maxy := -999999
	minx := 999999
	miny := 999999

	for m := l.Front(); m != nil; m = m.Next() {
		point := m.Value.(*Point)
		updatePoint(point)

		// Set minimums and maximums
		if minx > point.currentposition.x {
			minx = point.currentposition.x
		}
		if miny > point.currentposition.y {
			miny = point.currentposition.y
		}

		if maxx < point.currentposition.x {
			maxx = point.currentposition.x
		}

		if maxy < point.currentposition.y {
			maxy = point.currentposition.y
		}
	}

	found := true

	// Figure out if the points are real close together
	// This is not elegant, but given the letters are formed
	// by contiguous '#', I think it works
	for x := l.Front(); x != nil; x = x.Next() {
		point := x.Value.(*Point)
		for y := l.Front(); y != nil; y = y.Next() {
			secondpoint := y.Value.(*Point)
			if point.currentposition.x-secondpoint.currentposition.x > 200 || point.currentposition.y-secondpoint.currentposition.y > 200 {
				found = false
			}
		}
	}

	if found {

		fmt.Println()
		fmt.Printf("After %d second(s):\n", seconds)

		var grid map[int]map[int]rune
		grid = make(map[int]map[int]rune)
		for x := minx - 200; x < maxx+200; x++ {
			grid[x] = make(map[int]rune)
			for y := miny - 200; y < maxy+200; y++ {
				grid[x][y] = '.'
			}
		}

		for m := l.Front(); m != nil; m = m.Next() {
			point := m.Value.(*Point)
			grid[point.currentposition.x][point.currentposition.y] = '#'
		}

		for x := minx - 60; x < maxx+60; x++ {
			for y := miny; y < maxy+160; y++ {
				fmt.Printf("%s", string(grid[y][x]))
			}
			fmt.Println()
		}

		filename := fmt.Sprintf("%d.txt", seconds)
		f, _ := os.Create(filename)
		defer f.Close()
		w := bufio.NewWriter(f)
		for x := l.Front(); x != nil; x = x.Next() {
			point := x.Value.(*Point)
			pointstring := fmt.Sprintf("%d,%d\n", point.currentposition.y, point.currentposition.x)
			w.WriteString(pointstring)
		}
		w.Flush()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	l := list.New()

	// Read in the lines
	for scanner.Scan() {
		pstring := scanner.Text()
		p := parsePoints(pstring)
		l.PushBack(p)
	}

	// Let's simulate a bunch of seconds, I had no real idea how long it'd take
	for s := 1; s < 50000; s++ {
		simulate(s, l)
	}

}
