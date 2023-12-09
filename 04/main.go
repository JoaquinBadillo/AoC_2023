package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
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

	reg := regexp.MustCompile(`^Card \d+:\s`)

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = reg.ReplaceAllString(line, "${1}")

		left, right, err := split2(line, " | ")
		if err != nil {
			log.Fatal(err)
		}

		winning := make(map[int]struct{}, 0)
		winningCards := 0

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
				winningCards++
			}
		}

		prod := math.Pow(2, float64(winningCards)-1)

		if winningCards >= 1 {
			total += int(prod)
		}
	}

	fmt.Println(total)
}
