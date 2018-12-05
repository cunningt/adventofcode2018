package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// swap the case of a rune
func swapCaseRune(elem rune) rune {
	if unicode.IsUpper(elem) {
		return unicode.ToLower(elem)
	} else {
		return unicode.ToUpper(elem)
	}
}

// swap the case of a string
func swapCaseString(elem string) string {
	chars := []rune(elem)
	for i := 0; i < len(chars); i++ {
		chars[i] = swapCaseRune(chars[i])
	}
	return string(chars)
}

// Tried this, it took too long, gotta make this faster
func reaction(reactString string, pos int) string {
	if pos == len(reactString)-1 {
		return reactString
	}

	first := rune(reactString[pos])
	second := rune(reactString[pos+1])

	chars := []rune(reactString)

	if second == swapCaseRune(first) {
		newString := string(string(chars[:pos]) + string(chars[pos+2:]))
		return reaction(newString, 0)
	} else {
		return reaction(reactString, pos+1)
	}
}

// This is faster
func replaceReaction(reactString string, pos int) string {
	if pos == len(reactString)-1 {
		return reactString
	}

	// Let's replace everything from Aa to Zz
	for i := 65; i < (65 + 26); i++ {
		r := rune(i)
		swapped := swapCaseRune(r)
		replacement := string(r) + string(swapped)
		reactString = strings.Replace(reactString, replacement, "", -1)
	}

	// Let's replace everything from aA to zZ
	for i := 97; i < (97 + 26); i++ {
		r := rune(i)
		swapped := swapCaseRune(r)
		replacement := string(r) + string(swapped)
		reactString = strings.Replace(reactString, replacement, "", -1)
	}

	// Go to the next position
	return replaceReaction(reactString, pos+1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		reactString := scanner.Text()

		// Here are our placeholders for the smallest char / size
		smallestRune := rune(65)
		smallestSize := len(reactString)

		// Let's get the answer for part 1, do the reaction on the string
		result := replaceReaction(reactString, 0)
		resultLength := len(result)
		fmt.Printf("Answer to part 1 is %d\n", resultLength)

		// iterate over the alphabet to test the shortest polymer reaction
		for i := 65; i < (65 + 26); i++ {
			tempString := reactString

			r := rune(i)
			swapped := swapCaseRune(r)
			fmt.Printf("Removing all %c%c units results in ", r, swapped)

			// replace upper case, then replace lower case
			tempString = strings.Replace(tempString, string(r), "", -1)
			tempString = strings.Replace(tempString, string(swapped), "", -1)

			// do the reaction, report the result
			tempString = replaceReaction(tempString, 0)
			length := len(tempString)
			fmt.Printf("length of %d\n", length)

			// swap out smallest if necessary
			if length < smallestSize {
				smallestRune = r
				smallestSize = length
			}
		}

		fmt.Printf("Smallest combination was %c at %d\n", smallestRune, smallestSize)

	}
}
