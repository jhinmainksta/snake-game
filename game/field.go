package game

import (
	"fmt"
	"math/rand"
	"snake/utils"
	"strconv"
	"strings"
)

type Snake struct {
	direction string
	poses     [][2]int
	len       int
}

type Field struct {
	row   int
	col   int
	score int
	food  [2]int
	snake *Snake
}

func (f *Field) initSnake(length int) {

	randRow := rand.Intn(f.row/2) + f.row/4
	randCol := rand.Intn(f.col/2) + f.col/4

	snakePos := make([][2]int, length)
	snakePos[0] = [2]int{randRow, randCol}

	for i := 1; i < length; i++ {
		poses := make([][2]int, 0)
		if snakePos[i-1][0] != 0 {
			newPos := [2]int{snakePos[i-1][0] - 1, snakePos[i-1][1]}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, f.col, f.row) {
				poses = append(poses, newPos)
			}
		}

		if snakePos[i-1][0] != f.row-1 {
			newPos := [2]int{snakePos[i-1][0] + 1, snakePos[i-1][1]}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, f.col, f.row) {
				poses = append(poses, newPos)
			}
		}

		if snakePos[i-1][1] != 0 {
			newPos := [2]int{snakePos[i-1][0], snakePos[i-1][1] - 1}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, f.col, f.row) {
				poses = append(poses, newPos)
			}
		}

		if snakePos[i-1][1] != f.col-1 {
			newPos := [2]int{snakePos[i-1][0], snakePos[i-1][1] + 1}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, f.col, f.row) {
				poses = append(poses, newPos)
			}
		}

		rngPos := rand.Intn(len(poses))
		snakePos[i] = poses[rngPos]
	}

	dx := snakePos[length-1][1] - snakePos[length-2][1]
	dy := snakePos[length-1][0] - snakePos[length-2][0]

	direction := ""

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

	f.snake = &Snake{
		direction: direction,
		poses:     snakePos,
		len:       length,
	}
}

func (f *Field) moveSnake() {
	switch f.snake.direction {
	case "w":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0] - 1, f.snake.poses[len(f.snake.poses)-1][1]}
		nextPos = border(nextPos, f.row, f.col)
		utils.UpdSlice(f.snake.poses, append(f.snake.poses[1:], nextPos))
	case "a":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0], f.snake.poses[len(f.snake.poses)-1][1] - 1}
		nextPos = border(nextPos, f.row, f.col)
		utils.UpdSlice(f.snake.poses, append(f.snake.poses[1:], nextPos))
	case "s":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0] + 1, f.snake.poses[len(f.snake.poses)-1][1]}
		nextPos = border(nextPos, f.row, f.col)
		utils.UpdSlice(f.snake.poses, append(f.snake.poses[1:], nextPos))
	case "d":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0], f.snake.poses[len(f.snake.poses)-1][1] + 1}
		nextPos = border(nextPos, f.row, f.col)
		utils.UpdSlice(f.snake.poses, append(f.snake.poses[1:], nextPos))
	}
}

func (f *Field) growSnake() {

	switch f.snake.direction {
	case "w":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0] - 1, f.snake.poses[len(f.snake.poses)-1][1]}
		nextPos = border(nextPos, f.row, f.col)
		f.snake.poses = append(f.snake.poses, nextPos)
	case "a":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0], f.snake.poses[len(f.snake.poses)-1][1] - 1}
		nextPos = border(nextPos, f.row, f.col)
		f.snake.poses = append(f.snake.poses, nextPos)
	case "s":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0] + 1, f.snake.poses[len(f.snake.poses)-1][1]}
		nextPos = border(nextPos, f.row, f.col)
		f.snake.poses = append(f.snake.poses, nextPos)
	case "d":
		nextPos := [2]int{f.snake.poses[len(f.snake.poses)-1][0], f.snake.poses[len(f.snake.poses)-1][1] + 1}
		nextPos = border(nextPos, f.row, f.col)
		f.snake.poses = append(f.snake.poses, nextPos)
	}

	f.snake.len += 1
	f.score += 1
}

func (f *Field) initFood() {

	f.food[0] = rand.Intn(f.row)
	f.food[1] = rand.Intn(f.col)

	for utils.ContainPos(f.snake.poses, f.food) {
		f.food[0] = rand.Intn(f.row)
		f.food[1] = rand.Intn(f.col)
	}
}

func (f *Field) renderField() {
	scoreStr := strconv.Itoa(f.score)
	fmt.Println(strings.Repeat(" ", 2*f.col+1-len(scoreStr)) + scoreStr)

	fmt.Println("┌" + strings.Repeat("──", f.col) + "┐")

	for i := 0; i < f.row; i++ {

		fmt.Print("│")

		for j := 0; j < f.col; j++ {

			symbol := " "

			if f.food == [2]int{i, j} {
				symbol = "*"
			}

			if utils.ContainPos(f.snake.poses, [2]int{i, j}) {
				if [2]int{i, j} == f.snake.poses[f.snake.len-1] {
					symbol = "☭"
				} else {
					symbol = "∼"
				}
			}

			fmt.Printf("%s ", symbol)
		}

		fmt.Print("│")
		fmt.Println()

	}

	fmt.Println("└" + strings.Repeat("──", f.col) + "┘")
}

func (f *Field) ProcessTheMove() string {
	if f.snake.poses[f.snake.len-1] == f.food {
		f.growSnake()
		f.initFood()
	} else {
		f.moveSnake()
	}

	f.renderField()

	if utils.ContainPos(f.snake.poses[:f.snake.len-1], f.snake.poses[f.snake.len-1]) {
		return "You are proigral, prostofilya!"
	}
	return ""
}
