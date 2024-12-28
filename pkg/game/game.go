package game

import (
	"log"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/nsf/termbox-go"
)

const (
	defaultGameSpeed = 150
	defaultRow       = 6
	defaultCol       = 6
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

	row        int
	col        int
	borderMode bool
	wasEaten   bool
	score      int
	food       [2]int
	snakeLen   int
	snake      *Snake
}

func NewGame() *Game {
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

		row:        defaultRow,
		col:        defaultCol,
		borderMode: false,
		wasEaten:   false,
		score:      0,
		food:       [2]int{},
		snakeLen:   4,
		snake:      &Snake{},
	}
}

func InitGame() {
	Game := NewGame()

	go Game.HandleInput()
	Game.StartMenu()
}

func (g *Game) StartMenu() {

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	for {
		g.renderMenu()
		select {
		case <-g.startChan:
			for g.isStarted {
				g.runGame()
				g.afterGame()
			}
		case <-g.escChan:
			termbox.Clear(defaultColour, defaultColour)
			termbox.Flush()
			return
		case <-g.borderChan:
			g.borderMode = !g.borderMode
		}
	}
}

func (g *Game) runGame() {

	g.directionsQuerry = [2]string{}
	g.initSnake()
	g.initFood()
	g.score = 0

	g.renderGame()

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
				g.renderPaused()
			}
		default:
			if !g.isPaused {
				g.setDirection()

				if g.snake.len == g.col*g.row {

					g.renderWinMsg()
					g.toMenu = true
					g.isPaused = false
					return
				}
				if failed := g.ProcessTheMove(); failed {

					g.renderLossMsg()
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

	g.renderAfterGame()

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
				} else if g.snake.direction != "w" {
					g.directionsQuerry[0] = "w"
				}
			case char == 'a' || key == keyboard.KeyArrowLeft:
				if g.directionsQuerry[0] != "" {
					g.directionsQuerry[1] = "a"
				} else if g.snake.direction != "a" {
					g.directionsQuerry[0] = "a"
				}
			case char == 's' || key == keyboard.KeyArrowDown:
				if g.directionsQuerry[0] != "" {
					g.directionsQuerry[1] = "s"
				} else if g.snake.direction != "s" {
					g.directionsQuerry[0] = "s"
				}
			case char == 'd' || key == keyboard.KeyArrowRight:
				if g.directionsQuerry[0] != "" {
					g.directionsQuerry[1] = "d"
				} else if g.snake.direction != "d" {
					g.directionsQuerry[0] = "d"
				}
			}
		}
	}
}

func (g *Game) setDirection() {
	if g.directionsQuerry[0] != "" {
		if !(g.directionsQuerry[0] == "w" && g.snake.direction == "s" ||
			g.directionsQuerry[0] == "s" && g.snake.direction == "w" ||
			g.directionsQuerry[0] == "a" && g.snake.direction == "d" ||
			g.directionsQuerry[0] == "d" && g.snake.direction == "a") {
			g.snake.direction = g.directionsQuerry[0]
		} else {
			g.directionsQuerry[1] = ""
		}
		g.directionsQuerry[0] = g.directionsQuerry[1]
		g.directionsQuerry[1] = ""
	}
}
