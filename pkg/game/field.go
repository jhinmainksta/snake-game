package game

import (
	"math/rand"
	"snake/pkg/utils"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"
)

type Snake struct {
	direction string
	poses     [][2]int
	len       int
}

type Field struct {
	row        int
	col        int
	borderMode bool
	wasEaten   bool
	score      int
	food       [2]int
	snakeLen   int
	snake      *Snake
}

func (f *Field) initSnake() {

	randRow := rand.Intn(f.row/2) + f.row/4
	randCol := rand.Intn(f.col/2) + f.col/4

	snakePos := make([][2]int, f.snakeLen)
	snakePos[0] = [2]int{randRow, randCol}

	for i := 1; i < f.snakeLen; i++ {

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

	dx := snakePos[f.snakeLen-1][1] - snakePos[f.snakeLen-2][1]
	dy := snakePos[f.snakeLen-1][0] - snakePos[f.snakeLen-2][0]

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
		len:       f.snakeLen,
	}
}

func (f *Field) moveSnake() {

	nextPos := [2]int{}
	switch f.snake.direction {
	case "w":
		nextPos = [2]int{f.snake.poses[len(f.snake.poses)-1][0] - 1, f.snake.poses[len(f.snake.poses)-1][1]}
	case "a":
		nextPos = [2]int{f.snake.poses[len(f.snake.poses)-1][0], f.snake.poses[len(f.snake.poses)-1][1] - 1}
	case "s":
		nextPos = [2]int{f.snake.poses[len(f.snake.poses)-1][0] + 1, f.snake.poses[len(f.snake.poses)-1][1]}
	case "d":
		nextPos = [2]int{f.snake.poses[len(f.snake.poses)-1][0], f.snake.poses[len(f.snake.poses)-1][1] + 1}
	}

	if !f.borderMode {
		nextPos = border(nextPos, f.row, f.col)
	}

	if f.wasEaten {
		f.snake.poses = append(f.snake.poses, nextPos)
		f.snake.len++
		f.wasEaten = false
	} else {
		utils.UpdSlice(f.snake.poses, append(f.snake.poses[1:], nextPos))
	}

	if f.foodIsEaten() {
		f.wasEaten = true
		f.score += 1
		f.initFood()
	}
}

func (f *Field) foodIsEaten() bool {
	return f.snake.poses[f.snake.len-1] == f.food
}

func (f *Field) ProcessTheMove() bool {

	f.moveSnake()

	f.renderGame()

	snakeHead := f.snake.poses[f.snake.len-1]
	snakeHitItself := utils.ContainPos(f.snake.poses[:f.snake.len-1], snakeHead)
	snakeHitTheWall := false

	if f.borderMode {
		snakeHitTheWall = snakeHead[0] == -1 || snakeHead[0] == f.row || snakeHead[1] == -1 || snakeHead[1] == f.col
	}

	return snakeHitItself || snakeHitTheWall
}

func (f *Field) initFood() {

	f.food[0] = rand.Intn(f.row)
	f.food[1] = rand.Intn(f.col)

	for utils.ContainPos(f.snake.poses, f.food) {
		f.food[0] = rand.Intn(f.row)
		f.food[1] = rand.Intn(f.col)
	}
}

func (f *Field) renderInfo() {
	utils.Tbprint(f.col+1-3, f.row+3, defaultColour, defaultColour, "paused")
	termbox.Flush()
}

func (f *Field) renderScoreAndMode() {
	scoreStr := strconv.Itoa(f.score)
	msg := "borders off"
	if f.borderMode {
		msg = "borders on"
	}
	backspace := (f.col+1)*2 - len(scoreStr) - len(msg)
	utils.Tbprint(0, 0, defaultColour, defaultColour, msg+strings.Repeat(" ", backspace)+scoreStr)
}

func (f *Field) renderField() {

	utils.Tbprint(0, 1, defaultColour, defaultColour, "┌"+strings.Repeat("─", (f.col+1)*2-1)+"┐")

	for i := 1; i < f.row+1; i++ {
		termbox.SetCell(0, i+1, '│', defaultColour, defaultColour)
		termbox.SetCell((f.col+1)*2, i+1, '│', defaultColour, defaultColour)
	}

	termbox.SetCell((f.food[1]+1)*2, f.food[0]+2, '*', termbox.ColorRed, defaultColour)

	utils.Tbprint(0, f.row+2, defaultColour, defaultColour, "└"+strings.Repeat("─", (f.col+1)*2-1)+"┘")

	for i, pos := range f.snake.poses {
		char := '~'
		if i == f.snake.len-1 {
			char = '☭'
		}
		termbox.SetCell((pos[1]+1)*2, pos[0]+2, char, termbox.ColorGreen, defaultColour)
	}
}

func (f *Field) renderAfterGame() {

	utils.Tbprint(2, f.row+4, defaultColour, defaultColour, "┌─────────────────────┐")
	utils.Tbprint(2, f.row+5, defaultColour, defaultColour, "│  Enter - next game  │")
	utils.Tbprint(2, f.row+6, defaultColour, defaultColour, "│    Esc - menu       │")
	utils.Tbprint(2, f.row+7, defaultColour, defaultColour, "└─────────────────────┘")
	termbox.Flush()
}

func (f *Field) renderGame() {
	termbox.Clear(defaultColour, defaultColour)
	f.renderField()
	f.renderScoreAndMode()
	termbox.Flush()
}

func (g *Game) renderMenu() {
	termbox.Clear(defaultColour, defaultColour)
	utils.Tbprint(0, 0, defaultColour, defaultColour, "        ~~~snake~game~~☭ *        ")

	if g.field.borderMode {
		utils.Tbprint(0, 1, defaultColour, defaultColour, "         border mode: on           ")
	} else {
		utils.Tbprint(0, 1, defaultColour, defaultColour, "         border mode: off          ")
	}
	utils.Tbprint(0, 3, defaultColour, defaultColour, " press Enter to play snake         ")
	utils.Tbprint(0, 4, defaultColour, defaultColour, " press Esc to exit                 ")
	utils.Tbprint(0, 5, defaultColour, defaultColour, " press Space to switch border mode ")
	termbox.Flush()
}

func (g *Game) renderLossMsg() {
	utils.Tbprint(0, g.field.row+3, defaultColour, defaultColour, "You are proigral, prostofilya))0)")
	termbox.Flush()
}

func (g *Game) renderWinMsg() {
	utils.Tbprint(0, g.field.row+3, defaultColour, defaultColour, "Luchshiy, igrok v computer, graz!")
	termbox.Flush()
}
