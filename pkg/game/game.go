package game

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

const (
	defaultGameSpeed = 150
	defaultRow       = 10
	defaultCol       = 12
)

type Game struct {
	escChan    chan struct{}
	isPaused   bool
	pauseChan  chan struct{}
	isStarted  bool
	startChan  chan struct{}
	toMenu     bool
	borderChan chan struct{}

	gameSpeed        int
	directionsQuerry [2]string
	field            *Field
}

func InitGame() *Game {
	return &Game{
		escChan:    make(chan struct{}),
		isPaused:   false,
		pauseChan:  make(chan struct{}),
		isStarted:  false,
		startChan:  make(chan struct{}),
		toMenu:     false,
		borderChan: make(chan struct{}),

		gameSpeed:        defaultGameSpeed,
		directionsQuerry: [2]string{},
		field: &Field{
			row:        defaultRow,
			col:        defaultCol,
			borderMode: false,
			score:      0,
			food:       [2]int{},
			snakeLen:   6,
			snake:      &Snake{},
		},
	}
}

func (g *Game) runGame() {

	screen.Clear()
	screen.MoveTopLeft()

	g.field.initSnake()
	g.field.initFood()
	g.field.score = 0

	fmt.Println()
	g.field.renderField()

	time.Sleep(time.Second * 1)

	for {
		select {
		case <-g.escChan:
			g.isPaused = false
			g.toMenu = false
			g.isStarted = false
			return
		case <-g.pauseChan:
			g.isPaused = !g.isPaused
			if g.isPaused {
				g.field.renderInfo()
			}
		default:
			if !g.isPaused {

				screen.Clear()
				screen.MoveTopLeft()
				g.setDirection()

				if failed := g.field.ProcessTheMove(); failed {
					fmt.Println("You proigral, prostofilya))0)")
					g.toMenu = true
					g.isPaused = false
					return
				}
				if g.field.snakeLen == g.field.col*g.field.row {
					fmt.Println("Luchshiy iz lyudey pered monitorom, graz!")
					time.Sleep(time.Second * 2)
					g.toMenu = true
					g.isPaused = false
					return
				}
				time.Sleep(time.Millisecond * time.Duration(g.gameSpeed))
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func (g *Game) afterGame() {
	if !g.toMenu {
		return
	}

	defer func() { g.toMenu = false }()

	fmt.Println()
	fmt.Println("┌─────────────────────┐")
	fmt.Println("│  Enter - next game  │")
	fmt.Println("│    Esc - menu       │")
	fmt.Println("└─────────────────────┘")
	for {
		select {
		case <-g.startChan:
			return
		case <-g.escChan:
			g.isStarted = false
			return
		}
	}
}

func (g *Game) StartMenu() {
	for {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println("        ~~~snake~game~~☭ *        ")
		fmt.Println()
		if g.field.borderMode {
			fmt.Println("         border mode: on           ")
		} else {
			fmt.Println("         border mode: off          ")
		}
		fmt.Println()
		fmt.Println(" press Enter to play snake         ")
		fmt.Println(" press Esc to exit                 ")
		fmt.Println(" press Space to switch border mode ")
		select {
		case <-g.startChan:
			for g.isStarted {
				g.runGame()
				g.afterGame()
			}
		case <-g.escChan:
			screen.Clear()
			screen.MoveTopLeft()
			os.Exit(0)
		case <-g.borderChan:
			g.field.borderMode = !g.field.borderMode
		}
	}
}

func (g *Game) HandleInput() {

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if !g.isStarted {
			switch key {
			case keyboard.KeyEnter:
				g.isStarted = true
				g.startChan <- struct{}{}
			case keyboard.KeyEsc:
				g.escChan <- struct{}{}
			case keyboard.KeySpace:
				g.borderChan <- struct{}{}
			}
		} else {
			switch {
			case key == keyboard.KeyEnter:
				if g.toMenu {
					g.startChan <- struct{}{}
				}
			case key == keyboard.KeySpace:
				if !g.toMenu {
					g.pauseChan <- struct{}{}
				}
			case key == keyboard.KeyEsc:
				g.isStarted = false
				g.escChan <- struct{}{}
			case char == 'w' || key == keyboard.KeyArrowUp:
				if g.directionsQuerry[0] != "" {
					g.directionsQuerry[1] = "w"
				} else if g.field.snake.direction != "w" {
					g.directionsQuerry[0] = "w"
				}
			case char == 'a' || key == keyboard.KeyArrowLeft:
				if g.directionsQuerry[0] != "" {
					g.directionsQuerry[1] = "a"
				} else if g.field.snake.direction != "a" {
					g.directionsQuerry[0] = "a"
				}
			case char == 's' || key == keyboard.KeyArrowDown:
				if g.directionsQuerry[0] != "" {
					g.directionsQuerry[1] = "s"
				} else if g.field.snake.direction != "s" {
					g.directionsQuerry[0] = "s"
				}
			case char == 'd' || key == keyboard.KeyArrowRight:
				if g.directionsQuerry[0] != "" {
					g.directionsQuerry[1] = "d"
				} else if g.field.snake.direction != "d" {
					g.directionsQuerry[0] = "d"
				}
			}
		}
	}
}

func (g *Game) setDirection() {
	if g.directionsQuerry[0] != "" {
		if !(g.directionsQuerry[0] == "w" && g.field.snake.direction == "s" ||
			g.directionsQuerry[0] == "s" && g.field.snake.direction == "w" ||
			g.directionsQuerry[0] == "a" && g.field.snake.direction == "d" ||
			g.directionsQuerry[0] == "d" && g.field.snake.direction == "a") {
			g.field.snake.direction = g.directionsQuerry[0]
		} else {
			g.directionsQuerry[1] = ""
		}
		g.directionsQuerry[0] = g.directionsQuerry[1]
		g.directionsQuerry[1] = ""
	}
}
