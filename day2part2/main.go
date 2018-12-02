package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	l := list.New()

	// Read in the input
	for scanner.Scan() {
		l.PushBack(scanner.Text())
	}

	for e := l.Front(); e != nil; e = e.Next() {

		firstString := e.Value.(string)
		for f := l.Front(); f != nil; f = f.Next() {
			secondString := f.Value.(string)

			counter := 0
			var m map[int]int
			m = make(map[int]int)

			for i := 0; i < len(firstString); i++ {
				if firstString[i] == secondString[i] {
					m[i] = 1
					counter++
				} else {
					m[i] = 0
				}
			}

			if counter == (len(firstString) - 1) {
				fmt.Println("firstString " + firstString + " second " + secondString)
				for i := 0; i < len(firstString); i++ {
					if firstString[i] == secondString[i] {
						fmt.Print(string(firstString[i]))
					}
				}
				fmt.Println()

			}
		}
	}
}
