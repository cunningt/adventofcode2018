package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	u int
	v int
	x int
	y int
}

func parsePoint(point string) *Point {
	p := new(Point)

	re := regexp.MustCompile("([-0-9]+),([-0-9]+),([-0-9]+),([-0-9]+)")
	if re.MatchString(point) {
		matches := re.FindAllStringSubmatch(point, -1)
		p.u, _ = strconv.Atoi(matches[0][1])
		p.v, _ = strconv.Atoi(matches[0][2])
		p.x, _ = strconv.Atoi(matches[0][3])
		p.y, _ = strconv.Atoi(matches[0][4])
	}
	return p
}

// Compute the distance
func distance(pointone *Point, pointtwo *Point) float64 {
	sum := math.Abs(float64(pointone.u-pointtwo.u)) + math.Abs(float64(pointone.v-pointtwo.v)) + math.Abs(float64(pointone.x-pointtwo.x)) + math.Abs(float64(pointone.y-pointtwo.y))
	return sum
}

func printPoint(point *Point) {
	fmt.Printf("%d,%d,%d,%d\n", point.u, point.v, point.x, point.y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	points := list.New()
	constellations := list.New()

	// Read in all of the points
	for scanner.Scan() {
		point := scanner.Text()
		p := parsePoint(point)

		points.PushBack(p)
	}

	pointmap := make(map[int]Point)
	counter := 0
	mapcounter := 0
	for {
		if points.Len() == 0 {
			break
		}

		if constellations.Len() == 0 {
			p := points.Front()
			point := points.Front().Value.(*Point)
			pointmap[mapcounter] = *point
			mapcounter++
			//fmt.Printf("Initial constellation %d for %d,%d,%d,%d\n", counter, point.u, point.v, point.x, point.y)
			constellations.PushBack(pointmap)
			points.Remove(p)
		}

		found := false
		for l := points.Front(); l != nil; l = l.Next() {
			point := l.Value.(*Point)
			//fmt.Printf("Testing point %d,%d,%d,%d\n", point.u, point.v, point.x, point.y)
			for _, value := range pointmap {

				dist := distance(&value, point)
				//fmt.Printf("Distance %f from %d,%d,%d,%d ... point %d,%d,%d,%d to constellation %d\n", dist, point.u, point.v, point.x, point.y, value.u, value.v, value.x, value.y, counter)

				if dist <= 3 {
					found = true
					p := l.Value.(*Point)
					pointmap[mapcounter] = *p
					mapcounter++
					points.Remove(l)
					break
					//fmt.Printf("Added : distance %f from %d,%d,%d,%d ... point %d,%d,%d,%d to constellation %d\n", dist, point.u, point.v, point.x, point.y, value.u, value.v, value.x, value.y, counter)

				}
			}
		}

		if !found {
			pointmap = make(map[int]Point)
			p := points.Front()
			point := points.Front().Value.(*Point)
			pointmap[0] = *point
			constellations.PushBack(pointmap)
			points.Remove(p)
			counter++
			fmt.Printf("New constellation %d for %d,%d,%d,%d\n", counter, point.u, point.v, point.x, point.y)
		}

	}

	// Part one
	fmt.Printf("Number of constellations %d\n", constellations.Len())

	// Part two - I don't have enough stars yet

}
