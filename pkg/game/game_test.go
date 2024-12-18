package game

import (
	"errors"
	"testing"
)

func TestInitSnake(t *testing.T) {
	g := InitGame()
	result := func() (result error) {
		defer func() {
			if r := recover(); r != nil {
				result = errors.New(" ")
			}
		}()
		for range 1000 {
			g.field.initSnake(6)
		}
		return nil

	}()

	if result != nil {
		t.Error("Error while generating snake. Bad algorithm")
	}
}
