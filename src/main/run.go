package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"image"
	"os"

	_ "image/png"

	"ecs"
	"ecs/system"
	"fmt"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	ecs := ecs.NewECS()

	ecs.SysManager.AddSystem(system.ShoutSystem{})

	ecs.LoadEntityTypeData("res/ecs/data")

	ecs.InstantiateEntity("object", "obj")

	ecs.EntityLookup.PrintEntityTree()

	//SETUP For Pixel Window
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Game",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		//ecs.UpdateSystems()

		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
