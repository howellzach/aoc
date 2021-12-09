package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type number struct {
	num    string
	marked int
}

func newRow(input []string) []number {
	var row []number
	for _, v := range input {
		row = append(row, number{num: v, marked: 0})
	}
	return row
}

func createAllBoards(numgroups [][]string) (allboards [][][]number) {
	var row []number
	var board [][]number

	boardnum := len(numgroups)
	x := 0
	for x < boardnum {
		row = newRow(numgroups[x])
		board = append(board, row)
		x++
		if x%5 == 0 {
			allboards = append(allboards, board)
			row = []number{}
			board = [][]number{}
		}
	}
	return allboards
}

func addDrawnNum(boards [][][]number, drawnNum string) [][][]number {
	for i, x := range boards {
		for j, y := range x {
			for k, z := range y {
				if z.num == drawnNum {
					boards[i][j][k].marked = 1
				}
			}
		}
	}
	return boards
}

func checkHorizontal(input [][][]number) (winnerFound bool, winningBoards []int) {
	for i, v := range input {
		board := v
		for _, row := range board {
			sum := 0
			for _, x := range row {
				sum += x.marked
				if sum == 5 {
					winnerFound = true
					winningBoards = append(winningBoards, i)
				}
			}
		}
	}
	return winnerFound, winningBoards
}

func checkVertical(input [][][]number) (winnerFound bool, winningBoards []int) {
	var verticalLines []number
	for i, v := range input {
		board := v
		for row := 0; row < len(board); row++ {
			for i := 0; i < len(board[0]); i++ {
				verticalLines = append(verticalLines, board[i][row])
			}
			sum := 0
			for _, num := range verticalLines {
				sum += num.marked

			}
			if sum == 5 {
				winnerFound = true
				winningBoards = append(winningBoards, i)
				break
			}
			verticalLines = []number{}
		}
	}
	return winnerFound, winningBoards
}

func countWinningBoards(allboards [][][]number, drawNums []string) int {
	var winning_board [][]number
	var winning_board_list [][][]number
	var winning_board_index_list []int

	for _, v := range drawNums {
		localboard := addDrawnNum(allboards, v)

		winner_found, winning_board_indexes := checkHorizontal(localboard)
		if winner_found == true {
			for _, index := range winning_board_indexes {
				winning_board = localboard[index]
				if contains(winning_board_index_list, index) == false {
					winning_board_list = append(winning_board_list, winning_board)
				}
			}
			for _, index := range winning_board_indexes {
				winning_board_index_list = append(winning_board_index_list, index)
			}

		}

		winner_found, winning_board_indexes = checkVertical(localboard)
		if winner_found == true {
			for _, index := range winning_board_indexes {
				winning_board = localboard[index]
				if contains(winning_board_index_list, index) == false {
					winning_board_list = append(winning_board_list, winning_board)
				}
			}
			for _, index := range winning_board_indexes {
				winning_board_index_list = append(winning_board_index_list, index)
			}
		}
	}
	winnerCount := len(winning_board_list)
	return winnerCount
}

func runBingo(allboards [][][]number, drawNums []string, winnerCount int) (result1 int, result2 int) {
	var last_drawn_num int
	var winning_board [][]number
	var winning_board_list [][][]number
	var winning_board_index_list []int
	var gotres1 bool
	var gotres2 bool

	for _, v := range drawNums {
		allboards = addDrawnNum(allboards, v)

		winner_found, winning_board_indexes := checkHorizontal(allboards)
		if winner_found == true {
			for _, index := range winning_board_indexes {
				winning_board = allboards[index]
				if contains(winning_board_index_list, index) == false {
					winning_board_list = append(winning_board_list, winning_board)
					winning_board_index_list = append(winning_board_index_list, index)
					last_drawn_num, _ = strconv.Atoi(v)
				}
			}
			if len(winning_board_list) == 1 && gotres1 == false {
				first_board := winning_board_list[0]
				result1 = getResult(first_board, last_drawn_num)
				gotres1 = true
			}
			if len(winning_board_list) == winnerCount && gotres2 == false {
				last_winning_board := winning_board_list[len(winning_board_list)-1]
				result2 = getResult(last_winning_board, last_drawn_num)
				gotres2 = true
			}
		}

		winner_found, winning_board_indexes = checkVertical(allboards)
		if winner_found == true {
			for _, index := range winning_board_indexes {
				winning_board = allboards[index]
				if contains(winning_board_index_list, index) == false {
					winning_board_list = append(winning_board_list, winning_board)
					winning_board_index_list = append(winning_board_index_list, index)
					last_drawn_num, _ = strconv.Atoi(v)
				}
			}
			if len(winning_board_list) == 1 && gotres1 == false {
				first_board := winning_board_list[0]
				result1 = getResult(first_board, last_drawn_num)
			}
			if len(winning_board_list) == winnerCount && gotres2 == false {
				last_winning_board := winning_board_list[len(winning_board_list)-1]
				result2 = getResult(last_winning_board, last_drawn_num)
				gotres2 = true
			}
		}

	}

	return result1, result2
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getResult(board [][]number, lastDrawnNum int) int {
	unmarkedSum := 0
	for i := 0; i < len(board); i++ {
		for _, y := range board[i] {
			if y.marked == 0 {
				v, _ := strconv.Atoi(y.num)
				unmarkedSum += v
			}
		}
	}
	return unmarkedSum * lastDrawnNum
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

	drawNums := strings.Split(obj[0], ",")

	var numstring []string
	var numgroups [][]string

	for _, v := range obj[2:] {
		if v == "" {
			continue
		} else {
			numstring = strings.Fields(v)
			numgroups = append(numgroups, numstring)
		}
	}

	allboards := createAllBoards(numgroups)
	allboards2 := createAllBoards(numgroups)

	winner_count := countWinningBoards(allboards, drawNums)
	result_1, result_2 := runBingo(allboards2, drawNums, winner_count)
	fmt.Printf("Result 1: %v\n", result_1)
	fmt.Printf("Result 2: %v\n", result_2)
}
