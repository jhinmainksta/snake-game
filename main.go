package main

import "snake/game"

func main() {
	Game := game.InitGame()

	go Game.HandleInput()
	Game.StartMenu()
}
