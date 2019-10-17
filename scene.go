package main

import (
	"context"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	time int

	bg    *sdl.Texture
	coals []*sdl.Texture
	tree  *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, pathToBackgroundImg)
	if err != nil {
		return nil, fmt.Errorf("could not load background: %v", err)
	}

	var coals []*sdl.Texture
	for i := 0; i < 4; i++ {
		coal, err := img.LoadTexture(r, fmt.Sprintf("%s/coal_%d.png", pathToImgs, i))
		if err != nil {
			return nil, fmt.Errorf("could not load coal: %v", err)
		}

		coals = append(coals, coal)
	}

	return &scene{bg: bg, coals: coals}, nil
}

func (s *scene) run(ctx context.Context, r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		for range time.Tick(time.Millisecond * 60) {
			select {
			case <-ctx.Done():
				return
			default:
				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}

func (s *scene) paint(r *sdl.Renderer) error {
	s.time++

	r.Clear()

	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	i := s.time % len(s.coals)
	rect := &sdl.Rect{W: 150, H: 100, X: 200, Y: 200}
	if err := r.Copy(s.coals[i], nil, rect); err != nil {
		return fmt.Errorf("could not copy coal: %v", err)
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
