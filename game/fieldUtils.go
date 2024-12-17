package game

import "snake/utils"

// move through field wals
func border(pos [2]int, row, col int) [2]int {

	if pos[0] == row {
		pos[0] = 0
	}

	if pos[1] == col {
		pos[1] = 0
	}

	if pos[0] == -1 {
		pos[0] = row - 1
	}

	if pos[1] == -1 {
		pos[1] = col - 1
	}

	return pos
}

// attempt to fix snake poses generation
func possibleMoves(Pos [][2]int, col, row int) bool {

	pos := make([][2]int, 0)

	if Pos[0][0] != 0 {
		newPos := [2]int{Pos[0][0] - 1, Pos[0][1]}
		if !utils.ContainPos(Pos, newPos) {
			pos = append(pos, newPos)
		}
	}

	if Pos[0][0] != row-1 {
		newPos := [2]int{Pos[0][0] + 1, Pos[0][1]}
		if !utils.ContainPos(Pos, newPos) {
			pos = append(pos, newPos)
		}
	}

	if Pos[0][1] != 0 {
		newPos := [2]int{Pos[0][0], Pos[0][1] - 1}
		if !utils.ContainPos(Pos, newPos) {
			pos = append(pos, newPos)
		}
	}

	if Pos[0][1] != col-1 {
		newPos := [2]int{Pos[0][0], Pos[0][1] + 1}
		if !utils.ContainPos(Pos, newPos) {
			pos = append(pos, newPos)
		}
	}

	if len(pos) > 1 {
		return true
	} else {
		return false
	}
}
