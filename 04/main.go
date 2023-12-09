package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func split2(s string, sep string) (string, string, error) {
	arr := strings.Split(s, sep)
	if len(arr) != 2 {
		return "", "", fmt.Errorf("split2: %s", s)
	}
	return arr[0], arr[1], nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	cards := make(map[int]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimLeft(line, "Card ")

		card, rest, err := split2(line, ": ")
		if err != nil {
			log.Fatal(err)
		}

		cardNum, err := strconv.Atoi(card)
		if err != nil {
			log.Fatal(err)
		}

		total++
		current := 1
		multiplier := 1

		if val, ok := cards[cardNum]; ok {
			multiplier += val
		}

		left, right, err := split2(rest, " | ")
		if err != nil {
			log.Fatal(err)
		}

		winning := make(map[int]struct{}, 0)

		for _, numStr := range strings.Split(left, " ") {
			if numStr == "" {
				continue
			}

			num, err := strconv.Atoi(numStr)

			if err != nil {
				continue
			}

			winning[num] = struct{}{}
		}

		for _, numStr := range strings.Split(right, " ") {
			if numStr == "" {
				continue
			}

			num, err := strconv.Atoi(numStr)

			if err != nil {
				continue
			}

			if _, ok := winning[num]; ok {
				if _, ok := cards[cardNum]; !ok {
					cards[cardNum+current] = multiplier
				} else {
					cards[cardNum+current] += multiplier
				}
				current++
				total += multiplier
			}
		}
	}

	fmt.Println(total)
}
