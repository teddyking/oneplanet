package main

import (
	"context"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	bg   *sdl.Texture
	coal *coal
	tree *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, pathToBackgroundImg)
	if err != nil {
		return nil, fmt.Errorf("could not load background: %v", err)
	}

	coal, err := newCoal(r)
	if err != nil {
		return nil, fmt.Errorf("could not create coal: %v", err)
	}

	return &scene{bg: bg, coal: coal}, nil
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
	r.Clear()

	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	if err := s.coal.paint(r); err != nil {
		return fmt.Errorf("could not pain coal: %v", err)
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
	s.coal.destroy()
}
