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

// Part 1

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

func parseConditions(filename string) (map[string]int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	conditions := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
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

	return conditions, nil
}

func addValidGames(filename string, conditions map[string]int) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

// Part 2

func cubePow(cubeSets []string) (int, error) {
	minCubes := make(map[string]int)
	for _, cubeSet := range cubeSets {
		for _, cubes := range strings.Split(cubeSet, ",") {
			cubes = strings.TrimLeft(cubes, " ")
			qStr, color, err := splitPair(cubes, " ")

			if err != nil {
				return -1, errors.New("failed to split quantity and color pair")
			}

			color = strings.Trim(color, " ")
			quantity, err := strconv.Atoi(qStr)

			if err != nil {
				log.Printf("Failed to cast [%s] to a number", qStr)
				return -1, errors.New("failed to cast quantity to int")
			}

			val, ok := minCubes[color]

			if !ok {
				minCubes[color] = quantity
			} else {
				minCubes[color] = max(val, quantity)
			}
		}
	}

	product := 1

	for _, value := range minCubes {
		product *= value
	}

	return product, nil
}

func sumOfCubePows(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		_, data, err := splitPair(line, ":")

		if err != nil {
			log.Fatal("Invalid input")
		}

		cubeSets := strings.Split(data, ";")

		pow, err := cubePow(cubeSets)

		if err != nil {
			log.Fatal("Invalid input")
		}

		sum += pow
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

	// Part 1
	conditions, err := parseConditions("conditions.txt")
	if err != nil {
		log.Fatal("Failed to parse conditions")
	}

	fmt.Printf("Part 1: %d\n", addValidGames(filename, conditions))

	// Part 2

	fmt.Printf("Part 2: %d\n", sumOfCubePows(filename))
}
