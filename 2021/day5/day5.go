package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func createOceanFloor(maxX int, maxY int) [][]int {
	var grid [][]int
	for i := 0; i < maxY+1; i++ {
		line := []int{}
		for j := 0; j < maxX+1; j++ {
			line = append(line, 0)
		}
		grid = append(grid, line)
	}
	return grid
}

func addVentsVertHori(lines [][][]int, grid [][]int) [][]int {
	var diff int
	var start int

	for i := 0; i < len(lines); i++ {
		if lines[i][0][0] == lines[i][1][0] {
			// Horizontal line
			firY := lines[i][0][1]
			secY := lines[i][1][1]
			xcol := lines[i][0][0]

			if firY < secY {
				diff = secY - firY + 1
				start = firY
			} else {
				diff = firY - secY + 1
				start = secY
			}

			for a := start; a < start+diff; a++ {
				grid[a][xcol] = grid[a][xcol] + 1
			}

		} else if lines[i][0][1] == lines[i][1][1] {
			// Vertical line
			firX := lines[i][0][0]
			secX := lines[i][1][0]
			yrow := lines[i][0][1]

			if firX < secX {
				diff = secX - firX + 1
				start = firX
			} else {
				diff = firX - secX + 1
				start = secX
			}
			for a := start; a < start+diff; a++ {
				grid[yrow][a] = grid[yrow][a] + 1
			}
		}
	}
	return grid
}

func addVentsDiag(lines [][][]int, grid [][]int) [][]int {
	var diff int
	var start int

	for i := 0; i < len(lines); i++ {
		if lines[i][0][0] != lines[i][1][0] && lines[i][0][1] != lines[i][1][1] {
			// Not vertical or horizontal = diagonal
			firX := lines[i][0][0]
			firY := lines[i][0][1]
			secX := lines[i][1][0]
			secY := lines[i][1][1]

			if firX < secX {
				start = firX
				diff = secX - firX + 1
				y := firY
				if secY > firY {
					for a := start; a < start+diff; a++ {
						grid[y][a] = grid[y][a] + 1
						y++
					}
				} else {
					for a := start; a < start+diff; a++ {
						grid[y][a] = grid[y][a] + 1
						y--
					}
				}
			} else {
				start = secX
				diff = firX - secX + 1
				y := secY
				if secY > firY {
					for a := start; a < start+diff; a++ {
						grid[y][a] = grid[y][a] + 1
						y--
					}
				} else {
					for a := start; a < start+diff; a++ {
						grid[y][a] = grid[y][a] + 1
						y++
					}
				}

			}

		}
	}
	return grid
}

func countOverlap(grid [][]int) int {
	count := 0
	for _, v := range grid {
		for _, k := range v {
			if k > 1 {
				count++
			}
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

	scanner := bufio.NewScanner(file)

	var obj []string
	for scanner.Scan() {
		obj = append(obj, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var lines [][]string
	for _, v := range obj {
		s := strings.Split(v, " -> ")
		lines = append(lines, s)
	}

	var seplines [][]int
	for _, set := range lines {
		for _, v := range set {
			x := strings.Split(v, ",")
			grp := []int{}
			for _, y := range x {
				z, _ := strconv.Atoi(y)
				grp = append(grp, z)
			}
			seplines = append(seplines, grp)
		}
	}

	maxX := 0
	maxY := 0
	for _, j := range seplines {
		if j[0] > maxX {
			maxX = j[0]
		}
		if j[1] > maxY {
			maxY = j[1]
		}
	}

	grplines := [][][]int{}
	for i := 0; i < len(seplines); i++ {
		var grp [][]int
		for j := 0; j < 2; j++ {
			grp = [][]int{seplines[i], seplines[i+1]}
			j++
		}
		grplines = append(grplines, grp)
		i++
	}

	oceanFloor := createOceanFloor(maxX, maxY)

	addVentsVertHori(grplines, oceanFloor)
	fmt.Printf("Result 1: %v\n", countOverlap(oceanFloor))

	addVentsDiag(grplines, oceanFloor)
	fmt.Printf("Result 2: %v\n", countOverlap(oceanFloor))
}
