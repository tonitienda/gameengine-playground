package world

import (
	"fmt"
	"time"
)

const gravity = 9.8 // m/s2
const pixelsPerMeter = 100

type RGB struct {
	R byte
	G byte
	B byte
}

type Shape struct {
	PosX            float64
	PosY            float64
	VelX            float64
	VelY            float64
	BackgroundColor RGB
}

type World struct {
	currentTime time.Time
	Shapes      []Shape
}

func (w *World) Update() {
	if w.currentTime.IsZero() {
		w.currentTime = time.Now()
		return
	}

	elapsed := time.Since(w.currentTime)

	for i := range w.Shapes {
		w.Shapes[i].PosX += w.Shapes[i].VelX * elapsed.Seconds()
		w.Shapes[i].PosY += w.Shapes[i].VelY * elapsed.Seconds()

		// Counting that radius is 50 pixels from main
		// Fix this hardcoded value
		if w.Shapes[i].PosY > 25 {
			w.Shapes[i].VelY -= gravity * elapsed.Seconds() * pixelsPerMeter
		} else {
			w.Shapes[i].VelY = 0
			w.Shapes[i].PosY = 25
		}

		fmt.Printf("Shape %d: PosX: %f, PosY: %f, VelX: %f, VelY: %f\n", i, w.Shapes[i].PosX, w.Shapes[i].PosY, w.Shapes[i].VelX, w.Shapes[i].VelY)
	}

	w.currentTime = time.Now()
}
