package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	twoscount := 0
	threescount := 0
	for scanner.Scan() {

		boxid := scanner.Text()
		var m map[string]int
		m = make(map[string]int)

		for i := 0; i < len(boxid); i++ {
			char := string(boxid[i])
			m[char]++
		}
		fmt.Println("Boxid = " + boxid + "\n")

		twoflag := false
		threeflag := false
		for key, value := range m {
			if value == 2 {
				fmt.Println("Key:", key, "Value:", value)
				twoflag = true
			}
			if value == 3 {
				fmt.Println("Key:", key, "Value:", value)
				threeflag = true
			}
		}

		if twoflag {
			twoscount++
		}

		if threeflag {
			threescount++
		}

		fmt.Printf("Twoscount = %d, Threescount = %d, Checksum=%d\n", twoscount, threescount, (twoscount * threescount))

	}

	fmt.Printf("Twoscount = %d, Threescount = %d, Checksum=%d\n", twoscount, threescount, (twoscount * threescount))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
