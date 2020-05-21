package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/jpeg"
	"log"
)

const (
	screenWidth  = 600
	screenHeight = 480
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Poke Peaches")
	ebiten.SetWindowResizable(true)
	stars, _, err := ebitenutil.NewImageFromFile("stars-2000x1500.jpg", ebiten.FilterNearest)
	if err != nil {
		log.Fatalln(err)
	}
	if err := ebiten.RunGame(Game{
		Background: NewBackground(stars),
	}); err != nil {
		log.Fatal(err)
	}
}
