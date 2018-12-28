package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var fountains map[int]map[int]rune
var veins *list.List

type Vein struct {
	minx int
	maxx int
	miny int
	maxy int
}

func parseString(veinString string) *Vein {
	v := new(Vein)

	re := regexp.MustCompile("x=([0-9]+), y=([0-9]+)\\.\\.([0-9]+)")
	matches := re.FindAllStringSubmatch(veinString, -1)
	if re.MatchString(veinString) {
		v.minx, _ = strconv.Atoi(matches[0][1])
		v.maxx, _ = strconv.Atoi(matches[0][1])

		v.miny, _ = strconv.Atoi(matches[0][2])
		v.maxy, _ = strconv.Atoi(matches[0][3])
	}

	re = regexp.MustCompile("y=([0-9]+), x=([0-9]+)\\.\\.([0-9]+)")
	matches = re.FindAllStringSubmatch(veinString, -1)
	if re.MatchString(veinString) {
		v.miny, _ = strconv.Atoi(matches[0][1])
		v.maxy, _ = strconv.Atoi(matches[0][1])

		v.minx, _ = strconv.Atoi(matches[0][2])
		v.maxx, _ = strconv.Atoi(matches[0][3])

	}

	return v
}

func printVein(v *Vein) {
	fmt.Printf("x=%d..%d y=%d..%d\n", v.minx, v.maxx, v.miny, v.maxy)
}

func printFountain() {
	for y := 0; y < 15; y++ {
		for x := 490; x < 510; x++ {
			fmt.Printf("%c", fountains[x][y])
		}
		fmt.Println()
	}
}

func borderSand(x int, y int) bool {
	if fountains[x-1][y] == '#' || fountains[x][y+1] == '#' || fountains[x+1][y] == '#' {
		return true
	}
	if fountains[x-1][y+1] == '#' || fountains[x+1][y+1] == '#' {
		return true
	}
	return false
}

func simulate(time int) {
	for t := 0; t < time; t++ {

		// drip the water
		for y := 1; y < 15; y++ {
			for x := 0; x < 510; x++ {

				// Case Fountain
				if fountains[x][y-1] == '+' && fountains[x][y] == '.' {
					fountains[x][y] = '|'
				}

				// Case Water
				if fountains[x][y-1] == '|' && fountains[x][y] == '.' {
					fountains[x][y] = '|'
				}

				// Case flooding
				if fountains[x][y] == '.' {
					if fountains[x-1][y] == '|' || fountains[x+1][y] == '|' {
						if borderSand(x, y) {
							fountains[x][y] = '|'
						}
					}
				}

				// Case water all around
				if fountains[x-1][y] == '|' && fountains[x+1][y] == '|' && fountains[x][y] == '.' &&
					fountains[x+1][y+1] == '|' && fountains[x][y+1] == '|' && fountains[x-1][y+1] == '|' {
					fountains[x][y] = '|'
				}

				// Case water all around
				if fountains[x+1][y] == '|' && fountains[x][y] == '.' &&
					fountains[x-1][y] == '|' &&
					(fountains[x+1][y+1] == '|' || fountains[x+1][y+1] == '#') &&
					(fountains[x][y+1] == '|' || fountains[x][y+1] == '#') &&
					(fountains[x-1][y+1] == '|' || fountains[x-1][y+1] == '#') {
					fountains[x][y] = '|'
				}

			}
		}

		// resting water?
		min := -1
		for y := 1; y <= 20; y++ {
			for x := 0; x <= 510; x++ {
				if fountains[x][y] == '#' {
					if min >= 0 && x > min+1 {
						// Test to see if there are all '|' between

						flag := true
						for k := min + 1; k <= x-1; k++ {
							// if we see running water, and below it is standing water or
							// '#', turn to standing water

							if fountains[k][y] == '|' && (fountains[k][y+1] == '~' || fountains[k][y+1] == '#') {
							} else {
								flag = false
							}
						}

						if flag {
							for k := min + 1; k <= x-1; k++ {
								fountains[k][y] = '~'
							}
						}
						// Set min to x
						min = x
					} else {
						min = x
					}
				}
			}
		}

		printFountain()

	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fountains = make(map[int]map[int]rune)
	veins := list.New()

	// Initialize
	for x := 0; x < 510; x++ {
		fountains[x] = make(map[int]rune)
		for y := 0; y < 20; y++ {
			fountains[x][y] = '.'
		}
	}

	// Set fountain point
	fountains[500][0] = '+'

	// Read in the lines
	for scanner.Scan() {
		mystring := scanner.Text()
		v := parseString(mystring)
		printVein(v)
		veins.PushBack(v)

		for x := v.minx; x <= v.maxx; x++ {
			for y := v.miny; y <= v.maxy; y++ {
				fountains[x][y] = '#'
			}
		}
	}

	printFountain()

	simulate(10)

}
