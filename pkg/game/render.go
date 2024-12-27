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

	termbox.SetCell((f.food[1]+1)*2, f.food[0]+2, '*', termbox.ColorYellow, defaultColour)

	utils.Tbprint(0, f.row+2, defaultColour, defaultColour, "└"+strings.Repeat("─", (f.col+1)*2-1)+"┘")

	for i, pos := range f.snake.poses {
		char := '~'
		colour := termbox.ColorGreen
		if i == f.snake.len-1 {
			char = '☭'
			colour = termbox.ColorRed
		}
		termbox.SetCell((pos[1]+1)*2, pos[0]+2, char, colour, defaultColour)
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
