package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var fieldRow = 10
var fieldCol = 20
var snakeLen = 6

func comparePoses(snakePoses [][2]int, posVar [2]int) bool {
	for _, val := range snakePoses {
		if val == posVar {
			return false
		}
	}
	return true
}

func initSnake() [][2]int {

	randRow := rand.Intn(fieldRow)
	randCol := rand.Intn(fieldCol)

	snakePoses := make([][2]int, snakeLen)
	snakePoses[0] = [2]int{randRow, randCol}

	for i := 1; i < snakeLen; i++ {
		pos := make([][2]int, 0)
		if snakePoses[i-1][0] != 0 {
			posVar := [2]int{snakePoses[i-1][0] - 1, snakePoses[i-1][1]}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][0] != fieldRow {
			posVar := [2]int{snakePoses[i-1][0] + 1, snakePoses[i-1][1]}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][1] != 0 {
			posVar := [2]int{snakePoses[i-1][0], snakePoses[i-1][1] - 1}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][1] != fieldCol {
			posVar := [2]int{snakePoses[i-1][0], snakePoses[i-1][1] + 1}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}

		rngPose := rand.Intn(len(pos))
		snakePoses[i] = pos[rngPose]
	}

	return snakePoses
}

func renderField(field [][]bool, snakePoses [][2]int) {

	for _, pose := range snakePoses {
		field[pose[0]][pose[1]] = true
	}

	for i := 0; i < fieldRow; i++ {
		fmt.Println("_" + strings.Repeat("__", fieldCol))
		for j := 0; j < fieldCol; j++ {
			symbol := ""
			if field[i][j] {
				symbol = "*"
			} else {
				symbol = " "
			}
			fmt.Printf("%s|", symbol)
		}
		fmt.Println()
	}
	fmt.Println("_" + strings.Repeat("__", fieldCol))
}

func main() {

	field := make([][]bool, fieldRow)
	for i := range field {
		field[i] = make([]bool, fieldCol)
	}

	snakePoses := initSnake()

	renderField(field, snakePoses)
}
