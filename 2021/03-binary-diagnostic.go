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

	var inputs []string

	// scan initial line to setup counters
	scanner.Scan()
	initial := scanner.Text()
	inputs = append(inputs, initial)
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
		inputs = append(inputs, binary)
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

	// get part2 ratings
	oxyGenRating, co2ScrubRating := binaryDiagnostic2(inputs)

	// bit conversion
	decGamma, err := strconv.ParseInt(gammaRate, 2, length+1)
	if err != nil {
		log.Fatal(err)
	}
	decEpsilon, err := strconv.ParseInt(epsilonRate, 2, length+1)
	if err != nil {
		log.Fatal(err)
	}
	decOxyGenRating, err := strconv.ParseInt(oxyGenRating, 2, length+1)
	if err != nil {
		log.Fatal(err)
	}
	decCo2ScrubberRating, err := strconv.ParseInt(co2ScrubRating, 2, length+1)
	if err != nil {
		log.Fatal(err)
	}

	// results
	fmt.Printf("1 ~ Gamma Rate: %v (%v)\n", decGamma, gammaRate)
	fmt.Printf("1 ~ Epsilon Rate: %v (%v)\n", decEpsilon, epsilonRate)
	fmt.Printf("1 ~ Result: %v\n", decGamma*decEpsilon)

	fmt.Printf("2 ~ Oxygen Generator Rating: %v (%v)\n", decOxyGenRating, oxyGenRating)
	fmt.Printf("2 ~ CO2 Scrubber Rating: %v (%v)\n", decCo2ScrubberRating, co2ScrubRating)
	fmt.Printf("2 ~ Result: %v\n", decOxyGenRating*decCo2ScrubberRating)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func binaryDiagnostic2(inputs []string) (oxyGenRating, co2ScrubRating string) {
	var positives []string
	var negatives []string
	dominantBit := 0 // < 0 means 0 is dominant, >= 0 means 1 is dominant

	for _, input := range inputs {
		if input[0] == '1' {
			dominantBit++
			positives = append(positives, input)
		} else {
			dominantBit--
			negatives = append(negatives, input)
		}
	}

	if dominantBit < 0 {
		return oxygenGeneratorRating(negatives), co2ScrubberRating(positives)
	} else {
		return oxygenGeneratorRating(positives), co2ScrubberRating(negatives)
	}
}

func oxygenGeneratorRating(inputs []string) string {
	position := 1
	for len(inputs) > 1 {
		var positives []string
		var negatives []string
		dominantBit := 0 // < 0 means 0 is dominant, >= 0 means 1 is dominant

		for _, input := range inputs {
			if input[position] == '1' {
				dominantBit++
				positives = append(positives, input)
			} else {
				dominantBit--
				negatives = append(negatives, input)
			}
		}

		if dominantBit < 0 {
			inputs = negatives
		} else {
			inputs = positives
		}
		position++
	}
	return inputs[0]
}

func co2ScrubberRating(inputs []string) string {
	position := 1
	for len(inputs) > 1 {
		var positives []string
		var negatives []string
		dominantBit := 0 // < 0 means 0 is dominant, >= 0 means 1 is dominant

		for _, input := range inputs {
			if input[position] == '1' {
				dominantBit++
				positives = append(positives, input)
			} else {
				dominantBit--
				negatives = append(negatives, input)
			}
		}

		if dominantBit < 0 {
			inputs = positives
		} else {
			inputs = negatives
		}
		position++
	}
	return inputs[0]
}
