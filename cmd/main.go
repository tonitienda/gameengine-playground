package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	background := imdraw.New(nil)

	background.Color = pixel.RGB(0, 0, 0)
	background.Push(pixel.V(0, 0))
	background.Push(pixel.V(1024, 0))
	background.Push(pixel.V(1024, 768))
	background.Push(pixel.V(0, 768))
	background.Polygon(0)

	circle := imdraw.New(nil)
	circle.Color = pixel.RGB(0, 0, 1)
	circle.Push(pixel.V(1024/2, 768/2))
	circle.Circle(50, 0)

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		background.Draw(win)
		circle.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
