package main

import (
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

func (s *scene) run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		tick := time.Tick(time.Millisecond * 60)

		for {
			select {
			case e := <-events:
				if done := s.handleEvent(e); done {
					return
				}
			case <-tick:
				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch e := event.(type) {
	case *sdl.AudioDeviceEvent:
		// ignore for now
	case *sdl.WindowEvent:
		// ignore for now
	case *sdl.MouseMotionEvent:
		// ignore for now
	case *sdl.MouseButtonEvent:
		s.handleClick(e)
	case *sdl.QuitEvent:
		return true
	default:
		fmt.Printf("default event: %T\n", event)
	}

	return false
}

func (s *scene) handleClick(event *sdl.MouseButtonEvent) {
	mouseX, mouseY := event.X, event.Y
	mouseW, mouseH := int32(1), int32(1)

	coalX, coalY := s.coal.position()
	coalW, coalH := int32(coalWidth), int32(coalHeight)

	if mouseX < coalX+coalW &&
		mouseX+mouseW > coalX &&
		mouseY < coalY+coalH &&
		mouseY+mouseH > coalY {
		fmt.Println("BANG")
	}
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
