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

type str struct {
	balance int
	Name    string
}

func newStr(balance int, Name string) *str {
	return &str{balance: balance, Name: Name}
}

var fieldRow = 8
var fieldCol = 8
var snakeLen = 6
var direction = ""
var selectedDirections [2]string
var snakePos = make([][2]int, snakeLen)
var food [2]int
var gameSpeed = 170
var gameScore = 0

var spaceChan = make(chan bool)

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

		if key == keyboard.KeySpace {
			spaceChan <- true
		}

		switch {
		case char == 'w' || key == keyboard.KeyArrowUp:
			if selectedDirections[0] != "" {
				selectedDirections[1] = "w"
			} else if direction != "w" {
				selectedDirections[0] = "w"
			}
		case char == 'a' || key == keyboard.KeyArrowLeft:
			if selectedDirections[0] != "" {
				selectedDirections[1] = "a"
			} else if direction != "a" {
				selectedDirections[0] = "a"
			}
		case char == 's' || key == keyboard.KeyArrowDown:
			if selectedDirections[0] != "" {
				selectedDirections[1] = "s"
			} else if direction != "s" {
				selectedDirections[0] = "s"
			}
		case char == 'd' || key == keyboard.KeyArrowRight:
			if selectedDirections[0] != "" {
				selectedDirections[1] = "d"
			} else if direction != "d" {
				selectedDirections[0] = "d"
			}
		}
		if key == keyboard.KeyEsc {
			break
		}
	}
}

func setDirection() {
	if selectedDirections[0] != "" {
		if !(selectedDirections[0] == "w" && direction == "s" ||
			selectedDirections[0] == "s" && direction == "w" ||
			selectedDirections[0] == "a" && direction == "d" ||
			selectedDirections[0] == "d" && direction == "a") {
			direction = selectedDirections[0]
		} else {
			selectedDirections[1] = ""
		}
		selectedDirections[0] = selectedDirections[1]
		selectedDirections[1] = ""
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
	fmt.Println("└─" + strings.Repeat("──", fieldCol) + "┘")
}

func main() {

	screen.Clear()
	screen.MoveTopLeft()

	initSnake()
	initFood()
	renderField()

	go handleInput()
	time.Sleep(time.Second * 2)

	isPaused := false

	for {
		select {
		case <-spaceChan:
			isPaused = !isPaused
			if isPaused == true {
				fmt.Println("pause")
			}
		default:
			if !isPaused {
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
					fmt.Println("You are proigral, prostofilya!")
					return
				}
				time.Sleep(time.Millisecond * time.Duration(gameSpeed))
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

}
