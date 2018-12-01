package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	frequency := 0

	for scanner.Scan() {

		inputfreq, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}

		if inputfreq > 0 {
			fmt.Printf("Current frequency %d, change of +%d; resulting frequency %d.\n", frequency,
				inputfreq, (frequency + inputfreq))
		} else {
			fmt.Printf("Current frequency %d, change of %d; resulting frequency %d.\n", frequency,
				inputfreq, (frequency + inputfreq))
		}
		frequency = frequency + inputfreq
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
