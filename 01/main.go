package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pair struct {
	first  int
	second int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := new(Pair)
		count := 0

		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				continue
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
