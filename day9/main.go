package main

import (
	"container/ring"
	"fmt"
)

// Debugging function - print the ring values
func printRing(circle *ring.Ring) {
	circle.Do(func(x interface{}) {
		fmt.Print(x)
		fmt.Print(" ")
	})
	fmt.Println()
}

func main() {

	// It'd be nicer if I did something with argv here or did
	// something with reading in the number of players / lastmarble value
	numplayers := 405
	lastmarble := 7095300

	var scores map[int]int
	scores = make(map[int]int)

	circle := ring.New(1)
	circle.Value = 0
	currentmarble := 1
	currentplayer := 1

	for i := 1; i <= lastmarble; i++ {

		// Special case - marble divisible by 23
		if currentmarble%23 == 0 {
			scores[currentplayer] += currentmarble
			for b := 0; b < 6; b++ {
				circle = circle.Prev()
			}

			scores[currentplayer] += circle.Value.(int)
			circle = circle.Prev()
			circle.Unlink(1)
		} else {
			circle = circle.Next()
			circle = circle.Next()
			newentry := ring.New(1)
			newentry.Value = i
			circle.Link(newentry)
		}
		//fmt.Printf("[%d] ", currentplayer)
		//printRing(circle)
		currentmarble++

		// Move to the next elf
		if currentplayer == numplayers {
			currentplayer = 1
		} else {
			currentplayer++
		}

	}

	// Print the scores
	maxscore := 0
	for k, v := range scores {
		if v > maxscore {
			maxscore = v
		}
		fmt.Printf("Player %d, Score %d\n", k, v)
	}
	fmt.Printf("Maximum score is %d\n", maxscore)
}
