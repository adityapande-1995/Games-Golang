// This module holds SDL Screen object

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// MyScreen : hold SDL objects
type MyScreen struct {
	winWidth, winHeight int
	window              *sdl.Window
	renderer            *sdl.Renderer
	tex                 *sdl.Texture
}

// NewScreen : acts like a constructor, returns screen and pixel array
func NewScreen(w, h int) (MyScreen, []byte) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(w), int32(h), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(w), int32(h))
	if err != nil {
		panic(err)
	}

	pixels := make([]byte, w*h*4)
	return MyScreen{w, h, window, renderer, tex}, pixels

}

// Shutdown : shutdown method for MyScreen class
func (MyScreen *MyScreen) Shutdown() {
	sdl.Quit()
	MyScreen.window.Destroy()
	MyScreen.renderer.Destroy()
	MyScreen.tex.Destroy()
}

// Update : This methods updates the screen with pixel array
func (MyScreen *MyScreen) Update(pixels []byte) {
	MyScreen.tex.Update(nil, pixels, MyScreen.winWidth*4)
	MyScreen.renderer.Copy(MyScreen.tex, nil, nil)
	MyScreen.renderer.Present()

}
