package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	pathToFont          = "assets/fonts/block-cartoon.ttf"
	pathToBackgroundImg = "assets/img/background.jpg"
	pathToImgs          = "assets/img"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("unable to initialize SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("unable to initialize TTF: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(1920, 1080, 0)
	if err != nil {
		return fmt.Errorf("unable to create window: %v", err)
	}
	defer w.Destroy()

	sdl.PumpEvents()

	if err := drawTitle(r); err != nil {
		return fmt.Errorf("unable to draw title: %v", err)
	}

	// disply title for 1 second
	time.Sleep(time.Millisecond * 100)

	s, err := newScene(r)
	if err != nil {
		return fmt.Errorf("unable to create scene: %v", err)
	}
	defer s.destroy()

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Second*3, cancel)

	return <-s.run(ctx, r)
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()

	f, err := ttf.OpenFont(pathToFont, 2048)
	if err != nil {
		return fmt.Errorf("could not open font: %v", err)
	}
	defer f.Close()

	s, err := f.RenderUTF8Solid("One Planet", sdl.Color{R: 50, G: 200, B: 50, A: 255})
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()
	return nil
}
