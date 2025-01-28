package game

import (
	"log"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/nsf/termbox-go"
)

const (
	gameName = "THE SNAKE"

	play       = "play snake"
	borderMode = "border mode"
	settings   = "settings"
	exit       = "exit"

	defaultGameSpeed = 150
	defaultRow       = 10
	defaultCol       = 12
	menuBackspace    = 3
	menuFirstRow     = 1
	menuMiddle       = 20
)

var menu = [...]string{play, borderMode, settings, exit}

type Game struct {
	escChan   chan struct{}
	enterChan chan struct{}
	isPaused  bool
	pauseChan chan struct{}
	isStarted bool
	startChan chan struct{}
	toMenu    bool
	menuChan  chan struct{}

	gameSpeed        int
	directionsQuerry [2]string
	menuIndex        int

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
		escChan:   make(chan struct{}),
		enterChan: make(chan struct{}),
		isPaused:  false,
		pauseChan: make(chan struct{}),
		isStarted: false,
		startChan: make(chan struct{}),
		toMenu:    false,
		menuChan:  make(chan struct{}),

		gameSpeed:        defaultGameSpeed,
		directionsQuerry: [2]string{},
		menuIndex:        0,

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

func (game *Game) StartMenu() {

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	for {
		game.renderMenu()
		select {
		case <-game.enterChan:
			switch game.menuIndex {
			case 0:
				game.isStarted = true
				for game.isStarted {
					game.runGame()
					game.afterGame()
				}
			case 1:
				game.borderMode = !game.borderMode
			case 2:
			case 3:
				termbox.Clear(defaultColour, defaultColour)
				termbox.Flush()
				return
			}
		case <-game.escChan:
			termbox.Clear(defaultColour, defaultColour)
			termbox.Flush()
			return
		case <-game.menuChan:
		}
	}
}

func (game *Game) runGame() {

	game.directionsQuerry = [2]string{}
	game.initSnake()
	game.initFood()
	game.score = 0

	game.renderGame()

	time.Sleep(time.Second * 1)

	for {
		select {
		case <-game.escChan:
			game.isPaused = false
			game.toMenu = false
			game.isStarted = false
			return
		case <-game.pauseChan:
			game.isPaused = !game.isPaused
			if game.isPaused {
				game.renderPaused()
			}
		default:
			if !game.isPaused {
				game.setDirection()
				if game.snake.len == game.col*game.row {
					game.renderWinMsg()
					game.toMenu = true
					game.isPaused = false
					return
				}
				if failed := game.ProcessTheMove(); failed {

					game.renderLossMsg()
					game.toMenu = true
					game.isPaused = false
					return
				}
				time.Sleep(time.Millisecond * time.Duration(game.gameSpeed))
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func (game *Game) afterGame() {
	if !game.toMenu {
		return
	}

	defer func() { game.toMenu = false }()

	game.renderAfterGame()

	for {
		select {
		case <-game.startChan:
			return
		case <-game.escChan:
			game.isStarted = false
			return
		}
	}
}

func (game *Game) HandleInput() {

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if !game.isStarted {
			switch {
			case key == keyboard.KeyEnter:
				game.enterChan <- struct{}{}
			case key == keyboard.KeyEsc:
				game.escChan <- struct{}{}
			case key == keyboard.KeyArrowDown || key == keyboard.KeyArrowRight:
				game.menuIndex++
				if game.menuIndex == len(menu) {
					game.menuIndex = 0
				}
				game.menuChan <- struct{}{}
			case key == keyboard.KeyArrowUp || key == keyboard.KeyArrowLeft:
				game.menuIndex--
				if game.menuIndex == -1 {
					game.menuIndex = len(menu) - 1
				}
				game.menuChan <- struct{}{}
			}

		} else {
			switch {
			case key == keyboard.KeyEnter:
				if game.toMenu {
					game.startChan <- struct{}{}
				}
			case key == keyboard.KeySpace:
				if !game.toMenu {
					game.pauseChan <- struct{}{}
				}
			case key == keyboard.KeyEsc:
				game.isStarted = false
				game.escChan <- struct{}{}
			case char == 'w' || key == keyboard.KeyArrowUp:
				if game.directionsQuerry[0] != "" {
					game.directionsQuerry[1] = "w"
				} else if game.snake.direction != "w" {
					game.directionsQuerry[0] = "w"
				}
			case char == 'a' || key == keyboard.KeyArrowLeft:
				if game.directionsQuerry[0] != "" {
					game.directionsQuerry[1] = "a"
				} else if game.snake.direction != "a" {
					game.directionsQuerry[0] = "a"
				}
			case char == 's' || key == keyboard.KeyArrowDown:
				if game.directionsQuerry[0] != "" {
					game.directionsQuerry[1] = "s"
				} else if game.snake.direction != "s" {
					game.directionsQuerry[0] = "s"
				}
			case char == 'd' || key == keyboard.KeyArrowRight:
				if game.directionsQuerry[0] != "" {
					game.directionsQuerry[1] = "d"
				} else if game.snake.direction != "d" {
					game.directionsQuerry[0] = "d"
				}
			}
		}
	}
}

func (game *Game) setDirection() {
	if game.directionsQuerry[0] != "" {
		if !(game.directionsQuerry[0] == "w" && game.snake.direction == "s" ||
			game.directionsQuerry[0] == "s" && game.snake.direction == "w" ||
			game.directionsQuerry[0] == "a" && game.snake.direction == "d" ||
			game.directionsQuerry[0] == "d" && game.snake.direction == "a") {
			game.snake.direction = game.directionsQuerry[0]
		} else {
			game.directionsQuerry[1] = ""
		}
		game.directionsQuerry[0] = game.directionsQuerry[1]
		game.directionsQuerry[1] = ""
	}
}
