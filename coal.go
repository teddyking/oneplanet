package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	coalHeight    = 100
	coalWidth     = 150
	coalStartX    = 0
	coalStartY    = 0
	coalDropSpeed = 10

	numberHeight = 50
	numberWidth  = 50
)

type coal struct {
	mu sync.RWMutex

	time           int
	textures       []*sdl.Texture
	numberTextures []*sdl.Texture
	life           int

	x int32
	y int32

	fellToEarth     bool
	stoppedInTracks bool
}

func newCoal(r *sdl.Renderer) (*coal, error) {
	var textures []*sdl.Texture
	for i := 0; i < 4; i++ {
		texture, err := img.LoadTexture(r, fmt.Sprintf("%s/coal_%d.png", pathToImgs, i))
		if err != nil {
			return nil, fmt.Errorf("could not load texture: %v", err)
		}

		textures = append(textures, texture)
	}

	var numberTextures []*sdl.Texture
	for i := 1; i < 4; i++ {
		texture, err := img.LoadTexture(r, fmt.Sprintf("%s/%d.png", pathToImgs, i))
		if err != nil {
			return nil, fmt.Errorf("could not load texture: %v", err)
		}

		numberTextures = append(numberTextures, texture)
	}

	return &coal{
		life:           3,
		textures:       textures,
		numberTextures: numberTextures,
		x:              coalStartX,
		y:              coalStartY,
	}, nil
}

func (c *coal) collision() {
	c.life--
}

func (c *coal) update() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.time++
	c.y += coalDropSpeed

	if c.y > windowHeight {
		c.fellToEarth = true
	}

	if c.life <= 0 {
		c.stoppedInTracks = true
	}
}

func (c *coal) paint(r *sdl.Renderer) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	i := c.time % len(c.textures)
	rect := &sdl.Rect{W: coalWidth, H: coalHeight, X: c.x, Y: c.y}

	if err := r.Copy(c.textures[i], nil, rect); err != nil {
		return fmt.Errorf("could not copy coal: %v", err)
	}

	numberRect := &sdl.Rect{W: numberWidth, H: numberHeight, X: c.x + (coalWidth / 2) - (numberWidth / 2), Y: c.y + (coalHeight / 2) - (numberHeight * 2)}
	if err := r.Copy(c.numberTextures[c.life-1], nil, numberRect); err != nil {
		return fmt.Errorf("could not copy coal: %v", err)
	}

	return nil
}

func (c *coal) destroy() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, t := range c.textures {
		t.Destroy()
	}
}

func (c *coal) position() (int32, int32) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.x, c.y
}

func (c *coal) FellToEarth() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.fellToEarth
}

func (c *coal) StoppedInTracks() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.stoppedInTracks
}

//TODO: add restart logic
func (c *coal) restart() {
	c.mu.Lock()
	defer c.mu.Unlock()

	os.Exit(0)
}
