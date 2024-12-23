package utils

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func ContainPos(posCollection [][2]int, pos [2]int) bool {
	for _, val := range posCollection {
		if val == pos {
			return true
		}
	}
	return false
}

func UpdSlice(sliceToUpdate [][2]int, givenSlice [][2]int) {
	for i := range sliceToUpdate {
		sliceToUpdate[i] = givenSlice[i]
	}
}

func Tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
