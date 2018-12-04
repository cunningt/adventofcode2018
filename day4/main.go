package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Event struct {
	guardnumber int
	timestamp   string
	action      string
}

type Minutes struct {
	minutes map[int]int
}

func parseString(guardString string) *Event {
	e := new(Event)

	re := regexp.MustCompile("\\[(.*)\\] Guard #([0-9]+) begins shift")
	asre := regexp.MustCompile("\\[(.*)\\] falls asleep")
	wmre := regexp.MustCompile("\\[(.*)\\] wakes up")

	if asre.MatchString(guardString) {
		matches := asre.FindAllStringSubmatch(guardString, -1)
		e.timestamp = matches[0][1]
		e.action = "falls asleep"
	} else if wmre.MatchString(guardString) {
		matches := wmre.FindAllStringSubmatch(guardString, -1)
		e.timestamp = matches[0][1]
		e.action = "wakes up"
	} else if re.MatchString(guardString) {
		matches := re.FindAllStringSubmatch(guardString, -1)
		e.timestamp = matches[0][1]
		e.guardnumber, _ = strconv.Atoi(matches[0][2])
		e.action = "begins shift"
	} else {
		// error
	}

	return e
}

func parseMinute(timest string) int {
	re := regexp.MustCompile("00:([0-9]+)")
	matches := re.FindAllStringSubmatch(timest, -1)
	min, _ := strconv.Atoi(matches[0][1])
	return min
}

func calcTime(begin string, end string) int {
	layout := "2006-01-02 15:04"
	t1, _ := time.Parse(layout, begin)
	t2, _ := time.Parse(layout, end)

	delta := t2.Sub(t1)

	return int(delta.Minutes())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Map of intersecting IDs
	var m map[int]int
	m = make(map[int]int)

	var minutemap map[int]Minutes
	minutemap = make(map[int]Minutes)

	// Read in the input
	currentGuardNumber := 0
	//shiftBegin := ""
	sleepTimestamp := ""

	for scanner.Scan() {
		guardEvent := parseString(scanner.Text())

		if guardEvent.action == "begins shift" {
			currentGuardNumber = guardEvent.guardnumber
		} else if guardEvent.action == "wakes up" {
			wakeupTimestamp := guardEvent.timestamp
			minutes := calcTime(sleepTimestamp, wakeupTimestamp)
			m[currentGuardNumber] = m[currentGuardNumber] + minutes

			startMinute := parseMinute(sleepTimestamp)
			endMinute := parseMinute(wakeupTimestamp)

			for i := startMinute; i < endMinute; i++ {
				tmpMin := minutemap[currentGuardNumber]
				if len(tmpMin.minutes) == 0 {
					tmpMin.minutes = make(map[int]int)
				}
				tmpMin.minutes[i]++
				minutemap[currentGuardNumber] = tmpMin
			}

			sleepTimestamp = ""
		} else if guardEvent.action == "falls asleep" {
			sleepTimestamp = guardEvent.timestamp
		}
	}

	for key, value := range m {
		fmt.Printf("%d,%d ", value, key)

		// sort the minutemap
		tmpMin := minutemap[key]
		if len(tmpMin.minutes) == 0 {
			tmpMin.minutes = make(map[int]int)
		}

		n := 0
		biggestmin := 0
		biggestvalue := 0
		for k, v := range tmpMin.minutes {
			if v > n {
				n = v
				biggestmin = k
				biggestvalue = v
			} else {
			}
		}
		fmt.Printf("minute %d value %d\n", biggestmin, biggestvalue)
	}
}
