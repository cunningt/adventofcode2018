package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Rule struct {
	state  string
	result rune
}

func printRule(rule *Rule) {
	fmt.Printf("%s => %c\n", rule.state, rule.result)
}

func parseInitialState(stateString string) string {

	var initialstate string = ""
	re := regexp.MustCompile("initial state: ([\\#\\.]+)")

	if re.MatchString(stateString) {
		matches := re.FindAllStringSubmatch(stateString, -1)

		initialstate = ".........." + matches[0][1] + "......................................................................................................................................................................................................................................................................."
	}

	return initialstate
}

func parseRule(ruleString string) *Rule {

	rule := new(Rule)
	re := regexp.MustCompile("([\\.\\#]+) => ([\\.\\#]+)")

	if re.MatchString(ruleString) {
		matches := re.FindAllStringSubmatch(ruleString, -1)

		rule.state = matches[0][1]

		rule.state = strings.Replace(rule.state, "#", "[\\#]", -1)
		rule.state = strings.Replace(rule.state, ".", "[\\.]", -1)

		r := []rune(matches[0][2])
		rule.result = r[0]

	}
	printRule(rule)

	return rule
}

func applyRule(initialstate string, currentstate string, rule *Rule) string {
	r := regexp.MustCompile(rule.state)
	initstate := "...." + initialstate + "...."
	strArr := []rune(currentstate)

	idx := 0
	for {
		index := r.FindStringIndex(initstate[idx:])
		if index == nil {
			break
		}
		nextidx := idx + index[0]

		// The position we want to change is + 2 (middle of 5char string), and -4 padding
		pos := nextidx + 2 - 4
		//fmt.Printf("%s INDEX %d\n", rule.state, pos)

		//fmt.Printf("POS %d RESULT %c LEN %d\n", pos, rule.result, len(strArr))
		if pos >= 0 && (nextidx < len(initialstate)+2) {
			strArr[pos] = rule.result
		}
		idx = nextidx + 1
	}
	return string(strArr)
}

func countPots(currentstring string) int {
	re := regexp.MustCompile("([\\#])")

	count := 0
	if re.MatchString(currentstring) {
		indexes := re.FindAllStringIndex(currentstring, -1)
		for _, element := range indexes {
			count += (element[0] - 10)
			//fmt.Printf("Found pot %d\n", (element[0] - 3))
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	l := list.New()

	var stringArr []string
	var stringcount int = 0

	// Read in the lines
	for scanner.Scan() {
		stringArr = append(stringArr, scanner.Text())
		stringcount++
	}

	for i := 2; i < len(stringArr); i++ {
		rule := parseRule(stringArr[i])
		l.PushBack(rule)
	}

	initialstate := parseInitialState(stringArr[0])

	currentstate := initialstate
	genstate := currentstate

	fmt.Printf("0: %s\n", initialstate)

	//lastcount := 0
	for gen := 1; gen <= 20; gen++ {
		currentstate = strings.Replace(initialstate, "#", ".", -1)
		//currentstate := "......................................."
		for r := l.Front(); r != nil; r = r.Next() {
			rule := r.Value.(*Rule)
			currentstate = applyRule(genstate, currentstate, rule)
		}
		fmt.Printf("%d: %s\n", gen, currentstate)
		genstate = currentstate

		//fmt.Printf("%d Number of Pots : %d, diff of %d\n", gen, countPots(genstate), (countPots(genstate) - lastcount))
		//lastcount = countPots(genstate)

	}
	fmt.Printf("Number of Pots : %d\n", countPots(genstate))

	// Part 2 - it repeats - after looking at the counts, they seem to increase by 96
	// 190 Number of Pots : 19327, diff of 96
	// 191 Number of Pots : 19423, diff of 96
	// 192 Number of Pots : 19519, diff of 96
	// 193 Number of Pots : 19615, diff of 96
	// Formula is X * 96 + 1087

	part2 := (50000000000 * 96) + 1087
	fmt.Printf("Answer for Part2 is %d\n\n", part2)

}
