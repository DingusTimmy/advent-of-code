package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2021/inputs/04-giant-squid.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scan initial line to get which numbers are called
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	fmt.Println(numbers)

	// create map for tracking position of numbers
	boardMap := make(map[string][]int)

	// scan file to create bingo boards
	numberOfBoards := 0
	var boards []Board
	row := 0
	boardSum := 0
	for scanner.Scan() {
		if row > 5 {
			numberOfBoards++
			if boardSum > 0 {
				boards = append(boards, Board{sum: boardSum})
			}
			row = 0
			boardSum = 0
		}
		boardValues := strings.Fields(scanner.Text())
		for i, val := range boardValues {
			boardMap[val] = append(boardMap[val], (numberOfBoards*100)+((i+1)*10)+row) // board x 100 + column x 10 + row
			intVal, _ := strconv.Atoi(val)
			boardSum += intVal
		}
		row++
	}
	if row > 5 {
		numberOfBoards++
		if boardSum > 0 {
			boards = append(boards, Board{sum: boardSum})
		}
	}

	// call numbers in order, check for bingos
	winningNumber := 0
	winningBoard := 0
	bingo := false
	for _, num := range numbers {
		if bingo {
			break
		}
		intNum, _ := strconv.Atoi(num)
		for _, coords := range boardMap[num] {
			fmt.Printf("Num: %v, Coords: %v\n", num, coords)
			b := coords / 100
			x := digitAtPlace(coords, 2) - 1
			y := digitAtPlace(coords, 1) - 1
			fmt.Printf("1 ~ B:%v, X :%v, Y:%v\n", b, x, y)

			// mark a hit on the board
			boards[b].hits[x][y] = 1
			boards[b].sum -= intNum

			fmt.Printf("1 ~ Board: %v, Remaining sum: %v\n", b, boards[b].sum)
			fmt.Printf("1 ~ Board: %v, hits: %v\n", b, boards[b].hits)

			// check for bingo
			xbingo := 0
			ybingo := 0
			for i := 0; i < 5; i++ {
				xbingo += boards[b].hits[i][y]
				ybingo += boards[b].hits[x][i]
				fmt.Printf("x: %v, y: %v\n", boards[b].hits[x][i], boards[b].hits[i][y])
			}
			fmt.Printf("1 ~ x: %v, XBingo: %v, y: %v, YBingo: %v\n", x, xbingo, y, ybingo)
			if xbingo == 5 || ybingo == 5 {
				winningNumber = intNum
				winningBoard = b
				bingo = true
				break
			}
		}
	}

	// results
	fmt.Printf("1 ~ Winning Number: %v\n", winningNumber)
	fmt.Printf("1 ~ Board: %v, remaining sum: %v\n", winningBoard, boards[winningBoard].sum)
	fmt.Printf("1 ~ Result: %v\n", winningNumber*boards[winningBoard].sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Board struct {
	hits [5][5]int
	sum  int
}

func digitAtPlace(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
