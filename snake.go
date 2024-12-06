package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var fieldRow = 10
var fieldCol = 20
var snakeLen = 6
var direction = "d"

func comparePoses(snakePoses [][2]int, posVar [2]int) bool {
	for _, val := range snakePoses {
		if val == posVar {
			return false
		}
	}
	return true
}

func moveIsPossible(Poses [][2]int) bool {

	pos := make([][2]int, 0)
	if Poses[0][0] != 0 {
		posVar := [2]int{Poses[0][0] - 1, Poses[0][1]}
		if comparePoses(Poses, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Poses[0][0] != fieldRow-1 {
		posVar := [2]int{Poses[0][0] + 1, Poses[0][1]}
		if comparePoses(Poses, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Poses[0][1] != 0 {
		posVar := [2]int{Poses[0][0], Poses[0][1] - 1}
		if comparePoses(Poses, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Poses[0][1] != fieldCol-1 {
		posVar := [2]int{Poses[0][0], Poses[0][1] + 1}
		if comparePoses(Poses, posVar) {
			pos = append(pos, posVar)
		}
	}
	if len(pos) > 1 {
		return true
	} else {
		return false
	}
}

func initSnake() ([][2]int, error) {

	randRow := rand.Intn(fieldRow/2) + fieldRow/4

	randCol := rand.Intn(fieldCol/2) + fieldCol/4

	snakePoses := make([][2]int, snakeLen)
	snakePoses[0] = [2]int{randRow, randCol}

	for i := 1; i < snakeLen; i++ {
		pos := make([][2]int, 0)
		if snakePoses[i-1][0] != 0 {
			posVar := [2]int{snakePoses[i-1][0] - 1, snakePoses[i-1][1]}
			curPoses := append([][2]int{posVar}, snakePoses...)
			if comparePoses(snakePoses, posVar) && moveIsPossible(curPoses) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][0] != fieldRow-1 {
			posVar := [2]int{snakePoses[i-1][0] + 1, snakePoses[i-1][1]}
			curPoses := append([][2]int{posVar}, snakePoses...)
			if comparePoses(snakePoses, posVar) && moveIsPossible(curPoses) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][1] != 0 {
			posVar := [2]int{snakePoses[i-1][0], snakePoses[i-1][1] - 1}
			curPoses := append([][2]int{posVar}, snakePoses...)
			if comparePoses(snakePoses, posVar) && moveIsPossible(curPoses) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][1] != fieldCol-1 {
			posVar := [2]int{snakePoses[i-1][0], snakePoses[i-1][1] + 1}
			curPoses := append([][2]int{posVar}, snakePoses...)
			if comparePoses(snakePoses, posVar) && moveIsPossible(curPoses) {

				pos = append(pos, posVar)
			}
		}
		rngPose := rand.Intn(len(pos))
		snakePoses[i] = pos[rngPose]
	}

	return snakePoses, nil
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
				if i == snakePoses[snakeLen-1][0] && j == snakePoses[snakeLen-1][1] {
					symbol = "A"
				} else {
					symbol = "*"
				}
			} else {
				symbol = " "
			}
			fmt.Printf("%s|", symbol)
		}
		fmt.Println()
	}
	fmt.Println("_" + strings.Repeat("__", fieldCol))
}

func moveSnake(snakePoses [][2]int) {
	switch direction {
	case "w":
		fmt.Print("w")
		fmt.Print("w")
	case "a":
		fmt.Print("a")
		fmt.Print("a")
	case "s":
		fmt.Print("s")
		fmt.Print("s")
	case "d":
		fmt.Print("d")
		fmt.Print("d")
	}
}

func main() {

	field := make([][]bool, fieldRow)
	for i := range field {
		field[i] = make([]bool, fieldCol)
	}

	snakePoses, _ := initSnake()
	fmt.Println(snakePoses)

	renderField(field, snakePoses)

	for i := 0; ; i++ {
		fmt.Println(i)
		snakePoses, err := initSnake()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(snakePoses)
		}
	}
}
