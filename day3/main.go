package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Claim struct {
	number int
	xpos   int
	ypos   int
	width  int
	height int
}

func newClaim(claimString string) *Claim {
	re := regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	matches := re.FindAllStringSubmatch(claimString, -1)

	c := new(Claim)
	c.number, _ = strconv.Atoi(matches[0][1])
	c.xpos, _ = strconv.Atoi(matches[0][2])
	c.ypos, _ = strconv.Atoi(matches[0][3])
	c.width, _ = strconv.Atoi(matches[0][4])
	c.height, _ = strconv.Atoi(matches[0][5])

	return c
}

func doesClaimIntersect(c Claim, row int, column int) bool {
	//fmt.Printf("%s %d %d\n", claim, row, column)

	//fmt.Printf("POSITION [%d,%d] CLAIMPOS[%d,%d] W/H[%d,%d]\n",
	//	row, column, claimxpos, claimypos, claimwidth, claimheight)

	if row > c.xpos && row <= (c.xpos+c.width) {
		//fmt.Println("%d < %d < %d", claimxpos, row, (claimxpos + claimwidth))
		if column > c.ypos && column <= (c.ypos+c.height) {
			return true
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	l := list.New()

	// Map of intersecting IDs
	var m map[int]bool
	m = make(map[int]bool)

	// Read in the input
	for scanner.Scan() {
		claimString := scanner.Text()
		c := newClaim(claimString)
		l.PushBack(c)
	}

	// Find claims that intersect more than twice
	patchCounter := 0
	for row := 0; row <= 1000; row++ {
		for column := 0; column <= 1000; column++ {
			claimCounter := 0
			var claimmap map[int]bool
			claimmap = make(map[int]bool)

			for e := l.Front(); e != nil; e = e.Next() {

				claim := e.Value.(*Claim)
				claimIntersect := doesClaimIntersect(*claim, row, column)

				if claimIntersect {
					claimCounter++
					claimmap[claim.number] = true
				}
			}

			if claimCounter >= 2 {
				patchCounter++

				for key, _ := range claimmap {
					m[key] = true
				}

			}

		}
	}

	fmt.Printf("There are %d intersecting claims...\n", patchCounter)

	// Find non-intersecting claim
	for e := l.Front(); e != nil; e = e.Next() {

		claim := e.Value.(*Claim)
		if _, ok := m[claim.number]; ok {
		} else {
			fmt.Printf("Claim %d does not intersect\n", claim.number)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
