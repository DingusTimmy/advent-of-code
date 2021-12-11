package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2021/inputs/02-dive.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// part 1 setup
	xpos1 := 0
	depth1 := 0

	// part 2 setup
	aim := 0
	xpos2 := 0
	depth2 := 0

	// scan file
	for scanner.Scan() {
		// Assuming the input follows the "command value" format
		// details[0] is the command, details[1] is the value
		details := strings.Fields(scanner.Text())
		value, err := strconv.Atoi(details[1])
		if err != nil {
			log.Fatal(err)
		}

		xIncrease, depthIncrease := dive(details[0], value)
		xpos1 += xIncrease
		depth1 += depthIncrease

		aimIncrease, xIncrease, depthIncrease := dive2(details[0], value, aim)
		aim += aimIncrease
		xpos2 += xIncrease
		depth2 += depthIncrease
	}

	fmt.Printf("1 ~ X-Position: %v\n", xpos1)
	fmt.Printf("1 ~ Depth: %v\n", depth1)
	fmt.Printf("1 ~ Result: %v\n", xpos1*depth1)

	fmt.Printf("2 ~ X-Position: %v\n", xpos2)
	fmt.Printf("2 ~ Depth: %v\n", depth2)
	fmt.Printf("2 ~ Result: %v\n", xpos2*depth2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func dive(command string, value int) (xIncrease, depthIncrease int) {

	switch command {
	case "forward":
		return value, 0
	case "down":
		return 0, value
	case "up":
		return 0, -value
	}

	return 0, 0
}

func dive2(command string, value, currentAim int) (aimIncrease, xIncrease, depthIncrease int) {

	switch command {
	case "forward":
		return 0, value, currentAim * value
	case "down":
		return value, 0, 0
	case "up":
		return -value, 0, 0
	}

	return 0, 0, 0
}
