package main

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 500, 500

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}
func setPixels(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}

}

// *************  Start actual logic
const side int = 10

type color struct {
	r, g, b byte
}

// Class rect
type rect struct {
	x, y, s int
	color   color
}

func (rect *rect) draw(pixels []byte) { //top left corner mode draw
	//startX := rect.x - rect.s/2
	//startY := rect.y - rect.s/2

	for y := 0; y < rect.s; y++ {
		for x := 0; x < rect.s; x++ {
			//setPixels(startX+x, startY+y, rect.color, pixels)
			setPixels(rect.x+x, rect.y+y, rect.color, pixels)
		}
	}
}

// End class rect

// Class snake
type snake struct {
	body   []rect
	vx, vy int
	alive  bool
}

func (snake *snake) draw(pixels []byte) {
	for i := range snake.body {
		snake.body[i].draw(pixels)
	}
}
func (snake *snake) update(keyState []uint8) {
	// Check keyboard control
	if keyState[sdl.SCANCODE_UP] != 0 && snake.vy != 1 {
		snake.vx = 0
		snake.vy = -1
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 && snake.vy != -1 {
		snake.vx = 0
		snake.vy = 1
	}
	if keyState[sdl.SCANCODE_RIGHT] != 0 && snake.vx != -1 {
		snake.vx = 1
		snake.vy = 0
	}
	if keyState[sdl.SCANCODE_LEFT] != 0 && snake.vx != 1 {
		snake.vx = -1
		snake.vy = 0
	}

	// Check if alive
	if snake.body[len(snake.body)-1].x > winWidth || snake.body[len(snake.body)-1].x < 0 || snake.body[len(snake.body)-1].y > winHeight || snake.body[len(snake.body)-1].y < 0 {
		fmt.Println("Game over !!")
		snake.alive = false
		snake.body[len(snake.body)-1].color = color{255, 0, 0}
	}

	// Update position
	if snake.alive == true {
		head := snake.body[len(snake.body)-1]
		newhead := rect{head.x + snake.vx*side, head.y + snake.vy*side, side, color{255, 255, 255}}
		snake.body = append(snake.body, newhead) // Insert new head
		snake.body = snake.body[1:]              // Delete tail
	}

}
func (snake *snake) eat(f food) bool {
	return snake.body[len(snake.body)-1].x == f.r.x && snake.body[len(snake.body)-1].y == f.r.y
}
func (snake *snake) inclength() {
	head := snake.body[len(snake.body)-1]
	newhead := rect{head.x + snake.vx*side, head.y + snake.vy*side, side, color{255, 255, 255}}
	snake.body = append(snake.body, newhead) // Insert new head
}

// End class snake

// Class food
type food struct {
	r rect
}

func (food *food) draw(pixels []byte) {
	food.r.draw(pixels)
}
func (food *food) update() {
	food.r.x = (rand.Intn((winWidth/side)-3) + 1) * side
	food.r.y = (rand.Intn((winHeight/side)-3) + 1) * side
	fmt.Println("New food at :", food.r.x, food.r.y)
}

// Class food end

// Main func
func main() {

	S, pixels := NewScreen(winWidth, winHeight)

	// Actual game loop start
	b1 := []rect{rect{100, 100, side, color{255, 255, 255}}, rect{100, 100 + side, side, color{255, 255, 255}}}
	s1 := snake{b1, 0, 1, true}
	f := food{rect{200, 200, side, color{0, 255, 0}}}

	keyState := sdl.GetKeyboardState()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		clear(pixels)

		if s1.eat(f) {
			s1.inclength()
			f.update()
		}
		s1.update(keyState)
		s1.draw(pixels)
		f.draw(pixels)

		S.Update(pixels)
		sdl.Delay(50) // Game will run at 20 fps
	}

	S.Shutdown()

}
