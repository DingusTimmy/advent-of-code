package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// setup 4 int window
	var window [4]string
	window[0] = previous
	currentLine := 1
	windowIncreaseCount := 0

	// scan and compare remainder of file
	for scanner.Scan() {
		currentLine++
		current := scanner.Text()

		increased, printString := sonarSweep(previous, current)
		previous = current
		if increased {
			increaseCount++
		}
		fmt.Println(printString)

		window[3] = window[2]
		window[2] = window[1]
		window[1] = window[0]
		window[0] = current
		if currentLine > 3 {
			windowIncreased, windowPrintString := sonarSweep2(window[0], window[1], window[2], window[3])
			if windowIncreased {
				windowIncreaseCount++
			}
			fmt.Println(windowPrintString)
		}
	}

	fmt.Printf("Total increases: %v\n", increaseCount)
	fmt.Printf("Total window increases: %v\n", windowIncreaseCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sonarSweep(previous, current string) (increased bool, printString string) {
	printString = current
	increased = false

	intPrevious, err := strconv.Atoi(previous)
	if err != nil {
		log.Fatal(err)
	}
	intCurrent, err := strconv.Atoi(current)
	if err != nil {
		log.Fatal(err)
	}

	if intCurrent > intPrevious {
		printString = printString + " (increased)"
		increased = true
	} else if intCurrent < intPrevious {
		printString = printString + " (decreased)"
	} else {
		printString = printString + " (equal)"
	}

	return increased, printString
}

func sonarSweep2(a, b, c, d string) (increased bool, printString string) {
	intA, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	intB, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	intC, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}
	intD, err := strconv.Atoi(d)
	if err != nil {
		log.Fatal(err)
	}
	previous := intB + intC + intD
	current := intA + intB + intC
	printString = fmt.Sprintf("%v -> %v", previous, current)
	increased = false

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
