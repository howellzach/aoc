package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func MakeGroups(input []int) []int {
	c := 0
	var groups []int
	for i := 1; i < len(input)-1; i++ {
		group := input[c] + input[c+1] + input[c+2]
		c++
		groups = append(groups, group)
	}
	return groups
}

func CountIncreases(input []int) int {
	count := 0
	for i, v := range input[:len(input)-1] {
		val := v
		nextval := input[i+1]
		if val < nextval {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums, err := ReadInts(file)

	count1 := CountIncreases(nums)
	fmt.Printf("Result 1: %v\n", count1)

	groups := MakeGroups(nums)
	count2 := CountIncreases(groups)
	fmt.Printf("Result 2: %v\n", count2)
}
