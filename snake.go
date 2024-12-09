package main

import (
	"fmt"
	"log"
	"math/rand"
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
var snakeSkin = make([]string, snakeLen)

func containPos(snakePos [][2]int, posVar [2]int) bool {
	for _, val := range snakePos {
		if val == posVar {
			return false
		}
	}
	return true
}

func moveIsPossible(Pos [][2]int) bool {

	pos := make([][2]int, 0)
	if Pos[0][0] != 0 {
		posVar := [2]int{Pos[0][0] - 1, Pos[0][1]}
		if containPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Pos[0][0] != fieldRow-1 {
		posVar := [2]int{Pos[0][0] + 1, Pos[0][1]}
		if containPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Pos[0][1] != 0 {
		posVar := [2]int{Pos[0][0], Pos[0][1] - 1}
		if containPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}

	if Pos[0][1] != fieldCol-1 {
		posVar := [2]int{Pos[0][0], Pos[0][1] + 1}
		if containPos(Pos, posVar) {
			pos = append(pos, posVar)
		}
	}
	if len(pos) > 1 {
		return true
	} else {
		return false
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
			if containPos(snakePos, posVar) && moveIsPossible(curPos) {
				pos = append(pos, posVar)
			}
		}
		if snakePos[i-1][0] != fieldRow-1 {
			posVar := [2]int{snakePos[i-1][0] + 1, snakePos[i-1][1]}
			curPos := append([][2]int{posVar}, snakePos...)
			if containPos(snakePos, posVar) && moveIsPossible(curPos) {
				pos = append(pos, posVar)
			}
		}
		if snakePos[i-1][1] != 0 {
			posVar := [2]int{snakePos[i-1][0], snakePos[i-1][1] - 1}
			curPos := append([][2]int{posVar}, snakePos...)
			if containPos(snakePos, posVar) && moveIsPossible(curPos) {
				pos = append(pos, posVar)
			}
		}
		if snakePos[i-1][1] != fieldCol-1 {
			posVar := [2]int{snakePos[i-1][0], snakePos[i-1][1] + 1}
			curPos := append([][2]int{posVar}, snakePos...)
			if containPos(snakePos, posVar) && moveIsPossible(curPos) {

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

func renderField() {

	fmt.Println("┌─" + strings.Repeat("──", fieldCol) + "┐")
	for i := 0; i < fieldRow; i++ {
		fmt.Print("│ ")

		for j := 0; j < fieldCol; j++ {
			symbol := ""
			if !containPos(snakePos, [2]int{i, j}) {
				symbol = "*"
			} else {
				symbol = " "
			}
			fmt.Printf("%s ", symbol)
		}

		fmt.Print("│")
		fmt.Println()
	}
	fmt.Print("└─" + strings.Repeat("──", fieldCol) + "┘")
}

func updateSnake(a [][2]int, b [][2]int) {
	for i, _ := range a {
		a[i] = b[i]
	}
}

func border(pos [2]int) [2]int {
	if pos[0] == 10 {
		pos[0] = 0
	}
	if pos[1] == 20 {
		pos[1] = 0
	}
	if pos[0] == -1 {
		pos[0] = 9
	}
	if pos[1] == -1 {
		pos[1] = 19
	}
	return pos
}

func moveSnake() {
	switch direction {
	case "w":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0] - 1, snakePos[len(snakePos)-1][1]})
		updateSnake(snakePos, append(snakePos[1:], nextPos))
	case "a":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0], snakePos[len(snakePos)-1][1] - 1})
		updateSnake(snakePos, append(snakePos[1:], nextPos))
	case "s":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0] + 1, snakePos[len(snakePos)-1][1]})
		updateSnake(snakePos, append(snakePos[1:], nextPos))
	case "d":
		nextPos := border([2]int{snakePos[len(snakePos)-1][0], snakePos[len(snakePos)-1][1] + 1})
		updateSnake(snakePos, append(snakePos[1:], nextPos))
	}
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

func main() {

	initSnake()
	renderField()

	go handleInput()
	time.Sleep(time.Second * 2)

	for {
		screen.Clear()
		screen.MoveTopLeft()
		setDirection()
		moveSnake()
		renderField()
		time.Sleep(time.Millisecond * 500)
	}
}
