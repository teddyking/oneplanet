package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	coalHeight    = 100
	coalWidth     = 150
	coalStartX    = 0
	coalStartY    = 0
	coalDropSpeed = 10
)

type coal struct {
	time     int
	textures []*sdl.Texture
	life     int

	x int32
	y int32
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

	return &coal{
		life:     3,
		textures: textures,
		x:        coalStartX,
		y:        coalStartY,
	}, nil
}

func (c *coal) paint(r *sdl.Renderer) error {
	c.time++
	c.y += coalDropSpeed

	i := c.time % len(c.textures)
	rect := &sdl.Rect{W: coalWidth, H: coalHeight, X: c.x, Y: c.y}

	if err := r.Copy(c.textures[i], nil, rect); err != nil {
		return fmt.Errorf("could not copy coal: %v", err)
	}

	return nil
}

func (c *coal) destroy() {
	for _, t := range c.textures {
		t.Destroy()
	}
}

func (c *coal) position() (int32, int32) {
	return c.x, c.y
}
