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

func solveProblem(numplayers int, lastmarble int) int {
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
	for _, v := range scores {
		if v > maxscore {
			maxscore = v
		}
	}
	return maxscore
}

func main() {

	// Part 1
	fmt.Println("====== Part1")
	maxscore := solveProblem(405, 70953)
	fmt.Printf("Maximum score is %d\n", maxscore)

	// Part 2
	fmt.Println("====== Part2")
	maxscore = solveProblem(405, 7095300)
	fmt.Printf("Maximum score is %d\n", maxscore)

}
