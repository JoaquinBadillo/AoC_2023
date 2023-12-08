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

func splitPair(str string, sep string) (string, string, error) {
	vals := strings.Split(str, sep)

	if len(vals) != 2 {
		return "", "", errors.New("invalid size")
	}

	return vals[0], vals[1], nil
}

func validSet(cubeSet string, conditions map[string]int) bool {
	for _, cubes := range strings.Split(cubeSet, ",") {
		cubes = strings.TrimLeft(cubes, " ")
		qStr, color, err := splitPair(cubes, " ")

		if err != nil {
			log.Printf("Failed to split [%s] to a quantity and color pair", cubes)
			return false
		}

		color = strings.Trim(color, " ")
		limit, ok := conditions[color]

		if !ok {
			log.Printf("%s not found in map", color)
			return false
		}

		quantity, err := strconv.Atoi(qStr)

		if err != nil {
			log.Printf("Failed to cast [%s] to a number", qStr)
			return false
		}

		if quantity > limit {
			return false
		}
	}

	return true
}

func validGame(cubeSets []string, conditions map[string]int) bool {
	for _, cubeSet := range cubeSets {
		if !validSet(cubeSet, conditions) {
			return false
		}
	}

	return true
}

func addValidGames(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	conditions := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimRight(line, "\n") == "#" {
			break
		}

		color, qStr, err := splitPair(line, ",")

		if err != nil {
			log.Fatal("Invalid input file, failed to parse initial cubes")
		}

		val, err := strconv.Atoi(strings.TrimRight(qStr, "\n"))

		if err != nil {
			log.Fatal("Invalid input file")

		}

		conditions[color] = val
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game, data, err := splitPair(line, ":")

		if err != nil {
			log.Fatal("Invalid input")
		}

		num, err := strconv.Atoi(strings.Trim(game, "Game "))

		if err != nil {
			log.Fatal("Invalid input")
		}

		cubeSets := strings.Split(data, ";")

		if validGame(cubeSets, conditions) {
			sum += num
		}
	}

	return sum
}

func main() {
	var filename string

	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "input.txt"
	}

	fmt.Println(addValidGames(filename))
}
