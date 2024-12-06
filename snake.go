package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

var fieldRow = 10
var fieldCol = 20
var snakeLen = 6
var direction = "w"

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

func renderField(snakePoses [][2]int) {

	for i := 0; i < fieldRow; i++ {
		fmt.Println("_" + strings.Repeat("__", fieldCol-1))
		for j := 0; j < fieldCol; j++ {
			symbol := ""
			if !comparePoses(snakePoses, [2]int{i, j}) {
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
	fmt.Println("_" + strings.Repeat("__", fieldCol-1))
}

func changeSlice(a [][2]int, b [][2]int) {
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

func moveSnake(snakePoses [][2]int) [][2]int {
	switch direction {
	case "w":

		nextPose := border([2]int{snakePoses[len(snakePoses)-1][0] - 1, snakePoses[len(snakePoses)-1][1]})
		changeSlice(snakePoses, append(snakePoses[1:], nextPose))
	case "a":
		nextPose := border([2]int{snakePoses[len(snakePoses)-1][0], snakePoses[len(snakePoses)-1][1] - 1})
		changeSlice(snakePoses, append(snakePoses[1:], nextPose))
	case "s":
		nextPose := border([2]int{snakePoses[len(snakePoses)-1][0] + 1, snakePoses[len(snakePoses)-1][1]})
		changeSlice(snakePoses, append(snakePoses[1:], nextPose))
	case "d":
		nextPose := border([2]int{snakePoses[len(snakePoses)-1][0], snakePoses[len(snakePoses)-1][1] + 1})
		changeSlice(snakePoses, append(snakePoses[1:], nextPose))
	}
	return snakePoses
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
			direction = "w"
		case 'a':
			direction = "a"
		case 's':
			direction = "s"
		case 'd':
			direction = "d"
		}
		if key == keyboard.KeyEsc {
			break
		}
	}
}

func main() {

	snakePoses, _ := initSnake()

	renderField(snakePoses)

	go handleInput()
	time.Sleep(time.Second * 2)

	for {
		fmt.Print("\033[H\033[2J")
		renderField(moveSnake(snakePoses))
		time.Sleep(time.Millisecond * 500)
	}
}
