package main

import (
	//"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

// Class ball
type ball struct {
	pos
	radius int
	xv     float32
	yv     float32
	color  color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixels(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}

		}
	}

}
func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {
	ball.x += ball.xv
	ball.y += ball.yv

	if int(ball.y)-ball.radius < 0 || int(ball.y)+ball.radius > winHeight {
		ball.yv = -ball.yv
	}

	if ball.x < 0 || int(ball.x) > winWidth {
		ball.x = 300
		ball.y = 300
	}

	if int(ball.x) < int(leftPaddle.x)+leftPaddle.w/2 {
		if int(ball.y) > int(leftPaddle.y)-leftPaddle.h/2 && int(ball.y) < int(leftPaddle.y)+leftPaddle.h/2 {
			ball.xv = -ball.xv
		}

	}
	if int(ball.x) > int(rightPaddle.x)-rightPaddle.w/2 {
		if int(ball.y) > int(rightPaddle.y)-rightPaddle.h/2 && int(ball.y) < int(rightPaddle.y)+rightPaddle.h/2 {
			ball.xv = -ball.xv
		}

	}

}

// Class paddle
type paddle struct {
	pos
	w     int
	h     int
	color color
}

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x) - paddle.w/2
	startY := int(paddle.y) - paddle.h/2

	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixels(startX+x, startY+y, paddle.color, pixels)
		}
	}
}
func (paddle *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= 5
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += 5
	}

}
func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y
}

const winWidth, winHeight int = 800, 600

func clear(pixels []byte) {
	for i, _ := range pixels {
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

func main() {
	// Setting up SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		panic(err)
	}
	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	// Actual game loop start
	player1 := paddle{pos{50, 100}, 20, 100, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth) - 50, 100}, 20, 100, color{255, 255, 255}}
	ball := ball{pos{300, 300}, 20, 2, 2, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		clear(pixels)

		player1.update(keyState)
		ball.update(&player1, &player2)
		player1.draw(pixels)
		player2.draw(pixels)
		player2.aiUpdate(&ball)
		ball.draw(pixels)

		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16)
	}

}