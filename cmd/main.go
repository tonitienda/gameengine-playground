package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/tonitienda/gameengine-playground/pkg/world"
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

	myworld := &world.World{}
	myworld.Shapes = append(myworld.Shapes, world.Shape{
		PosX:            1024 / 2,
		PosY:            740,
		VelX:            0,
		VelY:            0,
		BackgroundColor: world.RGB{R: 0, G: 0, B: 1},
	})

	for !win.Closed() {
		win.Clear(colornames.Black)

		// Assuming all shapes are circles for now
		// Recreating and redrawing the circle every frame
		// TODO - Optimize
		for _, shape := range myworld.Shapes {
			circle := imdraw.New(nil)
			circle.Color = pixel.RGB(0, 0, 1)
			circle.Push(pixel.V(shape.PosX, shape.PosY))
			circle.Circle(25, 0)
			circle.Draw(win)
		}

		win.Update()
		myworld.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
