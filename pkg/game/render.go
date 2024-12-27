package game

import (
	"snake/pkg/utils"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"
)

const (
	defaultColour = termbox.ColorDefault
)

func (game *Game) renderInfo() {
	utils.Tbprint(game.col+1-3, game.row+3, defaultColour, defaultColour, "paused")
	termbox.Flush()
}

func (game *Game) renderScoreAndMode() {
	scoreStr := strconv.Itoa(game.score)
	msg := "borders off"
	if game.borderMode {
		msg = "borders on"
	}
	backspace := (game.col+1)*2 - len(scoreStr) - len(msg)
	utils.Tbprint(0, 0, defaultColour, defaultColour, msg+strings.Repeat(" ", backspace)+scoreStr)
}

func (game *Game) renderField() {

	utils.Tbprint(0, 1, defaultColour, defaultColour, "┌"+strings.Repeat("─", (game.col+1)*2-1)+"┐")

	for i := 1; i < game.row+1; i++ {
		termbox.SetCell(0, i+1, '│', defaultColour, defaultColour)
		termbox.SetCell((game.col+1)*2, i+1, '│', defaultColour, defaultColour)
	}

	termbox.SetCell((game.food[1]+1)*2, game.food[0]+2, '*', termbox.ColorYellow, defaultColour)

	utils.Tbprint(0, game.row+2, defaultColour, defaultColour, "└"+strings.Repeat("─", (game.col+1)*2-1)+"┘")

	for i, pos := range game.snake.poses {
		char := '~'
		colour := termbox.ColorGreen
		if i == game.snake.len-1 {
			char = '☭'
			colour = termbox.ColorRed
		}
		termbox.SetCell((pos[1]+1)*2, pos[0]+2, char, colour, defaultColour)
	}
}

func (game *Game) renderAfterGame() {

	utils.Tbprint(2, game.row+4, defaultColour, defaultColour, "┌─────────────────────┐")
	utils.Tbprint(2, game.row+5, defaultColour, defaultColour, "│  Enter - next game  │")
	utils.Tbprint(2, game.row+6, defaultColour, defaultColour, "│    Esc - menu       │")
	utils.Tbprint(2, game.row+7, defaultColour, defaultColour, "└─────────────────────┘")
	termbox.Flush()
}

func (game *Game) renderGame() {
	termbox.Clear(defaultColour, defaultColour)
	game.renderField()
	game.renderScoreAndMode()
	termbox.Flush()
}

func (game *Game) renderMenu() {
	termbox.Clear(defaultColour, defaultColour)
	utils.Tbprint(0, 0, defaultColour, defaultColour, "        ~~~snake~game~~☭ *        ")

	if game.borderMode {
		utils.Tbprint(0, 1, defaultColour, defaultColour, "         border mode: on           ")
	} else {
		utils.Tbprint(0, 1, defaultColour, defaultColour, "         border mode: off          ")
	}
	utils.Tbprint(0, 3, defaultColour, defaultColour, " press Enter to play snake         ")
	utils.Tbprint(0, 4, defaultColour, defaultColour, " press Esc to exit                 ")
	utils.Tbprint(0, 5, defaultColour, defaultColour, " press Space to switch border mode ")
	termbox.Flush()
}

func (game *Game) renderLossMsg() {
	utils.Tbprint(0, game.row+3, defaultColour, defaultColour, "You are proigral, prostofilya))0)")
	termbox.Flush()
}

func (game *Game) renderWinMsg() {
	utils.Tbprint(0, game.row+3, defaultColour, defaultColour, "Luchshiy, igrok v computer, graz!")
	termbox.Flush()
}
