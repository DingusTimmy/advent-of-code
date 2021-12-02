package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2021/inputs/01-sonar-sweep.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	increaseCount := 0

	// scan initial line
	scanner.Scan()
	fmt.Println(scanner.Text())
	previous := scanner.Text()

	// scan and compare remainder of file
	for scanner.Scan() {
		current := scanner.Text()
		increased, printString := sonarSweep(previous, current)
		previous = current
		if increased {
			increaseCount++
		}
		fmt.Println(printString)
	}

	fmt.Printf("Total increases: %v", increaseCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sonarSweep(previous, current string) (increased bool, printString string) {
	printString = current

	if current > previous {
		printString = printString + " (increased)"
		increased = true
	} else if current < previous {
		printString = printString + " (decreased)"
	} else {
		printString = printString + " (equal)"
	}

	return increased, printString
}
