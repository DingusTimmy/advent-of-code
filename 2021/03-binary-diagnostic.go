package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2021/inputs/03-binary-diagnostic.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scan initial line to setup counters
	scanner.Scan()
	initial := scanner.Text()
	length := len(initial)
	counts := make([]int, length)
	for i := 0; i < length; i++ {
		switch initial[i] {
		case '0':
			counts[i]++
		case '1':
			counts[i]--
		}
	}

	// scan file
	for scanner.Scan() {
		binary := scanner.Text()
		for i := 0; i < length; i++ {
			switch binary[i] {
			case '0':
				counts[i]++
			case '1':
				counts[i]--
			}
		}
	}

	// create bit arrays
	gammaRate := ""
	epsilonRate := ""
	for _, count := range counts {
		if count <= 0 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	// bit conversion
	decGamma, err := strconv.ParseInt(gammaRate, 2, length+1)
	if err != nil {
		log.Fatal(err)
	}
	decEpsilon, err := strconv.ParseInt(epsilonRate, 2, length+1)
	if err != nil {
		log.Fatal(err)
	}

	// results
	fmt.Printf("1 ~ Gamma Rate: %v (%v)\n", decGamma, gammaRate)
	fmt.Printf("1 ~ Epsilon Rate: %v (%v)\n", decEpsilon, epsilonRate)
	fmt.Printf("1 ~ Result: %v\n", decGamma*decEpsilon)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
