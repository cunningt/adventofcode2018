package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	l := list.New()

	frequency := 0
	var m map[int]int
	m = make(map[int]int)

	for scanner.Scan() {
		inputfreq, err := strconv.Atoi(scanner.Text())
		l.PushBack(inputfreq)

		if err != nil {
			fmt.Println(err)
		}
	}

	e := l.Front()

	for true {

		if e == nil {
			e = l.Front()
		}
		inputfreq := e.Value.(int)

		if m[frequency] != 0 {
			fmt.Printf("Frequency %d repeated\n", frequency)
			os.Exit(0)
		} else {
			m[frequency] = 1
		}

		fmt.Printf("Current frequency %d, change of %d; resulting frequency %d.\n", frequency,
			inputfreq, (frequency + inputfreq))
		frequency = frequency + inputfreq

		if e != nil {
			e = e.Next()
		} else {
			e = l.Front()
		}

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
