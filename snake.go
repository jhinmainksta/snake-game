package main

import (
	"fmt"
	"log"
	"math/rand"
	"snake/utils"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

var fieldRow = 10
var fieldCol = 20
var snakeLen = 6
var direction = ""
var selectedDirection = ""
var snakePos = make([][2]int, snakeLen)
var food [2]int
var gameSpeed = 150
var gameScore = 0

func moveIsPossible(Pos [][2]int) bool {

	pos := make([][2]int, 0)
	if Pos[0][0] != 0 {
		posVar := [2]int{Pos[0][0] - 1, Pos[0][1]}
		if !utils.ContainPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Pos[0][0] != fieldRow-1 {
		posVar := [2]int{Pos[0][0] + 1, Pos[0][1]}
		if !utils.ContainPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Pos[0][1] != 0 {
		posVar := [2]int{Pos[0][0], Pos[0][1] - 1}
		if !utils.ContainPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Pos[0][1] != fieldCol-1 {
		posVar := [2]int{Pos[0][0], Pos[0][1] + 1}
		if !utils.ContainPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}
	if len(pos) > 1 {
		return true
	} else {
		return false
	}
}

func initFood() {

	food[0] = rand.Intn(fieldRow)
	food[1] = rand.Intn(fieldCol)

	for utils.ContainPos(snakePos, food) {
		food[0] = rand.Intn(fieldRow)
		food[1] = rand.Intn(fieldCol)
	}
}

func initSnake() {

	randRow := rand.Intn(fieldRow/2) + fieldRow/4
	randCol := rand.Intn(fieldCol/2) + fieldCol/4

	snakePos[0] = [2]int{randRow, randCol}

	for i := 1; i < snakeLen; i++ {
		pos := make([][2]int, 0)
		if snakePos[i-1][0] != 0 {
			posVar := [2]int{snakePos[i-1][0] - 1, snakePos[i-1][1]}
			curPos := append([][2]int{posVar}, snakePos...)
			if !(utils.ContainPos(snakePos, posVar)) && moveIsPossible(curPos) {
				pos = append(pos, posVar)
			}
		}
		if snakePos[i-1][0] != fieldRow-1 {
			posVar := [2]int{snakePos[i-1][0] + 1, snakePos[i-1][1]}
			curPos := append([][2]int{posVar}, snakePos...)
			if !(utils.ContainPos(snakePos, posVar)) && moveIsPossible(curPos) {
				pos = append(pos, posVar)
			}
		}
		if snakePos[i-1][1] != 0 {
			posVar := [2]int{snakePos[i-1][0], snakePos[i-1][1] - 1}
			curPos := append([][2]int{posVar}, snakePos...)
			if !(utils.ContainPos(snakePos, posVar)) && moveIsPossible(curPos) {
				pos = append(pos, posVar)
			}
		}
		if snakePos[i-1][1] != fieldCol-1 {
			posVar := [2]int{snakePos[i-1][0], snakePos[i-1][1] + 1}
			curPos := append([][2]int{posVar}, snakePos...)
			if !(utils.ContainPos(snakePos, posVar)) && moveIsPossible(curPos) {
				pos = append(pos, posVar)
			}
		}
		rngPos := rand.Intn(len(pos))
		snakePos[i] = pos[rngPos]
	}

	dx := snakePos[snakeLen-1][1] - snakePos[snakeLen-2][1]
	dy := snakePos[snakeLen-1][0] - snakePos[snakeLen-2][0]

	if dx == 1 {
		direction = "d"
	}
	if dx == -1 {
		direction = "a"
	}
	if dy == 1 {
		direction = "s"
	}
	if dy == -1 {
		direction = "w"
	}
}

func border(pos [2]int) [2]int {
	if pos[0] == fieldRow {
		pos[0] = 0
	}
	if pos[1] == fieldCol {
		pos[1] = 0
	}
	if pos[0] == -1 {
		pos[0] = fieldRow - 1
	}
	if pos[1] == -1 {
		pos[1] = fieldCol - 1
	}
	return pos
}

func moveSnake() {
	switch direction {
	case "w":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0] - 1, snakePos[len(snakePos)-1][1]})
		utils.UpdSlice(snakePos, append(snakePos[1:], nextPos))
	case "a":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0], snakePos[len(snakePos)-1][1] - 1})
		utils.UpdSlice(snakePos, append(snakePos[1:], nextPos))
	case "s":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0] + 1, snakePos[len(snakePos)-1][1]})
		utils.UpdSlice(snakePos, append(snakePos[1:], nextPos))
	case "d":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0], snakePos[len(snakePos)-1][1] + 1})
		utils.UpdSlice(snakePos, append(snakePos[1:], nextPos))
	}
}

