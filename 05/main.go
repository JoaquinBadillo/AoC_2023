package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SeedMapping struct {
	num     int
	changed bool
}

func parseMapping(line string) (int, int, int, error) {
	mapData := strings.Split(line, " ")

	if len(mapData) != 3 {
		return 0, 0, 0, errors.New("invalid mapping")
	}

	res := make([]int, 3)
	for i, s := range mapData {
		num, err := strconv.Atoi(s)

		if err != nil {
			return 0, 0, 0, errors.New("invalid mapping")
		}

		res[i] = num
	}

	return res[0], res[1], res[2], nil

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	seedData := make(map[int]*SeedMapping)

	line := scanner.Text()
	line = strings.TrimLeft(line, "seeds: ")
	for _, s := range strings.Split(line, " ") {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		seedData[num] = &SeedMapping{num: num, changed: false}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimRight(line, "\n") == "" {
			// Reset changed state
			for _, value := range seedData {
				value.changed = false
			}

			// Clear title
			scanner.Scan()
			continue
		}

		destination, source, length, err := parseMapping(line)

		if err != nil {
			log.Fatal(err)
		}

		for _, value := range seedData {
			if !value.changed && value.num >= source && value.num < source+length {
				value.changed = true
				value.num = destination + (value.num - source)
			}
		}
	}

	min := [2]int{-1, -1}

	for key, value := range seedData {
		if min[0] == -1 || value.num < min[0] {
			min[0] = value.num
			min[1] = key
		}
	}

	fmt.Println(min[0])
}
