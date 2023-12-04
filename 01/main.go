/*
	Advent of Code 2023: Day 1
	Trebuchet?!
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	first  int
	second int
}

func kmpPreprocess(pattern string) []int {
	n := len(pattern)
	border := make([]int, n+1)

	border[0] = -1
	i := 0
	j := -1

	for i < n {
		for j >= 0 && pattern[i] != pattern[j] {
			j = border[j]
		}
		i++
		j++
		border[i] = j
	}

	return border
}

func kmp(text string, pattern string) (bool, int) {
	n := len(text)
	m := len(pattern)
	border := kmpPreprocess(pattern)

	i := 0
	j := 0

	for i < n {
		for j >= 0 && text[i] != pattern[j] {
			j = border[j]
		}
		i++
		j++

		if j == m {
			j = border[j]
			return true, i
		}
	}

	return false, -1
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := new(Pair)
		count := 0

		var sb strings.Builder

		for _, char := range line {
			num, err := strconv.Atoi(string(char))

			if err != nil {
				sb.WriteRune(char)
				match := false
				for key, value := range numbers {
					found, start := kmp(sb.String(), key)

					if found {
						num = value
						match = true

						prev := sb.String()
						sb.Reset()
						sb.WriteString(prev[start-1:])
						break
					}
				}

				if !match {
					continue
				}
			} else {
				sb.Reset()
			}

			if count == 0 {
				digits.first = num
				count++
			}
			digits.second = num
		}

		sum += digits.first*10 + digits.second
	}

	fmt.Println(sum)
}
