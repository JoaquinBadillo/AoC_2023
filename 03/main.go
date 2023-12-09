package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Position struct {
	x int
	y int
}

type Num struct {
	value  int
	pos    []Position
	isPart bool
}

type Asterisk struct {
	adjacentNums map[*Num]struct{}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	symbols := make(map[Position]struct{})
	asterisks := make(map[Position]*Asterisk)

	nums := make([]*Num, 0)
	row := 0
	sum := 0

	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		positions := make([]Position, 0)

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				sb.WriteString(string(line[i]))
				positions = append(positions, Position{row, i})
			} else {
				if sb.Len() > 0 {
					num, err := strconv.Atoi(sb.String())

					if err != nil {
						log.Fatal(err)
					}

					nums = append(nums, &Num{num, positions, false})
					sb.Reset()
					positions = make([]Position, 0)
				}

				if line[i] != '.' {
					if line[i] == '*' {
						asterisks[Position{row, i}] = &Asterisk{make(map[*Num]struct{}, 0)}
					} else {
						symbols[Position{row, i}] = struct{}{}
					}
				}
			}
		}

		// Don't forget end of line
		if sb.Len() > 0 {
			num, err := strconv.Atoi(sb.String())

			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, &Num{num, positions, false})
			sb.Reset()
		}

		row++
	}

	directions := [8]Position{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for _, num := range nums {
		if num.isPart {
			continue
		}

		for _, pos := range num.pos {
			for _, direction := range directions {
				neihbour := Position{pos.x + direction.x, pos.y + direction.y}
				if _, ok := symbols[neihbour]; ok {
					if !num.isPart {
						num.isPart = true
						sum += num.value
					}
				} else if asterisk, ok := asterisks[neihbour]; ok {
					if _, ok := asterisk.adjacentNums[num]; !ok {
						asterisk.adjacentNums[num] = struct{}{}
					}

					if !num.isPart {
						num.isPart = true
						sum += num.value
					}
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sum)

	gearRatios := 0
	for _, asterisk := range asterisks {
		if len(asterisk.adjacentNums) == 2 {
			product := 1
			for num := range asterisk.adjacentNums {
				product *= num.value
			}
			gearRatios += product
		}
	}

	fmt.Printf("Part 2: %d\n", gearRatios)
}