func growSnake() {
	switch direction {
	case "w":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0] - 1, snakePos[len(snakePos)-1][1]})
		snakePos = append(snakePos, nextPos)
	case "a":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0], snakePos[len(snakePos)-1][1] - 1})
		snakePos = append(snakePos, nextPos)
	case "s":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0] + 1, snakePos[len(snakePos)-1][1]})
		snakePos = append(snakePos, nextPos)
	case "d":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0], snakePos[len(snakePos)-1][1] + 1})
		snakePos = append(snakePos, nextPos)
	}

	snakeLen += 1
}

func handleInput() {

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'w':
			selectedDirection = "w"
		case 'a':
			selectedDirection = "a"
		case 's':
			selectedDirection = "s"
		case 'd':
			selectedDirection = "d"
		}
		if key == keyboard.KeyEsc {
			break
		}
	}
}

func setDirection() {
	if selectedDirection != "" && !(selectedDirection == "w" && direction == "s" ||
		selectedDirection == "s" && direction == "w" ||
		selectedDirection == "a" && direction == "d" ||
		selectedDirection == "d" && direction == "a") {
		direction = selectedDirection
	}

}

func renderField() {

	scoreStr := strconv.Itoa(gameScore)
	fmt.Println(strings.Repeat(" ", 2*fieldCol+1-len(scoreStr)) + scoreStr)

	fmt.Println("┌─" + strings.Repeat("──", fieldCol) + "┐")
	for i := 0; i < fieldRow; i++ {
		fmt.Print("│ ")

		for j := 0; j < fieldCol; j++ {
			symbol := ""
			if utils.ContainPos(snakePos, [2]int{i, j}) {
				if [2]int{i, j} == snakePos[snakeLen-1] {
					symbol = "☭"
				} else {
					symbol = "∼"
				}
			} else {
				symbol = " "
			}
			if food == [2]int{i, j} {
				symbol = "*"
			}
			fmt.Printf("%s ", symbol)
		}

		fmt.Print("│")
		fmt.Println()
	}
	fmt.Print("└─" + strings.Repeat("──", fieldCol) + "┘")
}

func main() {
	// fmt.Print("How fast would you like to play ( In miliseconds. More miliseconds, slower the game): ")
	// fmt.Scan(&gameSpeed)

	screen.Clear()
	screen.MoveTopLeft()
	fieldCol = 10

	initSnake()
	initFood()
	renderField()

	go handleInput()
	time.Sleep(time.Second * 2)

	for {
		screen.Clear()
		screen.MoveTopLeft()
		setDirection()
		if food == [2]int{-1, -1} {
			initFood()
			growSnake()
			gameScore += 1
		} else {
			moveSnake()
		}

		if snakePos[snakeLen-1] == food {
			food = [2]int{-1, -1}
		}

		renderField()

		if utils.ContainPos(snakePos[:snakeLen-1], snakePos[snakeLen-1]) {
			fmt.Println()
			fmt.Println("You are proigral, prostofilya!")
			break
		}
		time.Sleep(time.Millisecond * time.Duration(gameSpeed))
	}
}
