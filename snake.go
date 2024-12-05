package main

import (
	"fmt"
	"math/rand"
)

var field [15][15]bool
var snakeLen = 6

func comparePoses(snakePoses [][2]int, posVar [2]int) bool {
	for _, val := range snakePoses {
		if val == posVar {
			return false
		}
	}
	return true
}

func main() {

	random := rand.Intn(15*15 - 1)

	snakePoses := make([][2]int, snakeLen)
	snakePoses[0] = [2]int{random / 15, random % 15}

	field[snakePoses[0][0]][snakePoses[0][1]] = true

	for i := 1; i < snakeLen; i++ {
		pos := make([][2]int, 0)
		if snakePoses[i-1][0] != 0 {
			posVar := [2]int{snakePoses[i-1][0] - 1, snakePoses[i-1][1]}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][0] != 14 {
			posVar := [2]int{snakePoses[i-1][0] + 1, snakePoses[i-1][1]}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][1] != 0 {
			posVar := [2]int{snakePoses[i-1][0], snakePoses[i-1][1] - 1}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}
		if snakePoses[i-1][1] != 14 {
			posVar := [2]int{snakePoses[i-1][0], snakePoses[i-1][1] + 1}
			if comparePoses(snakePoses, posVar) {
				pos = append(pos, posVar)
			}
		}

		rngPose := rand.Intn(len(pos))
		snakePoses[i] = pos[rngPose]
	}

	for i := 0; i < len(field); i++ {
		fmt.Print("_______________________________\n|")
		for j := 0; j < len(field[0]); j++ {
			symbol := ""
			if field[i][j] {
				symbol = "*"
			} else {
				symbol = " "
			}
			fmt.Printf("%s|", symbol)
		}
		fmt.Println()
	}
	fmt.Print("________________________________\n")

	fmt.Printf("row: %d, column: %d\n", random/15, random%15)
	fmt.Println(snakePoses)
}
