package game

import (
	"errors"
	"testing"
)

func TestInitSnake(t *testing.T) {
	game := NewGame()

	game.snakeLen = 8
	game.col = 20
	game.row = 20
	numOfInit := 5000

	result := func() (result error) {
		defer func() {
			if r := recover(); r != nil {
				result = errors.New(" ")
			}
		}()
		for range numOfInit {
			game.initSnake()
		}
		return nil

	}()

	if result != nil {
		t.Error("Error while generating snake. Bad algorithm")
	}
}
