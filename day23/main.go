package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	v int
	x int
	y int
	r int
}

type Range struct {
	min int
	max int
}

func distance(pointone *Point, pointtwo *Point) float64 {
	sum := math.Abs(float64(pointone.v-pointtwo.v)) + math.Abs(float64(pointone.x-pointtwo.x)) + math.Abs(float64(pointone.y-pointtwo.y))
	return sum
}

func (r Range) InRange(i int) bool {
	return i >= r.min && i <= r.max
}

func parsePoint(stateString string) *Point {

	re := regexp.MustCompile("pos=<([\\-0-9]+),([\\-0-9]+),([\\-0-9]+)>, r=([\\-0-9]+)")

	p := new(Point)
	if re.MatchString(stateString) {
		matches := re.FindAllStringSubmatch(stateString, -1)
		p.v, _ = strconv.Atoi(matches[0][1])
		p.x, _ = strconv.Atoi(matches[0][2])
		p.y, _ = strconv.Atoi(matches[0][3])
		p.r, _ = strconv.Atoi(matches[0][4])
	}

	return p
}

func (p Point) Total() int {
	return int(math.Abs(float64(p.v))) + int(math.Abs(float64(p.x))) + int(math.Abs(float64(p.y)))
}

func (p Point) InRange(p2 Point) bool {
	d := distance(&p, &p2)
	if float64(p.r) >= d {
		return true
	}
	return false
}

func printPoint(p Point) {
	fmt.Printf("%d,%d,%d r=%d\n", p.v, p.x, p.y, p.r)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var points []Point

	for scanner.Scan() {
		p := parsePoint(scanner.Text())
		printPoint(*p)
		points = append(points, *p)
	}

	var distances map[int]int
	distances = make(map[int]int)
	maxindex := -1
	maxnum := -1
	maxradius := -1

	maxv := -1
	maxx := -1
	maxy := -1

	for i := 0; i < len(points); i++ {
		pone := points[i]
		counter := 0
		for x := 0; x < len(points); x++ {
			ptwo := points[x]
			d := distance(&pone, &ptwo)
			if pone.r >= int(d) {
				distances[x] = distances[x] + 1
				counter++
			}
		}

		if math.Abs(float64(pone.v)) > float64(maxv) {
			maxv = int(math.Abs(float64(pone.v)))
		}
		if math.Abs(float64(pone.x)) > float64(maxx) {
			maxx = int(math.Abs(float64(pone.x)))
		}
		if math.Abs(float64(pone.y)) > float64(maxy) {
			maxy = int(math.Abs(float64(pone.y)))
		}

		if pone.r > maxradius {
			maxradius = pone.r
			maxindex = i
			maxnum = counter
		}

	}

	fmt.Printf("Maximum radius is %d.  There are %d points within its range.   It is at ", maxradius, maxnum)
	ans := points[maxindex]
	printPoint(ans)

	maxindex = -1
	maxnum = -1
	for k, v := range distances {
		if v > maxnum {
			maxnum = v
			maxindex = k
		}
	}
	answer := points[maxindex]
	fmt.Printf("Bot in range of most other bots is ")
	printPoint(answer)
	fmt.Printf("It is in range of %d bots\n", maxnum)

	// Part 2
	// Kind of can't believe this works - the more elegant solution is to search smaller and smaller boxes
	// to see how many nanobots are in each box, and then subdivide accordingly.    What we can also do here
	// is figure out the distance from 0, plus or minus the radius for each nanobot.    Between the most extreme
	// distance +- r deltas, we search how many points are in range.    The distance that's within the most ranges
	// wins.   I don't think this would work for a large dataset, but we're only working with 1,000 points here.

	var ranges []Range
	minmin := 10000000000
	maxmax := -10000000000
	for i := 0; i < len(points); i++ {
		min := points[i].Total() - points[i].r
		max := points[i].Total() + points[i].r
		r := Range{min, max}
		ranges = append(ranges, r)

		if min < minmin {
			minmin = min
		}
		if max > maxmax {
			maxmax = max
		}
	}

	maxr := -1
	maxcounter := 0
	for r := minmin; r < maxmax; r++ {
		counter := 0
		for i := 0; i < len(ranges); i++ {
			if ranges[i].InRange(r) {
				counter++
			}
		}
		if counter > maxcounter {
			maxcounter = counter
			maxr = r
		}
	}

	fmt.Printf("MaxCounter = %d MaxR = %d\n", maxcounter, maxr)
}
