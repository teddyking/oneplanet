package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type coal struct {
	time     int
	textures []*sdl.Texture

	y float64
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

	return &coal{textures: textures, y: 200}, nil
}

func (c *coal) paint(r *sdl.Renderer) error {
	c.time++
	c.y -= 10

	i := c.time % len(c.textures)
	rect := &sdl.Rect{W: 150, H: 100, X: 300, Y: (200 - int32(c.y))}

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
