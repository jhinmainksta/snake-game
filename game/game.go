package game

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

type Game struct {
	gameSpeed        int
	escChan          chan struct{}
	isPaused         bool
	pauseChan        chan struct{}
	isStarted        bool
	startChan        chan struct{}
	gameoverChan     chan struct{}
	snakeLen         int
	directionsQuerry [2]string
	field            *Field
}

func InitGame() *Game {
	return &Game{
		gameSpeed:        170,
		escChan:          make(chan struct{}),
		isPaused:         false,
		pauseChan:        make(chan struct{}),
		isStarted:        false,
		startChan:        make(chan struct{}),
		gameoverChan:     make(chan struct{}),
		snakeLen:         6,
		directionsQuerry: [2]string{},
		field: &Field{
			row:   10,
			col:   10,
			score: 0,
			food:  [2]int{},
			snake: &Snake{},
		},
	}
}

func (g *Game) runGame() {

	screen.Clear()
	screen.MoveTopLeft()

	g.field.initSnake(g.snakeLen)
	g.field.initFood()
	g.field.score = 0
	g.field.renderField()

	time.Sleep(time.Second * 2)

	for {
		select {
		case <-g.escChan:
			return
		case <-g.pauseChan:
			g.isPaused = !g.isPaused
		default:
			if !g.isPaused {

				screen.Clear()
				screen.MoveTopLeft()
				g.setDirection()

				msg := g.field.ProcessTheMove()
				if msg != "" {
					fmt.Println(msg)
					g.isStarted = !g.isStarted
					time.Sleep(time.Second * 1)
					return
				}
				time.Sleep(time.Millisecond * time.Duration(g.gameSpeed))
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func (g *Game) StartMenu() {
	for {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println("->     ~~~snake~game~~☭ *      <-")
		fmt.Println("->  press Enter to play snake  <-")
		fmt.Println("->      press Esc to exit      <-")
		select {
		case <-g.startChan:
			g.runGame()
		case <-g.escChan:
			screen.Clear()
			screen.MoveTopLeft()
			os.Exit(0)
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
			}
		} else {
			switch {
			case key == keyboard.KeySpace:
				g.pauseChan <- struct{}{}
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
