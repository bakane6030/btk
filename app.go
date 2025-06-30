package btk

import (
	"io"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type DrawFunc func(c Ctx)

type Ctx interface {
	Text(text string)

	Textbox(hint string, input *string)

	Slider(min, max int, input *int)
	SliderFloat(min, max float32, input *float32)

	Image(image io.Reader)

	Add(d DrawFunc)

	Horizontal(d DrawFunc)
}

type App struct {
	renderer *sdl.Renderer
	window   *sdl.Window

	quit chan struct{}
}

func Init(title string, width, height int) (*App, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(width),
		int32(height),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return nil, err
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "linear")

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	renderer.SetDrawColor(hex("FF0000"))
	renderer.Clear()

	renderer.SetDrawColor(hex("00FF00"))
	renderer.DrawLine(100, 100, 200, 200)

	renderer.Present()

	quit := make(chan struct{}, 1)
	app := &App{renderer, window, quit}

	go app.pollEvents()

	return app, nil
}

func (a *App) Close() error {
	a.renderer.Destroy()
	a.window.Destroy()
	sdl.Quit()
	return nil
}

func (a *App) Quit() <-chan struct{} {
	return a.quit
}

func (a *App) pollEvents() {
	for {
		event := sdl.PollEvent()
		if event == nil {
			continue
		}

		switch event.(type) {
		case *sdl.QuitEvent:
			a.quit <- struct{}{}
			return
		}

		time.Sleep(16 * time.Millisecond)
	}
}

func (a *App) DrawWindow(d DrawFunc) {
}
