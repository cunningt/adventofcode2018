package main

import (
	"bufio"
	"fmt"
	"os"
)

var forest map[int]map[int]rune
var original map[int]map[int]rune

const tree rune = '|'
const open rune = '.'
const lumberyard rune = '#'
const maxx = 50
const maxy = 50

func up(mychar rune, x int, y int) int {
	if mychar == original[x][y+1] {
		return 1
	}
	return 0
}

func down(mychar rune, x int, y int) int {
	if mychar == original[x][y-1] {
		return 1
	}
	return 0
}

func right(mychar rune, x int, y int) int {
	if mychar == original[x+1][y] {
		return 1
	}
	return 0
}

func left(mychar rune, x int, y int) int {
	if mychar == original[x-1][y] {
		return 1
	}
	return 0
}

func upright(mychar rune, x int, y int) int {
	if mychar == original[x+1][y+1] {
		return 1
	}
	return 0
}

func upleft(mychar rune, x int, y int) int {
	if mychar == original[x-1][y+1] {
		return 1
	}
	return 0
}

func downright(mychar rune, x int, y int) int {
	if mychar == original[x+1][y-1] {
		return 1
	}
	return 0
}

func downleft(mychar rune, x int, y int) int {
	if mychar == original[x-1][y-1] {
		return 1
	}
	return 0
}

func applyRules(x int, y int) {
	// open ground (.), trees (|), or a lumberyard (#)

	// An open acre will become filled with trees if three or more
	// adjacent acres contained trees. Otherwise, nothing happens.
	count := 0
	if original[x][y] == open {
		count += up(tree, x, y)
		count += down(tree, x, y)
		count += right(tree, x, y)
		count += left(tree, x, y)

		count += upright(tree, x, y)
		count += upleft(tree, x, y)
		count += downright(tree, x, y)
		count += downleft(tree, x, y)

		if count >= 3 {
			forest[x][y] = tree
		}
	}

	// An acre filled with trees will become a lumberyard if three
	// or more adjacent acres were lumberyards. Otherwise, nothing happens.
	count = 0
	if original[x][y] == tree {
		count += up(lumberyard, x, y)
		count += down(lumberyard, x, y)
		count += right(lumberyard, x, y)
		count += left(lumberyard, x, y)

		count += upright(lumberyard, x, y)
		count += upleft(lumberyard, x, y)
		count += downright(lumberyard, x, y)
		count += downleft(lumberyard, x, y)

		if count >= 3 {
			forest[x][y] = lumberyard
		}
	}

	// An acre containing a lumberyard will remain a lumberyard if it
	// was adjacent to at least one other lumberyard and at least one acre containing
	// trees. Otherwise, it becomes open.
	lumbercount := 0
	treecount := 0
	if original[x][y] == lumberyard {
		treecount += up(tree, x, y)
		treecount += down(tree, x, y)
		treecount += right(tree, x, y)
		treecount += left(tree, x, y)

		treecount += upright(tree, x, y)
		treecount += upleft(tree, x, y)
		treecount += downright(tree, x, y)
		treecount += downleft(tree, x, y)

		lumbercount += up(lumberyard, x, y)
		lumbercount += down(lumberyard, x, y)
		lumbercount += right(lumberyard, x, y)
		lumbercount += left(lumberyard, x, y)

		lumbercount += upright(lumberyard, x, y)
		lumbercount += upleft(lumberyard, x, y)
		lumbercount += downright(lumberyard, x, y)
		lumbercount += downleft(lumberyard, x, y)

		if lumbercount >= 1 && treecount >= 1 {
			forest[x][y] = lumberyard
		} else {
			forest[x][y] = open
		}
	}
}

func simulate(days int) {

	for d := 0; d < days; d++ {

		answer := scoreForest()
		fmt.Printf("Day %d, Score %d\n", d, answer)

		for x := 0; x < maxx; x++ {
			if d == 0 {
				original[x] = make(map[int]rune)
			}
			for y := 0; y < maxy; y++ {
				original[x][y] = forest[x][y]
			}
		}

		for x := 0; x < maxx; x++ {
			for y := 0; y < maxy; y++ {
				applyRules(x, y)
			}
		}
	}
}

func printForest() {
	for x := 0; x < maxx; x++ {
		for y := 0; y < maxy; y++ {
			fmt.Printf("%c", forest[x][y])
		}
		fmt.Println()
	}
}

func parseConstruction(txt string) map[int]rune {
	var con map[int]rune
	con = make(map[int]rune)
	r := []rune(txt)
	for i := 0; i < len(r); i++ {
		con[i] = r[i]
	}
	return con
}

func scoreForest() int {
	treecount := 0
	lumbercount := 0
	for x := 0; x < maxx; x++ {
		for y := 0; y < maxy; y++ {
			if tree == forest[x][y] {
				treecount++
			} else if lumberyard == forest[x][y] {
				lumbercount++
			}
		}
	}

	//fmt.Printf("Treecount : %d\n", treecount)
	//fmt.Printf("Lumbercount : %d\n", lumbercount)
	return (treecount * lumbercount)
}

func simulate(time int) {
	for t := 0; t <= time; t++ {

	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	forest = make(map[int]map[int]rune)
	original = make(map[int]map[int]rune)
	counter := 0

	// Read in the lines
	for scanner.Scan() {
		forest[counter] = make(map[int]rune)
		mystring := scanner.Text()
		forest[counter] = parseConstruction(mystring)
		counter++
	}

	// Part 1
	simulate(10)
	answer := scoreForest()
	fmt.Printf("Score : %d\n", answer)

	// Part 2 is a repeating pattern
	// 3277 = 206987
	// 3305 = 206987, 28
	total := 1000000000 - 3277
	modulus := total % 28
	fmt.Printf("Modulus = %d\n", modulus)
	// Looked at answers, the 19th repeating is
	// Day 3296, Score 233058
	fmt.Printf("Answer = %d\n", 233058)
}
