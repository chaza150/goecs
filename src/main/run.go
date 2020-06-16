package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"image"
	"os"

	_ "image/png"

	"ecs"
	"ecs/component"
	"ecs/entity"
	"ecs/io"
	"ecs/system"
	//"fmt"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	ecs := ecs.NewECS()
	ecs.EntManager.AddNewEntity("e1")
	ecs.EntManager.AddNewEntity("e2")
	ecs.EntManager.AddNewEntity("e3")
	ecs.EntManager.AddNewEntity("e4")

	ecs.EntManager.AddEntity(entity.NewPlayerEntity())

	ecs.EntManager.GetEntity("e1").AddComponent(component.ShoutComponent{"e1"})
	ecs.EntManager.GetEntity("e2").AddComponent(component.ShoutComponent{"e2"})
	ecs.EntManager.GetEntity("e3").AddComponent(component.ShoutComponent{"e3"})
	ecs.EntManager.GetEntity("e4").AddComponent(component.ShoutComponent{"e4"})

	shoutSys := system.ShoutSystem{}

	ecs.SysManager.AddSystem(shoutSys)

	io.LoadFile()

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
		ecs.UpdateSystems()

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
