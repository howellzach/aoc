package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dive struct {
	direction string
	units     int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var obj []string
	for scanner.Scan() {
		obj = append(obj, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var travel []dive

	for i := range obj {
		x := strings.Fields(obj[i])
		dir := x[0]
		num, err := strconv.Atoi(x[1])
		if err != nil {
			log.Fatal(err)
		}
		z := dive{direction: dir, units: num}
		travel = append(travel, z)
	}

	hori1 := 0
	depth1 := 0

	for i := range travel {
		if travel[i].direction == "forward" {
			hori1 = hori1 + travel[i].units
		} else if travel[i].direction == "down" {
			depth1 = depth1 + travel[i].units
		} else if travel[i].direction == "up" {
			depth1 = depth1 - travel[i].units
		}
	}

	result1 := hori1 * depth1

	fmt.Printf("Horizontal: %d\n", hori1)
	fmt.Printf("Depth: %d\n", depth1)
	fmt.Println("")
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Println("")

	hori2 := 0
	depth2 := 0
	aim := 0

	for i := range travel {
		if travel[i].direction == "forward" {
			hori2 = hori2 + travel[i].units
			depth2 = depth2 + (aim * travel[i].units)
		} else if travel[i].direction == "down" {
			aim = aim + travel[i].units
		} else if travel[i].direction == "up" {
			aim = aim - travel[i].units
		}
	}

	result2 := hori2 * depth2

	fmt.Printf("Aim: %d\n", aim)
	fmt.Printf("Horizontal: %d\n", hori2)
	fmt.Printf("Depth: %d\n", depth2)
	fmt.Println("")
	fmt.Printf("Result 2: %d\n", result2)
}
