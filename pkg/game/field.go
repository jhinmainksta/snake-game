package game

import (
	"math/rand"
	"snake/pkg/utils"
)

type Snake struct {
	direction string
	poses     [][2]int
	len       int
}

func (game *Game) initSnake() {

	randRow := rand.Intn(game.row/2) + game.row/4
	randCol := rand.Intn(game.col/2) + game.col/4

	snakePos := make([][2]int, game.snakeLen)
	snakePos[0] = [2]int{randRow, randCol}

	for i := 1; i < game.snakeLen; i++ {

		poses := make([][2]int, 0)

		if snakePos[i-1][0] != 0 {
			newPos := [2]int{snakePos[i-1][0] - 1, snakePos[i-1][1]}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, game.col, game.row) {
				poses = append(poses, newPos)
			}
		}

		if snakePos[i-1][0] != game.row-1 {
			newPos := [2]int{snakePos[i-1][0] + 1, snakePos[i-1][1]}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, game.col, game.row) {
				poses = append(poses, newPos)
			}
		}

		if snakePos[i-1][1] != 0 {
			newPos := [2]int{snakePos[i-1][0], snakePos[i-1][1] - 1}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, game.col, game.row) {
				poses = append(poses, newPos)
			}
		}

		if snakePos[i-1][1] != game.col-1 {
			newPos := [2]int{snakePos[i-1][0], snakePos[i-1][1] + 1}
			curPos := append([][2]int{newPos}, snakePos...)
			if !(utils.ContainPos(snakePos, newPos)) && possibleMoves(curPos, game.col, game.row) {
				poses = append(poses, newPos)
			}
		}

		rngPos := rand.Intn(len(poses))
		snakePos[i] = poses[rngPos]
	}

	dx := snakePos[game.snakeLen-1][1] - snakePos[game.snakeLen-2][1]
	dy := snakePos[game.snakeLen-1][0] - snakePos[game.snakeLen-2][0]

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

	game.snake = &Snake{
		direction: direction,
		poses:     snakePos,
		len:       game.snakeLen,
	}
}

func (game *Game) moveSnake() {

	nextPos := [2]int{}
	switch game.snake.direction {
	case "w":
		nextPos = [2]int{game.snake.poses[len(game.snake.poses)-1][0] - 1, game.snake.poses[len(game.snake.poses)-1][1]}
	case "a":
		nextPos = [2]int{game.snake.poses[len(game.snake.poses)-1][0], game.snake.poses[len(game.snake.poses)-1][1] - 1}
	case "s":
		nextPos = [2]int{game.snake.poses[len(game.snake.poses)-1][0] + 1, game.snake.poses[len(game.snake.poses)-1][1]}
	case "d":
		nextPos = [2]int{game.snake.poses[len(game.snake.poses)-1][0], game.snake.poses[len(game.snake.poses)-1][1] + 1}
	}

	if !game.borderMode {
		nextPos = border(nextPos, game.row, game.col)
	}

	if game.wasEaten {
		game.snake.poses = append(game.snake.poses, nextPos)
		game.snake.len++
		game.wasEaten = false
	} else {
		utils.UpdSlice(game.snake.poses, append(game.snake.poses[1:], nextPos))
	}

	if game.foodIsEaten() {
		game.wasEaten = true
		if game.snake.len != game.col*game.row {
			game.score += 1
			game.initFood()
		} else {

		}
	}
}

func (game *Game) foodIsEaten() bool {
	return game.snake.poses[game.snake.len-1] == game.food
}

func (game *Game) ProcessTheMove() bool {

	game.moveSnake()

	game.renderGame()

	snakeHead := game.snake.poses[game.snake.len-1]
	snakeHitItself := utils.ContainPos(game.snake.poses[:game.snake.len-1], snakeHead)
	snakeHitTheWall := false

	if game.borderMode {
		snakeHitTheWall = snakeHead[0] == -1 || snakeHead[0] == game.row || snakeHead[1] == -1 || snakeHead[1] == game.col
	}

	return snakeHitItself || snakeHitTheWall
}

func (game *Game) initFood() {

	game.food[0] = rand.Intn(game.row)
	game.food[1] = rand.Intn(game.col)

	for utils.ContainPos(game.snake.poses, game.food) {
		game.food[0] = rand.Intn(game.row)
		game.food[1] = rand.Intn(game.col)
	}
}
