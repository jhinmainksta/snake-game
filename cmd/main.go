package main

import "snake/pkg/game"

func main() {

	Game := game.InitGame()

	go Game.HandleInput()
	Game.StartMenu()

}
