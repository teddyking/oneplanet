package main

import (
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	explosionHeight = 256
	explosionWidth  = 128
)

type explosion struct {
	mu sync.RWMutex

	texture *sdl.Texture

	x int32
	y int32
}

func newExplosion(r *sdl.Renderer) (*explosion, error) {
	var texture *sdl.Texture
	texture, err := img.LoadTexture(r, fmt.Sprintf("%s/explosion.png", pathToImgs))
	if err != nil {
		return nil, fmt.Errorf("could not load texture: %v", err)
	}

	return &explosion{
		texture: texture,
	}, nil
}

func (e *explosion) paint(r *sdl.Renderer, x, y int32) error {
	e.mu.RLock()
	defer e.mu.RUnlock()

	rect := &sdl.Rect{
		W: 100,
		H: 100,
		X: 100,
		Y: 100,
	}

	if err := r.Copy(e.texture, nil, rect); err != nil {
		return fmt.Errorf("could not copy explosion: %v", err)
	}

	return nil
}

func (e *explosion) destroy() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.texture.Destroy()
}
