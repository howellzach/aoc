package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func bitGroup(input [][]string) (output [][]string) {
	for j := 0; j < len(input[0]); j++ {
		var bitgrp []string
		for i := 0; i < len(input); i++ {
			k := input[i][j]
			bitgrp = append(bitgrp, k)
		}
		output = append(output, bitgrp)
	}
	return
}

func convertBinStrToDec(binstr string) (num int) {
	x, err := strconv.ParseInt(binstr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	num = int(x)
	return
}

func createGamaEpsil(groups [][]string) (gamastring string, epsilstring string) {
	for _, group := range groups {
		ones := 0
		zeros := 0

		for _, v := range group {
			if v == "1" {
				ones++
			} else {
				zeros++
			}
		}

		if ones > zeros {
			gamastring += "1"
			epsilstring += "0"
		} else {
			gamastring += "0"
			epsilstring += "1"
		}
	}
	return
}

func filterLines(input [][]string, criteria string, pos int) (output [][]string) {
	bitGroup := bitGroup(input)

	ones := 0
	zeros := 0

	for _, v := range bitGroup[pos] {
		if v == "1" {
			ones++
		} else {
			zeros++
		}
	}

	if criteria == "most" {
		if ones >= zeros {
			for x := 0; x < len(input); x++ {
				if input[x][pos] == "1" {
					output = append(output, input[x])
				}
			}

		} else {
			for x := 0; x < len(input); x++ {
				if input[x][pos] == "0" {
					output = append(output, input[x])
				}
			}
		}
	} else if criteria == "least" {
		if zeros <= ones {
			for x := 0; x < len(input); x++ {
				if input[x][pos] == "0" {
					output = append(output, input[x])
				}
			}
		} else {
			for x := 0; x < len(input); x++ {
				if input[x][pos] == "1" {
					output = append(output, input[x])
				}
			}
		}
	}

	return
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

	var lines [][]string
	for i := range obj {
		x := strings.Split(obj[i], "")
		lines = append(lines, x)
	}

	groups := bitGroup(lines)

	gamastring, epsilstring := createGamaEpsil(groups)

	fmt.Printf("Gamastring: %s\n", gamastring)
	fmt.Printf("Epsilstring: %s\n", epsilstring)
	fmt.Println("")

	gamaint := convertBinStrToDec(gamastring)
	epsilint := convertBinStrToDec(epsilstring)

	fmt.Printf("Gamastring: %d\n", gamaint)
	fmt.Printf("Epsilstring: %d\n", epsilint)
	fmt.Println("")

	result := gamaint * epsilint
	fmt.Printf("Result 1: %d\n", result)

	most := filterLines(lines, "most", 0)
	mrepeats := len(most)
	mcount := 1

	for mrepeats > 1 {
		most = filterLines(most, "most", mcount)
		mrepeats = len(most)
		mcount++
	}

	least := filterLines(lines, "least", 0)
	lrepeats := len(least)
	lcount := 1

	for lrepeats > 1 {
		least = filterLines(least, "least", lcount)
		lrepeats = len(least)
		lcount++
	}

	var mostStr string
	for _, v := range most[0] {
		mostStr += v
	}

	var leastStr string
	for _, v := range least[0] {
		leastStr += v
	}

	fmt.Println("")
	oxyGenRating := convertBinStrToDec(mostStr)
	co2ScrubRating := convertBinStrToDec(leastStr)
	fmt.Printf("Oxygen Generator Rating: %v\n", oxyGenRating)
	fmt.Printf("CO2 Scrubber Rating: %v\n", co2ScrubRating)

	fmt.Println("")
	lifeSupportRating := oxyGenRating * co2ScrubRating
	fmt.Printf("Life Support Rating (Result 2): %v\n", lifeSupportRating)
}
