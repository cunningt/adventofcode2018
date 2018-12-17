package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// Debugging function - print the ring values
func printRing(circle *ring.Ring) {
	circle.Do(func(x interface{}) {
		fmt.Print(x)
		fmt.Print(" ")
	})
	fmt.Println()
}

func solveProblem(numrecipes int, input string) *ring.Ring {
	size := len(input)
	circle := ring.New(size)

	// Initialize the Ring
	r := []rune(input)
	for i := 0; i < len(r); i++ {
		circle.Value, _ = strconv.Atoi(string(r[i]))
		circle = circle.Next()
	}

	firstElf := circle.Move(0)
	secondElf := circle.Move(1)

	for x := 0; x < numrecipes+10; x++ {
		firstValue := firstElf.Value.(int)
		secondValue := secondElf.Value.(int)

		sum := firstValue + secondValue

		if sum > 9 {
			var firstRecipe int = sum / 10
			var secondRecipe int = sum % 10

			newentry := ring.New(1)
			newentry.Value = firstRecipe
			circle.Move(-1).Link(newentry)

			newentry = ring.New(1)
			newentry.Value = secondRecipe
			circle.Move(-1).Link(newentry)

		} else {
			var recipe int = sum

			newentry := ring.New(1)
			newentry.Value = recipe
			circle.Move(-1).Link(newentry)

		}

		// Pick new recipes
		firstMoves := 1 + firstElf.Value.(int)
		secondMoves := 1 + secondElf.Value.(int)

		firstElf = firstElf.Move(firstMoves)
		secondElf = secondElf.Move(secondMoves)

	}

	return circle
}

func solveSecondProblem(input string, answer string) int {
	size := len(input)
	circle := ring.New(size)

	// Initialize the Ring
	r := []rune(input)
	for i := 0; i < len(r); i++ {
		circle.Value, _ = strconv.Atoi(string(r[i]))
		circle = circle.Next()
	}

	firstElf := circle.Move(0)
	secondElf := circle.Move(1)

	counter := 2

	for {
		firstValue := firstElf.Value.(int)
		secondValue := secondElf.Value.(int)

		//fmt.Printf("First [%d] Second [%d]\n", firstValue, secondValue)

		sum := firstValue + secondValue

		if sum > 9 {
			var firstRecipe int = sum / 10
			var secondRecipe int = sum % 10

			newentry := ring.New(1)
			newentry.Value = firstRecipe
			circle.Move(-1).Link(newentry)

			counter++
			if counter > 6 {
				digits := findAnswer(counter-6, 6, circle)
				//fmt.Printf("digits=%s answer=%s, counter=%d\n", digits, answer, counter)
				if digits == answer {
					return (counter - 6)
				}
			}

			newentry = ring.New(1)
			newentry.Value = secondRecipe
			circle.Move(-1).Link(newentry)

			counter++

		} else {
			var recipe int = sum

			newentry := ring.New(1)
			newentry.Value = recipe
			circle.Move(-1).Link(newentry)

			counter++
		}

		// Pick new recipes
		firstMoves := 1 + firstElf.Value.(int)
		secondMoves := 1 + secondElf.Value.(int)

		firstElf = firstElf.Move(firstMoves)
		secondElf = secondElf.Move(secondMoves)

		if counter > 5 {
			digits := findAnswer(counter-6, 6, circle)
			//fmt.Printf("digits=%s answer=%s, counter=%d\n", digits, answer, counter)

			if digits == answer {

				return (counter - 6)
			}
		}
	}

	return -1
}

func findAnswer(num int, digits int, circle *ring.Ring) string {

	answer := ""
	for i := digits; i > 0; i-- {
		answer = answer + strconv.Itoa(circle.Move(0-i).Value.(int))
	}
	return answer
}

func findFirstAnswer(num int, digits int, circle *ring.Ring) string {

	answer := ""
	for i := 0; i < digits; i++ {
		answer = answer + strconv.Itoa(circle.Move(num+i).Value.(int))
	}
	return answer
}

func main() {

	// Part 1
	fmt.Println("====== Part1")
	circle := solveProblem(939601, "37")
	answer := findFirstAnswer(939601, 10, circle)
	fmt.Printf("Answer is %s\n", answer)

	// Part 2
	fmt.Println("====== Part2")
	twoAnswer := solveSecondProblem("37", "939601")
	fmt.Printf("Answer is %d\n", twoAnswer)

}
