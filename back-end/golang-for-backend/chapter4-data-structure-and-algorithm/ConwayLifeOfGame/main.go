package main

import (
	"conwaygame/game"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func guiRunner() {
	size, width, height := float64(2), float64(400), float64(400)
	windowConfig := pixelgl.WindowConfig{
		Title:  "Conway's Game Of Life",
		Bounds: pixel.R(0, 0, width*size, height*size),
		VSync:  true,
	}

	window, err := pixelgl.NewWindow(windowConfig)
	if err != nil {
		panic(err)
	}

	gol := game.NewConwayGameOfLife(int(width), int(height), int(size))
	gol.InitWithRandom()
	//gol.InitWithMethuselah()

	for !window.Closed() {
		gol.PlayRound()
		window.Canvas().SetPixels(gol.GetPixels().Colors)
		window.Update()
	}
}

func main() {
	pixelgl.Run(guiRunner)
}
