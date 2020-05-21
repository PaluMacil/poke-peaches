package main

import "github.com/hajimehoshi/ebiten"

type Background struct {
	image *ebiten.Image
	ratio float64
}

func (bg *Background) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	outsideWidth, outsideHeight := ebiten.WindowSize()
	bgWidth, bgHeight := bg.image.Size()
	screenRatio := float64(outsideWidth) / float64(outsideHeight)
	var scaleFactor float64
	if screenRatio < bg.ratio {
		scaleFactor = float64(outsideHeight) / float64(bgHeight)
	} else {
		scaleFactor = float64(outsideWidth) / float64(bgWidth)
	}
	opts.GeoM.Scale(scaleFactor, scaleFactor)
	screen.DrawImage(bg.image, opts)
}

func NewBackground(image *ebiten.Image) *Background {
	bgWidth, bgHeight := image.Size()
	return &Background{
		image: image,
		ratio: float64(bgWidth) / float64(bgHeight),
	}
}

type Game struct {
	Background *Background
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g Game) Update(screen *ebiten.Image) error {
	g.Background.Draw(screen)

	return nil
}
