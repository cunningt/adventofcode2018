package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Point struct {
	x int
	y int
}

func parsePoint(point string) *Point {
	p := new(Point)

	re := regexp.MustCompile("([0-9]+), ([0-9]+)")
	if re.MatchString(point) {
		matches := re.FindAllStringSubmatch(point, -1)
		p.x, _ = strconv.Atoi(matches[0][1])
		p.y, _ = strconv.Atoi(matches[0][2])
	}
	return p
}

// Compute the distance
func distance(pointone *Point, pointtwo *Point) float64 {
	sum := math.Abs(float64(pointone.x-pointtwo.x)) + math.Abs(float64(pointone.y-pointtwo.y))
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var points []*Point

	// Read in all of the points
	maxcoord := 0
	for scanner.Scan() {
		point := scanner.Text()
		p := parsePoint(point)

		if p.x > maxcoord {
			maxcoord = p.x
		}
		if p.y > maxcoord {
			maxcoord = p.y
		}
		points = append(points, p)
	}

	var counter map[int]int
	counter = make(map[int]int)

	// List of the infinites - check if they are the min x/y
	var infinite map[int]bool
	infinite = make(map[int]bool)

	regionCount := 0
	maxcoord = 400

	// Loop through the grid, then loop through the points to determine distance
	for x := 0; x < maxcoord; x++ {
		for y := 0; y < maxcoord; y++ {
			p := new(Point)
			p.x = x
			p.y = y

			var totalDistance float64
			var mindistance float64
			var minindex int

			mindistance = 1000000
			for l := 0; l < len(points); l++ {
				dist := distance(p, points[l])
				totalDistance += dist
				if dist < mindistance {
					mindistance = dist
					minindex = l
				} else if dist == mindistance {
					// Ignore ties
					minindex = -1
				}
			}

			// For part 2, check if the total distance is less than 10000
			if totalDistance < 10000 {
				regionCount++
			}

			// Mark infinites
			if x == 0 || x == (maxcoord-1) || y == 0 || y == (maxcoord-1) {
				infinite[minindex] = true
			}

			// If it wasn't a tie, then we increment
			if minindex > 0 {
				counter[minindex]++
			}
		}
	}

	// Find maximum
	maxindex := 0
	maxdistances := 0
	for k, v := range counter {
		if _, ok := infinite[k]; ok {
		} else {
			if v > maxdistances {
				maxindex = k
				maxdistances = v
			}
		}
	}

	// Sorting, don't print out infinites
	n := map[int][]int{}
	var a []int
	for k, v := range counter {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {

			if _, ok := infinite[s]; ok {
			} else {
				fmt.Printf("%d (%d,%d), %d\n", s, points[s].x, points[s].y, k)
			}
		}
	}
	fmt.Printf("Maximum is %d at index %d\n", maxdistances, maxindex)
	fmt.Printf("Region count : %d\n", regionCount)
}
